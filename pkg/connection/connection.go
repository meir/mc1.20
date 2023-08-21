package connection

import (
	"bytes"
	"io"
	"net"

	"github.com/meir/mc1.20/pkg/packets"
	"golang.org/x/exp/slog"
)

type Connection struct {
	Conn  net.Conn
	State ConnectionState

	ProtocolVersion int
	ServerAddress   string
	Port            uint16
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn:  conn,
		State: StateHandshake,
	}
}

func (c *Connection) Manage() {
	for {
		var reader *bytes.Reader
		data, err := io.ReadAll(c.Conn)
		if err != nil {
			if err == io.EOF {
				slog.Info("connection closed")
				return
			}
			slog.Error("failed to read from connection", "err", err)
			return
		}

		reader = bytes.NewReader(data)

		var packet packets.Packet
		err = packets.Unmarshal(reader, &packet)
		if err != nil {
			slog.Error("failed to unmarshal packet", "err", err)
		}

		ok, err := GetHandlers(c.State, PacketId(packet.ID)).Handle(c, reader, packet)
		if err != nil {
			slog.Error("failed to handle packet", "err", err)
		}

		if !ok {
			slog.Error("no resolve for packet", "packet_id", packet.ID)
		}
	}
}

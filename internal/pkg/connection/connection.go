package connection

import (
	"bufio"
	"errors"
	"io"
	"net"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
	"golang.org/x/exp/slog"
)

type Connection struct {
	Conn  net.Conn
	State ConnectionState

	ProtocolVersion int
	ServerAddress   string
	Port            uint16

	Username string
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn:  conn,
		State: StateHandshake,
	}
}

func (c *Connection) Manage() {
	for {
		reader := bufio.NewReader(c.Conn)

		var packet packets.Packet
		err := packets.Unmarshal(reader, &packet)
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Debug("connection closed")
				return
			}
			slog.Error("failed to unmarshal packet", "err", err)
		}

		slog.Debug("packet received", "length", packet.Length, "id", packet.ID, "state", c.State)

		ok, err := GetHandlers(c.State, PacketId(packet.ID)).Handle(c, reader, packet)
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Debug("connection closed")
				return
			}
			slog.Error("failed to handle packet", "err", err, "id", packet.ID, "state", c.State)
		}

		if !ok {
			slog.Error("no resolve for packet", "packet_id", packet.ID)
		}
	}
}

func (c *Connection) Write(id PacketId, packet any) error {
	data, err := packets.Marshal(packet)
	if err != nil {
		return err
	}

	slog.Debug("sending packet", "length", len(data), "id", id)

	varintParser := packets.GetParser("varint")

	idData, err := varintParser.Marshal(reflect.ValueOf(int(id)))
	if err != nil {
		return err
	}

	data = append(idData, data...)

	lengthData, err := varintParser.Marshal(reflect.ValueOf(int(len(data))))
	if err != nil {
		return err
	}

	data = append(lengthData, data...)

	_, err = c.Conn.Write(data)
	if err != nil {
		return err
	}

	slog.Debug("packet sent", "length", len(data), "id", id)
	return nil
}

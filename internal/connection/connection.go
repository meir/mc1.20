package connection

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"reflect"
	"sync"

	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
	"golang.org/x/exp/slog"
)

type Connection struct {
	Mutex       sync.Mutex
	Conn        net.Conn
	packetQueue chan *bufio.Reader
	State       ConnectionState
	Server      *Server

	ProtocolVersion int
	ServerAddress   string
	Port            uint16

	Username    string
	UUID        datatypes.UUID
	VerifyToken []byte
}

func NewConnection(conn net.Conn, server *Server) *Connection {
	return &Connection{
		Mutex:       sync.Mutex{},
		Conn:        conn,
		packetQueue: make(chan *bufio.Reader, 100),
		State:       StateHandshake,
		Server:      server,
	}
}

func (c *Connection) Manage() {
	go c.handle()

	for {
		length, err := c.readPacketLength()
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Debug("disconnected")
				return
			}
		}

		packet := make([]byte, length)
		_, err = c.Conn.Read(packet)
		if err != nil {
			if errors.Is(err, io.EOF) {
				slog.Debug("invalid packet length")
			}
			slog.Error("failed to read packet", "err", err, "length", length)
		}

		reader := bufio.NewReader(bytes.NewReader(packet))
		c.packetQueue <- reader
	}
}

func (c *Connection) handle() {
	for {
		reader := <-c.packetQueue

		length := reader.Size()
		var packet packets.Packet
		err := packets.Unmarshal(reader, &packet)
		if err != nil {
			slog.Error("failed to unmarshal packet", "err", err, "length", length)
		}
		packet.Length = length

		c.Mutex.Lock()
		handler := GetHandlers(c.State, PacketId(packet.ID))
		c.Mutex.Unlock()

		ok, err := handler.Handle(c, reader, packet)
		if err != nil {
			slog.Error("failed to handle packet", "err", err, "id", packet.ID, "state", c.State)
			return
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

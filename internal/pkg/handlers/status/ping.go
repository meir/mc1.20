package status

import (
	"bufio"

	"github.com/meir/mc1.20/internal/connection"
	"github.com/meir/mc1.20/pkg/packets"
	"golang.org/x/exp/slog"
)

func init() {
	connection.RegisterHandler(connection.StateStatus, connection.PacketId(connection.ServerPacketPingRequest), HandlePingRequest)
}

func HandlePingRequest(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
	slog.Debug("ping request event",
		"length", packet.Length,
		"id", packet.ID,
	)

	pingPacket := PacketPingRequest{}
	err := packets.Unmarshal(reader, &pingPacket)
	if err != nil {
		return false, err
	}

	err = conn.Write(
		connection.PacketId(connection.ClientPacketPingResponse),
		PacketPingResponse{
			Payload: pingPacket.Payload,
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

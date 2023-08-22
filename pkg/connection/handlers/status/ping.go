package status

import (
	"bufio"

	"github.com/meir/mc1.20/pkg/connection"
	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/status"
	"golang.org/x/exp/slog"
)

func init() {
	connection.RegisterHandler(connection.StateStatus, connection.PacketId(connection.ServerPacketPingRequest), HandlePingRequest)
}

func HandlePingRequest(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
	slog.Info("ping request event",
		"length", packet.Length,
		"id", packet.ID,
	)

	pingPacket := status.PacketPingRequest{}
	err := packets.Unmarshal(reader, &pingPacket)
	if err != nil {
		return false, err
	}

	err = conn.Write(
		connection.PacketId(connection.ClientPacketPingResponse),
		status.PacketPingResponse{
			Payload: pingPacket.Payload,
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

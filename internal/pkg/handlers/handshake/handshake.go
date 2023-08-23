package handshake

import (
	"bufio"

	"github.com/meir/mc1.20/internal/pkg/connection"
	"github.com/meir/mc1.20/pkg/packets"
	"golang.org/x/exp/slog"
)

func init() {
	connection.RegisterHandler(connection.StateHandshake, connection.PacketId(connection.ServerPacketHandshake), HandleHandshake)
}

func HandleHandshake(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
	var handshakePacket PacketHandshake
	err := packets.Unmarshal(reader, &handshakePacket)
	if err != nil {
		return false, err
	}

	conn.ProtocolVersion = handshakePacket.ProtocolVersion
	conn.ServerAddress = handshakePacket.ServerAddress
	conn.Port = handshakePacket.ServerPort
	conn.State = connection.ConnectionState(handshakePacket.NextState)

	slog.Debug("handshake event",
		"length", packet.Length,
		"id", packet.ID,
		"protocol_version", handshakePacket.ProtocolVersion,
		"server_address", handshakePacket.ServerAddress,
		"server_port", handshakePacket.ServerPort,
		"next_state", handshakePacket.NextState,
	)

	return true, nil
}

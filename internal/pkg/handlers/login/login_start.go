package login

import (
	"bufio"

	"github.com/meir/mc1.20/internal/pkg/connection"
	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	connection.RegisterHandler(connection.StateLogin, connection.PacketId(connection.ServerPacketLoginStart), HandleLoginStart)
}

func HandleLoginStart(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
	loginStart := PacketLoginStart{}
	err := packets.Unmarshal(reader, &loginStart)
	if err != nil {
		return false, err
	}

	conn.Username = loginStart.Name

	return true, nil
}

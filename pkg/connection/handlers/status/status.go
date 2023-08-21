package status

import (
	"bytes"

	"github.com/meir/mc1.20/pkg/connection"
	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/objects"
)

func init() {
	connection.RegisterHandler(connection.StateStatus, connection.PacketId(connection.ServerPacketStatusRequest), HandleStatusRequest)
}

func HandleStatusRequest(conn *connection.Connection, reader *bytes.Reader, packet packets.Packet) (bool, error) {
	status := objects.StatusResponse{
		Version: objects.StatusVersion{
			Name:     "1.20",
			Protocol: 754,
		},
	}
	return true, nil
}

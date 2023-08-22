package status

import (
	"bufio"
	"encoding/json"

	"github.com/meir/mc1.20/pkg/connection"
	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/objects"
	"github.com/meir/mc1.20/pkg/packets/status"
	"golang.org/x/exp/slog"
)

func init() {
	connection.RegisterHandler(connection.StateStatus, connection.PacketId(connection.ServerPacketStatusRequest), HandleStatusRequest)
}

func HandleStatusRequest(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
	statusResponse := objects.StatusResponse{
		Version: objects.StatusVersion{
			Name:     "1.20",
			Protocol: 763,
		},
		Players: objects.StatusPlayers{
			Max:    100,
			Online: 0,
		},
		Description: objects.Chat{
			Text: "Deepfryer bananas",
		},
		Favicon:            "",
		EnforcesSecureChat: false,
		PreviewsChat:       false,
	}

	slog.Info("status request event",
		"length", packet.Length,
		"id", packet.ID,
	)

	statusString, err := json.Marshal(statusResponse)
	if err != nil {
		return false, err
	}

	err = conn.Write(
		connection.PacketId(connection.ClientPacketStatusResponse),
		status.PacketStatusResponse{
			Status: string(statusString),
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

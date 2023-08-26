package status

import (
	"bufio"
	"encoding/json"

	"github.com/meir/mc1.20/internal/connection"
	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
	"golang.org/x/exp/slog"
)

func init() {
	connection.RegisterHandler(connection.StateStatus, connection.PacketId(connection.ServerPacketStatusRequest), HandleStatusRequest)
}

func HandleStatusRequest(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
	statusResponse := datatypes.StatusResponse{
		Version: datatypes.StatusVersion{
			Name:     "1.20",
			Protocol: 763,
		},
		Players: datatypes.StatusPlayers{
			Max:    69,
			Online: 1,
		},
		Description: datatypes.Chat{
			Text:  "This is bananas!",
			Color: "yellow",
			Bold:  true,
		},
		Favicon:            "",
		EnforcesSecureChat: false,
		PreviewsChat:       false,
	}

	slog.Debug("status request event",
		"length", packet.Length,
		"id", packet.ID,
	)

	statusString, err := json.Marshal(statusResponse)
	if err != nil {
		return false, err
	}

	err = conn.Write(
		connection.PacketId(connection.ClientPacketStatusResponse),
		PacketStatusResponse{
			Status: string(statusString),
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

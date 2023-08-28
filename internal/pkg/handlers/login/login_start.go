package login

import (
	"bufio"
	"crypto/rand"
	"crypto/x509"
	"fmt"

	"github.com/meir/mc1.20/internal/connection"
	"github.com/meir/mc1.20/internal/pkg/handlers/play"
	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
	"golang.org/x/exp/slog"
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

	slog.Debug(fmt.Sprintf("login request from %s", loginStart.Name))

	conn.Username = loginStart.Name
	if loginStart.PlayerUUID == nil {
		conn.UUID = *loginStart.PlayerUUID
	}

	// TODO: Disconnect user paths
	// conn.DisconnectUser(true, datatypes.Chat{
	// 	Text: "You have been disconnected.",
	// })

	conn.VerifyToken = make([]byte, 16)
	_, err = rand.Read(conn.VerifyToken)
	if err != nil {
		return false, err
	}

	//Create DER format of the RSA public key
	publicKey, err := x509.MarshalPKIXPublicKey(&conn.Server.PrivateKey.PublicKey)
	if err != nil {
		return false, err
	}

	encryptionRequest := PacketEncryptionRequest{
		ServerID:    "",
		PublicKey:   publicKey,
		VerifyToken: conn.VerifyToken,
	}

	err = conn.Write(connection.PacketId(connection.ClientPacketEncryptionRequest), encryptionRequest)
	if err != nil {
		return false, err
	}

	slog.Debug("wrote encryption request",
		"length", packet.Length,
		"id", packet.ID,
		"username", loginStart.Name,
	)

	err = conn.Write(connection.PacketId(connection.ClientPacketLoginPlay), play.PacketLoginPlay{
		EntityID:            0,
		IsHardcore:          false,
		Gamemode:            0,
		PreviousGamemode:    0,
		DimensionNames:      []string{"minecraft:overworld"},
		RegistryCodec:       nil,
		DimensionType:       "minecraft:overworld",
		DimensionName:       "minecraft:overworld",
		HashedSeed:          0,
		MaxPlayers:          0,
		ViewDistance:        0,
		SimulationDistance:  0,
		ReducedDebugInfo:    false,
		EnableRespawnScreen: false,
		IsDebug:             false,
		IsFlat:              false,
		DeathLocation: struct {
			Dimension string             `packet:"identifier"`
			Location  datatypes.Position `packet:"position"`
		}{
			Dimension: "minecraft:overworld",
			Location: datatypes.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
		},
		PortalCooldown: 0,
	})

	return true, nil
}

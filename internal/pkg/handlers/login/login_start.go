package login

import (
	"bufio"
	"crypto/rand"
	"crypto/x509"
	"fmt"

	"github.com/meir/mc1.20/internal/connection"
	"github.com/meir/mc1.20/pkg/packets"
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

	return true, nil
}

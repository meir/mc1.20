package login

import "github.com/meir/mc1.20/pkg/packets/datatypes"

type PacketLoginStart struct {
	Name       string          `packet:"string"`
	PlayerUUID *datatypes.UUID `packet:"uuid,optional"`
}

type PacketEncryptionRequest struct {
	ServerID    string `packet:"string"`
	PublicKey   []byte `packet:"byte,array"`
	VerifyToken []byte `packet:"byte,array"`
}

type PacketEncryptionResponse struct {
	SharedSecret []byte `packet:"byte,array"`
	VerifyToken  []byte `packet:"byte,array"`
}

type PacketLoginSuccess struct {
	UUID       datatypes.UUID `packet:"uuid"`
	Username   string         `packet:"string"`
	Properties []struct {
		Name      string  `packet:"string"`
		Value     string  `packet:"string"`
		Signature *string `packet:"string,optional"`
	} `packet:"struct,array"`
}

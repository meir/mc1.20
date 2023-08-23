package login

type PacketLoginStart struct {
	Name       string  `packet:"string"`
	PlayerUUID *string `packet:"uuid,optional"`
}

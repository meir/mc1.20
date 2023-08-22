package handshake

type PacketHandshake struct {
	ProtocolVersion int    `packet:"varint"`
	ServerAddress   string `packet:"string"`
	ServerPort      uint16 `packet:"ushort"`
	NextState       int    `packet:"varint"`
}

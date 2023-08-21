package handshake

import "github.com/meir/mc1.20/pkg/packets"

type PacketHandshake struct {
	*packets.Packet

	ProtocolVersion int    `packet:"varint"`
	ServerAddress   string `packet:"string"`
	ServerPort      uint16 `packet:"ushort"`
	NextState       int    `packet:"varint"`
}

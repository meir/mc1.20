package status

import "github.com/meir/mc1.20/pkg/packets"

type PacketStatusResponse struct {
	*packets.Packet

	Status string `packet:"string"`
}

type PacketPingResponse struct {
	*packets.Packet

	Payload int64 `packet:"long"`
}

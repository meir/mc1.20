package status

import "github.com/meir/mc1.20/pkg/packets"

type PacketStatusRequest struct {
	*packets.Packet
}

type PacketPingRequest struct {
	*packets.Packet

	Payload int64 `packet:"long"`
}

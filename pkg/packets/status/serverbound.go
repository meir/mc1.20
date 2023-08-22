package status

type PacketStatusRequest struct {
}

type PacketPingRequest struct {
	Payload int64 `packet:"long"`
}

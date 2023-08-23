package status

type PacketStatusRequest struct {
}

type PacketPingRequest struct {
	Payload int64 `packet:"long"`
}

type PacketStatusResponse struct {
	Status string `packet:"string"`
}

type PacketPingResponse struct {
	Payload int64 `packet:"long"`
}

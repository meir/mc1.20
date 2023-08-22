package status

type PacketStatusResponse struct {
	Status string `packet:"string"`
}

type PacketPingResponse struct {
	Payload int64 `packet:"long"`
}

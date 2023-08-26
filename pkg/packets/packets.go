package packets

// Packet is the initial data structure given at the start of every packet.
type Packet struct {
	Length int
	ID     int `packet:"varint"`
}

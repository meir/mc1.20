package packets

type Packet struct {
	Length int `packet:"varint"`
	ID     int `packet:"varint"`

	Data []byte
}

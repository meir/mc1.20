package connection

import "fmt"

const SEGMENT_BITS byte = 0x7f
const CONTINUE_BIT byte = 0x80

func (c *Connection) readPacketLength() (int, error) {
	// read varint
	length := 0
	pos := 0

	for {
		sb := make([]byte, 1)
		_, err := c.Conn.Read(sb)
		b := sb[0]
		if err != nil {
			return 0, err
		}

		length |= int(b&SEGMENT_BITS) << uint(pos)

		if b&CONTINUE_BIT == 0 {
			break
		}

		pos += 7

		if pos > 32 {
			return 0, fmt.Errorf("varint too big")
		}
	}

	return length, nil
}

package parsers

import (
	"bytes"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("varint32", &VarintParser{32})
	packets.RegisterParser("varint64", &VarintParser{64})
}

type VarintParser struct {
	bits int
}

const SEGMENT_BITS byte = 0x7f
const CONTINUE_BIT byte = 0x80

func (p *VarintParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Ptr,
		}
	}

	if value.Elem().Kind() != reflect.Int {
		return &packets.ErrInvalidKind{
			value.Elem().Kind(),
			reflect.Int,
		}
	}

	// decode varint
	v := int64(0)
	pos := 0

	for {
		b, err := data.ReadByte()
		if err != nil {
			return err
		}

		v |= int64(b&SEGMENT_BITS) << uint(pos)

		if b&CONTINUE_BIT == 0 {
			break
		}

		pos += 7

		if pos > p.bits {
			return packets.ErrVarintTooBig
		}
	}

	// set value
	value.Elem().SetInt(v)

	return nil
}

func (p *VarintParser) Marshal(data reflect.Value) ([]byte, error) {
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}

	if data.Kind() != reflect.Int {
		return nil, &packets.ErrInvalidKind{
			data.Kind(),
			reflect.Int,
		}
	}

	v := data.Int()
	bytes := []byte{}

	for {
		b := byte(v & 0x7f)

		v >>= 7

		if v != 0 {
			b |= CONTINUE_BIT
		}

		bytes = append(bytes, b)

		if v == 0 {
			break
		}
	}

	return bytes, nil
}

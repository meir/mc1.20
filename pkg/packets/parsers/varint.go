package parsers

import (
	"bufio"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("varint", &VarintParser{32})
	packets.RegisterParser("varint32", &VarintParser{32})
	packets.RegisterParser("varint64", &VarintParser{64})
	packets.RegisterParser("varlong", &VarintParser{64})
}

type VarintParser struct {
	bits int
}

const SEGMENT_BITS byte = 0x7f
const CONTINUE_BIT byte = 0x80

func (p *VarintParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectKind("varint", value, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64); err != nil {
		return err
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
			return ErrVarintTooBig
		}
	}

	// set value
	value.SetInt(v)

	return nil
}

func (p *VarintParser) Marshal(data reflect.Value) ([]byte, error) {
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}

	if err := expectKind("varint", data, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64); err != nil {
		return nil, err
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

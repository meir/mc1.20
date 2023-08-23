package parsers

import (
	"bufio"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	bp := &ByteParser{}
	packets.RegisterParser("byte", bp)
	packets.RegisterParser("ubyte", bp)
}

type ByteParser struct{}

func (bp *ByteParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectKind("byte", value, reflect.Uint8); err != nil {
		return err
	}

	b, err := data.ReadByte()
	if err != nil {
		return err
	}

	value.SetUint(uint64(b))

	return nil
}

func (b *ByteParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectKind("byte", value, reflect.Uint8); err != nil {
		return nil, err
	}

	return []byte{uint8(value.Uint())}, nil
}

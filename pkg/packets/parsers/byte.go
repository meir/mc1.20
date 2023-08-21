package parsers

import (
	"bytes"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	bp := &ByteParser{}
	packets.RegisterParser("byte", bp)
	packets.RegisterParser("ubyte", bp)
}

type ByteParser struct{}

func (bp *ByteParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Ptr,
		}
	}

	if value.Elem().Kind() != reflect.Uint8 {
		return &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Uint8,
		}
	}

	b, err := data.ReadByte()
	if err != nil {
		return err
	}

	value.Elem().SetUint(uint64(b))

	return nil
}

func (b *ByteParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Uint8 {
		return nil, &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Uint8,
		}
	}

	return []byte{uint8(value.Uint())}, nil
}

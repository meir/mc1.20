package parsers

import (
	"bufio"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("bool", &BoolParser{})
}

// BoolParser is a parser for bool type
// bool bytes are either 0x01 or 0x00
type BoolParser struct{}

func (p *BoolParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Bool {
		return &packets.ErrInvalidKind{
			"bool",
			value.Kind(),
			reflect.Bool,
		}
	}

	b, err := data.ReadByte()
	if err != nil {
		return err
	}

	value.SetBool(int(b) > 0)

	return nil
}

func (p *BoolParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Bool {
		return nil, &packets.ErrInvalidKind{
			"bool",
			value.Kind(),
			reflect.Bool,
		}
	}

	var b byte = 0x00
	if value.Bool() {
		b = 0x01
	}

	return []byte{b}, nil
}

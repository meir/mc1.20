package parsers

import (
	"bytes"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("string", &StringParser{
		varintParser: &VarintParser{
			32,
		},
	})
}

type StringParser struct {
	varintParser *VarintParser
}

func (p *StringParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Ptr,
		}
	}

	if value.Elem().Kind() != reflect.String {
		return &packets.ErrInvalidKind{
			value.Elem().Kind(),
			reflect.String,
		}
	}

	length := int32(0)
	err := p.varintParser.Unmarshal(data, reflect.ValueOf(&length))
	if err != nil {
		return err
	}

	if length < 0 {
		return &packets.ErrInvalidLength{
			int(length),
		}
	}

	if length > 32767 {
		return &packets.ErrInvalidLength{
			int(length),
		}
	}

	if length == 0 {
		value.Elem().SetString("")
		return nil
	}

	bytes := make([]byte, length)
	_, err = data.Read(bytes)
	if err != nil {
		return err
	}

	value.Elem().SetString(string(bytes))
	return nil
}

func (p *StringParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.String {
		return nil, &packets.ErrInvalidKind{
			value.Kind(),
			reflect.String,
		}
	}

	length := int32(len(value.String()))

	if length < 0 {
		return nil, &packets.ErrInvalidLength{
			int(length),
		}
	}

	if length > 32767 {
		return nil, &packets.ErrInvalidLength{
			int(length),
		}
	}

	varint, err := p.varintParser.Marshal(reflect.ValueOf(length))
	if err != nil {
		return nil, err
	}

	return append(varint, []byte(value.String())...), nil
}

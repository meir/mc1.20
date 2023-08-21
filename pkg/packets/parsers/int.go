package parsers

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("int8", &IntParser{8, reflect.Int8})
	packets.RegisterParser("uint8", &IntParser{8, reflect.Uint8})
	packets.RegisterParser("int16", &IntParser{16, reflect.Int16})
	packets.RegisterParser("uint16", &IntParser{16, reflect.Uint16})
	packets.RegisterParser("int32", &IntParser{32, reflect.Int32})
	packets.RegisterParser("uint32", &IntParser{32, reflect.Uint32})
	packets.RegisterParser("int64", &IntParser{64, reflect.Int64})
	packets.RegisterParser("uint64", &IntParser{64, reflect.Uint64})

	packets.RegisterParser("int", &IntParser{32, reflect.Int32})
	packets.RegisterParser("uint", &IntParser{32, reflect.Uint32})

	packets.RegisterParser("short", &IntParser{16, reflect.Int16})
	packets.RegisterParser("ushort", &IntParser{16, reflect.Uint16})
	packets.RegisterParser("long", &IntParser{64, reflect.Int64})
	packets.RegisterParser("ulong", &IntParser{64, reflect.Uint64})
}

type IntParser struct {
	bits int
	kind reflect.Kind
}

func (p *IntParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return &packets.ErrInvalidKind{
			Kind:   value.Kind(),
			Wanted: reflect.Ptr,
		}
	}

	if value.Elem().Kind() != p.kind {
		return &packets.ErrInvalidKind{
			Kind:   value.Elem().Kind(),
			Wanted: p.kind,
		}
	}

	var v int64
	b := make([]byte, p.bits/8)
	if _, err := data.Read(b); err != nil {
		return err
	}

	switch p.bits {
	case 8:
		v = int64(b[0])
	case 16:
		v = int64(binary.BigEndian.Uint16(b))
	case 32:
		v = int64(binary.BigEndian.Uint32(b))
	case 64:
		v = int64(binary.BigEndian.Uint64(b))
	}

	value.Elem().SetInt(v)

	return nil
}

func (p *IntParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != p.kind {
		return nil, &packets.ErrInvalidKind{
			Kind:   value.Kind(),
			Wanted: p.kind,
		}
	}

	b := make([]byte, p.bits/8)
	switch p.bits {
	case 8:
		b[0] = byte(value.Int())
	case 16:
		binary.BigEndian.PutUint16(b, uint16(value.Int()))
	case 32:
		binary.BigEndian.PutUint32(b, uint32(value.Int()))
	case 64:
		binary.BigEndian.PutUint64(b, uint64(value.Int()))
	}

	return b, nil
}

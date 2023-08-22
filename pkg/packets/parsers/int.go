package parsers

import (
	"bufio"
	"encoding/binary"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("int8", &IntParser{8, false, reflect.Int8})
	packets.RegisterParser("uint8", &IntParser{8, true, reflect.Uint8})
	packets.RegisterParser("int16", &IntParser{16, false, reflect.Int16})
	packets.RegisterParser("uint16", &IntParser{16, true, reflect.Uint16})
	packets.RegisterParser("int32", &IntParser{32, false, reflect.Int32})
	packets.RegisterParser("uint32", &IntParser{32, true, reflect.Uint32})
	packets.RegisterParser("int64", &IntParser{64, false, reflect.Int64})
	packets.RegisterParser("uint64", &IntParser{64, true, reflect.Uint64})

	packets.RegisterParser("int", &IntParser{32, false, reflect.Int32})
	packets.RegisterParser("uint", &IntParser{32, true, reflect.Uint32})

	packets.RegisterParser("short", &IntParser{16, false, reflect.Int16})
	packets.RegisterParser("ushort", &IntParser{16, true, reflect.Uint16})
	packets.RegisterParser("long", &IntParser{64, false, reflect.Int64})
	packets.RegisterParser("ulong", &IntParser{64, true, reflect.Uint64})
}

type IntParser struct {
	bits     int
	unsigned bool
	kind     reflect.Kind
}

func (p *IntParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != p.kind {
		return &packets.ErrInvalidKind{
			Kind:   value.Kind(),
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

	if p.unsigned {
		value.SetUint(uint64(v))
	} else {
		value.SetInt(v)
	}

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

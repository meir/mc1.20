package parsers

import (
	"bufio"
	"encoding/binary"
	"math"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("float32", &FloatParser{bits: 32, kind: reflect.Float32})
	packets.RegisterParser("float64", &FloatParser{bits: 64, kind: reflect.Float64})

	packets.RegisterParser("float", &FloatParser{bits: 32, kind: reflect.Float32})
	packets.RegisterParser("double", &FloatParser{bits: 64, kind: reflect.Float64})
}

type FloatParser struct {
	bits int
	kind reflect.Kind
}

func (p *FloatParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != p.kind {
		return &packets.ErrInvalidKind{
			Kind:   value.Kind(),
			Wanted: p.kind,
		}
	}

	var v float64
	b := make([]byte, p.bits/8)
	if _, err := data.Read(b); err != nil {
		return err
	}

	switch p.kind {
	case reflect.Float32:
		bits := binary.BigEndian.Uint32(b)
		v32 := math.Float32frombits(bits)
		v = float64(v32)
	case reflect.Float64:
		bits := binary.BigEndian.Uint64(b)
		v = math.Float64frombits(bits)
	}

	value.SetFloat(v)

	return nil
}

func (p *FloatParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != p.kind {
		return nil, &packets.ErrInvalidKind{
			Kind:   value.Kind(),
			Wanted: p.kind,
		}
	}

	var v float64

	switch p.kind {
	case reflect.Float32:
		v = float64(value.Float())
	case reflect.Float64:
		v = value.Float()
	}

	bits := math.Float64bits(v)
	b := make([]byte, p.bits/8)
	binary.BigEndian.PutUint64(b, bits)

	return b, nil
}

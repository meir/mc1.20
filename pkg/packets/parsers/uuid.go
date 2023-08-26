package parsers

import (
	"bufio"
	"encoding/binary"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
)

func init() {
	packets.RegisterParser("uuid", &UUIDParser{
		kind: reflect.TypeOf(datatypes.UUID{}),
		int64Parser: &IntParser{
			64,
			false,
			reflect.Int64,
		},
	})
}

type UUIDParser struct {
	kind reflect.Type

	int64Parser *IntParser
}

func (p *UUIDParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {

	if err := expectType("uuid", value, p.kind); err != nil {
		return err
	}

	var most, least int64

	err := p.int64Parser.Unmarshal(data, reflect.ValueOf(&most))
	if err != nil {
		return err
	}

	err = p.int64Parser.Unmarshal(data, reflect.ValueOf(&least))
	if err != nil {
		return err
	}

	uuid := datatypes.NewUUID()

	binary.LittleEndian.PutUint64(uuid[:8], uint64(most))
	binary.LittleEndian.PutUint64(uuid[8:], uint64(least))

	if value.Kind() == reflect.Ptr {
		value.Set(reflect.ValueOf(&uuid))
	} else {
		value.Set(reflect.ValueOf(uuid))
	}

	return nil
}

func (p *UUIDParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("uuid", value, p.kind); err != nil {
		return nil, err
	}

	uuid := value.Interface().(datatypes.UUID)

	var most, least int64

	most = int64(binary.LittleEndian.Uint64(uuid[:8]))
	least = int64(binary.LittleEndian.Uint64(uuid[8:]))

	mostBytes, err := p.int64Parser.Marshal(reflect.ValueOf(most))
	if err != nil {
		return nil, err
	}

	leastBytes, err := p.int64Parser.Marshal(reflect.ValueOf(least))
	if err != nil {
		return nil, err
	}

	return append(mostBytes, leastBytes...), nil
}

package parsers

import (
	"bufio"
	"reflect"

	"github.com/beito123/nbt"
	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("nbt", &NBTParser{
		reflect.TypeOf(&nbt.Stream{}),
	})
}

type NBTParser struct {
	entityMetadataType reflect.Type
}

func (p *NBTParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("nbt", value, p.entityMetadataType); err != nil {
		return err
	}

	stream, err := nbt.FromReader(data, nbt.LittleEndian)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(stream))

	return nil
}

func (p *NBTParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("nbt", value, p.entityMetadataType); err != nil {
		return nil, err
	}

	stream := value.Interface().(*nbt.Stream)

	return stream.Bytes(), nil
}

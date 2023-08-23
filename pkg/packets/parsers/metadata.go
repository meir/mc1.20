package parsers

import (
	"bufio"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
)

func init() {
	byteParser := &ByteParser{}
	varintParser := &VarintParser{
		32,
	}

	packets.RegisterParser("entity_metadata", &EnitytMetadataParser{
		reflect.TypeOf(datatypes.EntityMetadata{}),
		byteParser,
		varintParser,
	})
}

type EnitytMetadataParser struct {
	entityMetadataType reflect.Type

	byteParser   *ByteParser
	varintParser *VarintParser
}

func (p *EnitytMetadataParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("metadata", value, p.entityMetadataType); err != nil {
		return err
	}

	metadata := datatypes.EntityMetadata{}

	var b byte
	if err := p.byteParser.Unmarshal(data, reflect.ValueOf(&b)); err != nil {
		return err
	}

	metadata.Index = b

	if b == 0xff {
		value.Set(reflect.ValueOf(metadata))
		return nil
	}

	var t int32
	if err := p.varintParser.Unmarshal(data, reflect.ValueOf(&b)); err != nil {
		return err
	}

	metadata.Type = datatypes.MetadataType(t)

	return nil
}

func (p *EnitytMetadataParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("metadata", value, p.entityMetadataType); err != nil {
		return nil, err
	}

	var metadata datatypes.EntityMetadata
	if value.CanAddr() {
		metadata = value.Addr().Interface().(datatypes.EntityMetadata)
	} else {
		metadata = value.Interface().(datatypes.EntityMetadata)
	}

	if metadata.Index == 0xff {
		return []byte{0xff}, nil
	}

	bytes := []byte{}

	indexBytes, err := p.byteParser.Marshal(reflect.ValueOf(metadata.Index))
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, indexBytes...)

	typeBytes, err := p.varintParser.Marshal(reflect.ValueOf(metadata.Type))
	if err != nil {
		return nil, err
	}
	bytes = append(bytes, typeBytes...)

	return []byte{}, nil
}

type MetadataValueParser struct {
}

func (p *MetadataValueParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	return nil
}

func (p *MetadataValueParser) Marshal(value reflect.Value) ([]byte, error) {
	return []byte{}, nil
}

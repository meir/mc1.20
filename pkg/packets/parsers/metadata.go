package parsers

import (
	"bytes"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/objects"
)

func init() {
	byteParser := &ByteParser{}
	varintParser := &VarintParser{
		32,
	}

	packets.RegisterParser("entity_metadata", &EnitytMetadataParser{
		reflect.TypeOf(objects.EntityMetadata{}),
		byteParser,
		varintParser,
	})
}

type EnitytMetadataParser struct {
	entityMetadataType reflect.Type

	byteParser   *ByteParser
	varintParser *VarintParser
}

func (p *EnitytMetadataParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Ptr,
		}
	}

	if value.Elem().Kind() != reflect.Struct {
		return &packets.ErrInvalidKind{
			value.Elem().Kind(),
			reflect.Struct,
		}
	}

	if value.Elem().Type() != p.entityMetadataType {
		return &packets.ErrInvalidType{
			value.Elem().Type(),
			p.entityMetadataType,
		}
	}

	metadata := objects.EntityMetadata{}

	var b byte
	if err := p.byteParser.Unmarshal(data, reflect.ValueOf(&b)); err != nil {
		return err
	}

	metadata.Index = b

	if b == 0xff {
		value.Elem().Set(reflect.ValueOf(metadata))
		return nil
	}

	var t int32
	if err := p.varintParser.Unmarshal(data, reflect.ValueOf(&b)); err != nil {
		return err
	}

	metadata.Type = objects.MetadataType(t)

	return nil
}

func (p *EnitytMetadataParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Struct,
		}
	}

	if value.Type() != p.entityMetadataType {
		return nil, &packets.ErrInvalidType{
			value.Type(),
			p.entityMetadataType,
		}
	}

	var metadata objects.EntityMetadata
	if value.CanAddr() {
		metadata = value.Addr().Interface().(objects.EntityMetadata)
	} else {
		metadata = value.Interface().(objects.EntityMetadata)
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

func (p *MetadataValueParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	return nil
}

func (p *MetadataValueParser) Marshal(value reflect.Value) ([]byte, error) {
	return []byte{}, nil
}

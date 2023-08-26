package packets

import (
	"bufio"
	"errors"
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

// Unmarshal will accept a bufio.Reader and a pointer to a packet struct and
// will attempt to parse the data from the reader into the packet struct.
// It will do this using reflect and the struct tags.
// For this to work, the packet struct must have the `packet` tag on fields that should be filled.
// The tag should be in the format `packet:"[type],<'optional'>"`. e.g. `packet:"string,optional"`
// The type given in the tag should be a registered parser.
// If the optional tag is present, the parser will first attempt to parse a bool from the reader.
// If the bool is false, the field will be left as the zero value for that type.
// If the bool is true, the parser will attempt to parse the field as normal.
// Parsers are expected to be registered on init.
func Unmarshal(data *bufio.Reader, packet any) error {
	typeOfPacket := reflect.TypeOf(packet)
	valueOfPacket := reflect.ValueOf(packet)

	if typeOfPacket.Kind() != reflect.Ptr {
		return ErrPacketMustBePointer{typeOfPacket.Kind()}
	}

	typeOfPacket = typeOfPacket.Elem()
	valueOfPacket = valueOfPacket.Elem()

	for i := 0; i < typeOfPacket.NumField(); i++ {
		field := typeOfPacket.Field(i)
		tag := strings.Split(field.Tag.Get("packet"), ",")

		if tag[0] == "" {
			continue
		}

		if slices.Contains(tag, "optional") {
			boolParser := parsers["bool"]
			var included bool
			err := boolParser.Unmarshal(data, reflect.ValueOf(&included))
			if err != nil {
				return errors.Join(ErrOptFieldCouldNotBeParsed{
					field:  field.Name,
					kind:   tag[0],
					parser: "bool",
				}, err)
			}

			if !included {
				continue
			}
		}

		if slices.Contains(tag, "array") {
			varintParser := parsers["varint"]

			length := 0
			err := varintParser.Unmarshal(data, reflect.ValueOf(&length))
			if err != nil {
				return errors.Join(ErrFieldCouldNotBeParsed{
					field:  field.Name,
					kind:   tag[0],
					parser: "varint",
				}, err)
			}

			if length == 0 {
				continue
			}

			//create slice of type
			slice := reflect.MakeSlice(field.Type, length, length)

			for i := 0; i < length; i++ {
				// create value of type
				value := reflect.New(field.Type.Elem()).Elem()

				if tag[0] == "struct" {
					err := Unmarshal(data, value.Addr().Interface())
					if err != nil {
						return errors.Join(ErrFieldCouldNotBeParsed{
							field:  field.Name + "[]",
							kind:   tag[0],
							parser: tag[0],
						}, err)
					}

					slice.Index(i).Set(value)
					continue
				}

				err := parsers[tag[0]].Unmarshal(data, value)
				if err != nil {
					return errors.Join(ErrFieldCouldNotBeParsed{
						field:  field.Name,
						kind:   tag[0],
						parser: tag[0],
					}, err)
				}

				slice.Index(i).Set(value)
			}

			valueOfPacket.Field(i).Set(slice)
			continue
		}

		if tag[0] == "struct" {
			value := reflect.New(field.Type).Interface()
			err := Unmarshal(data, value)
			if err != nil {
				return errors.Join(ErrFieldCouldNotBeParsed{
					field:  field.Name,
					kind:   tag[0],
					parser: tag[0],
				}, err)
			}

			valueOfPacket.Field(i).Set(reflect.ValueOf(value).Elem())
			continue
		}

		if parser, ok := parsers[tag[0]]; ok {
			value := valueOfPacket.Field(i)
			err := parser.Unmarshal(data, value)
			if err != nil {
				return errors.Join(ErrFieldCouldNotBeParsed{
					field:  field.Name,
					kind:   tag[0],
					parser: tag[0],
				}, err)
			}
		}
	}

	return nil
}

// Marshal will accept a packet struct and will attempt to marshal it into a byte slice.
// It only parses fields that have the `packet` tag.
func Marshal(packet any) ([]byte, error) {
	typeOfPacket := reflect.TypeOf(packet)
	valueOfPacket := reflect.ValueOf(packet)

	if typeOfPacket.Kind() == reflect.Ptr {
		typeOfPacket = typeOfPacket.Elem()
		valueOfPacket = valueOfPacket.Elem()
	}

	var data []byte

	for i := 0; i < typeOfPacket.NumField(); i++ {
		field := typeOfPacket.Field(i)
		tag := strings.Split(field.Tag.Get("packet"), ",")

		if tag[0] == "" {
			continue
		}

		if slices.Contains(tag, "optional") {
			boolParser := parsers["bool"]
			var included bool

			// check if the field is null
			if valueOfPacket.Field(i).IsNil() {
				included = false
			} else {
				included = true
			}

			boolData, err := boolParser.Marshal(reflect.ValueOf(&included))
			if err != nil {
				return nil, errors.Join(ErrOptFieldCouldNotBeEncoded{
					field:  field.Name,
					kind:   tag[0],
					parser: "bool",
				}, err)
			}

			data = append(data, boolData...)

			if !included {
				continue
			}
		}

		if slices.Contains(tag, "array") {
			varintParser := parsers["varint"]

			length := valueOfPacket.Field(i).Len()
			varintData, err := varintParser.Marshal(reflect.ValueOf(&length))
			if err != nil {
				return nil, errors.Join(ErrFieldCouldNotBeEncoded{
					field:  field.Name,
					kind:   tag[0],
					parser: "varint",
				}, err)
			}

			data = append(data, varintData...)

			array := valueOfPacket.Field(i)

			for i := 0; i < length; i++ {
				value := array.Index(i)

				if tag[0] == "struct" {
					structData, err := Marshal(value.Interface())
					if err != nil {
						return nil, errors.Join(ErrFieldCouldNotBeEncoded{
							field:  field.Name + "[]",
							kind:   tag[0],
							parser: tag[0],
						}, err)
					}

					data = append(data, structData...)
					continue
				}

				packetData, err := parsers[tag[0]].Marshal(value)
				if err != nil {
					return nil, errors.Join(ErrFieldCouldNotBeEncoded{
						field:  field.Name + "[]",
						kind:   tag[0],
						parser: tag[0],
					}, err)
				}

				data = append(data, packetData...)
			}

			continue
		}

		if tag[0] == "struct" {
			structData, err := Marshal(valueOfPacket.Field(i).Interface())
			if err != nil {
				return nil, errors.Join(ErrFieldCouldNotBeEncoded{
					field:  field.Name,
					kind:   tag[0],
					parser: tag[0],
				}, err)
			}

			data = append(data, structData...)
			continue
		}

		if parser, ok := parsers[tag[0]]; ok {
			value := valueOfPacket.Field(i)
			packetData, err := parser.Marshal(value)
			if err != nil {
				return nil, errors.Join(ErrFieldCouldNotBeEncoded{
					field:  field.Name,
					kind:   tag[0],
					parser: tag[0],
				}, err)
			}

			data = append(data, packetData...)
		}
	}

	return data, nil
}

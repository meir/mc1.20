package packets

import (
	"bufio"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Unmarshal(data *bufio.Reader, packet any) error {
	typeOfPacket := reflect.TypeOf(packet)
	valueOfPacket := reflect.ValueOf(packet)

	if typeOfPacket.Kind() != reflect.Ptr {
		return errors.New("packet must be a pointer")
	}

	typeOfPacket = typeOfPacket.Elem()
	valueOfPacket = valueOfPacket.Elem()

	for i := 0; i < typeOfPacket.NumField(); i++ {
		field := typeOfPacket.Field(i)
		tag := strings.Split(field.Tag.Get("packet"), ",")

		if tag[0] == "" {
			continue
		}

		if len(tag) > 2 && tag[2] == "optional" {
			boolParser := parsers["bool"]
			var included bool
			err := boolParser.Unmarshal(data, reflect.ValueOf(&included))
			if err != nil {
				return errors.Join(fmt.Errorf("failed to parse optional field %s", tag[0]), err)
			}

			if !included {
				continue
			}
		}

		if parser, ok := parsers[tag[0]]; ok {
			value := valueOfPacket.Field(i)
			err := parser.Unmarshal(data, value)
			if err != nil {
				return errors.Join(fmt.Errorf("failed to parse field %s", tag[0]), err)
			}
		}
	}

	return nil
}

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

		if len(tag) > 2 && tag[2] == "optional" {
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
				return nil, errors.Join(fmt.Errorf("failed to parse optional field %s", tag[0]), err)
			}

			data = append(data, boolData...)

			if !included {
				continue
			}
		}

		if parser, ok := parsers[tag[0]]; ok {
			value := valueOfPacket.Field(i)
			packetData, err := parser.Marshal(value)
			if err != nil {
				return nil, errors.Join(fmt.Errorf("failed to parse field %s", tag[0]), err)
			}

			data = append(data, packetData...)
		}
	}

	return data, nil
}

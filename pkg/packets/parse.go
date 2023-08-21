package packets

import (
	"bytes"
	"reflect"
	"strings"
)

func Unmarshal(data *bytes.Reader, packet any) error {
	typeOfPacket := reflect.TypeOf(packet)
	valueOfPacket := reflect.ValueOf(packet)

	for i := 0; i < typeOfPacket.NumField(); i++ {
		field := typeOfPacket.Field(i)
		tag := field.Tag.Get("packet")

		if tag == "" {
			continue
		}

		optional := strings.Split(tag, ",")[1] == "optional"
		tag = strings.Split(tag, ",")[0]

		if optional {
			boolParser := parsers["bool"]
			var included bool
			err := boolParser.Unmarshal(data, reflect.ValueOf(&included))
			if err != nil {
				return err
			}

			if !included {
				continue
			}
		}

		if parser, ok := parsers[tag]; ok {
			value := valueOfPacket.Field(i)
			if optional && value.IsNil() {
				continue
			}

			err := parser.Unmarshal(data, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

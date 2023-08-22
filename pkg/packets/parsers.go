package packets

import (
	"bufio"
	"reflect"
)

type FieldParser interface {
	Marshal(reflect.Value) ([]byte, error)
	Unmarshal(*bufio.Reader, reflect.Value) error
}

var parsers = map[string]FieldParser{}

func RegisterParser(name string, parser FieldParser) {
	parsers[name] = parser
}

func GetParser(name string) FieldParser {
	return parsers[name]
}

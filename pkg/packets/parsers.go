// Package is a simple package for parsing and serializing tcp/udp packets.
// It is primarily made for the minecraft protocol https://wiki.vg/Protocol
// But there isnt any reason it couldnt be used for other protocols or even other data parsing.
// It works by using reflection to parse and serialize structs and get parsers for the fields.
//
// The parsers themselves need to self register by using the RegisterParser function.
package packets

import (
	"bufio"
	"reflect"
)

var parsers = map[string]FieldParser{}

// FieldParser is an interface for parsing and serializing fields
// a field would have a `packet:"type"` tag, the FieldParser would be specific to that type
// and would be used to parse and serialize the field
type FieldParser interface {
	Marshal(reflect.Value) ([]byte, error)
	Unmarshal(*bufio.Reader, reflect.Value) error
}

// RegisterParser registers a parser for a type
func RegisterParser(name string, parser FieldParser) {
	parsers[name] = parser
}

// GetParser gets a parser for a type
func GetParser(name string) FieldParser {
	if parser, ok := parsers[name]; ok {
		return parser
	}
	return nil
}

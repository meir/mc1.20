package parsers

import (
	"bytes"
	"reflect"
	"regexp"

	"github.com/meir/mc1.20/pkg/packets"
)

func init() {
	packets.RegisterParser("identifier", &IdentifierParser{
		stringParser: &StringParser{
			varintParser: &VarintParser{
				bits: 32,
			},
		},
		rule: regexp.MustCompile(`^[a-z0-9.-_]+:[a-z0-9.-_/]+$`),
	})
}

type IdentifierParser struct {
	stringParser *StringParser
	rule         *regexp.Regexp
}

func (p *IdentifierParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return &packets.ErrInvalidKind{
			Kind:   value.Kind(),
			Wanted: reflect.Ptr,
		}
	}

	if value.Elem().Kind() != reflect.String {
		return &packets.ErrInvalidKind{
			Kind:   value.Elem().Kind(),
			Wanted: reflect.String,
		}
	}

	var s string
	if err := p.stringParser.Unmarshal(data, reflect.ValueOf(&s)); err != nil {
		return err
	}

	if !p.rule.MatchString(s) {
		return &packets.ErrInvalidValue{
			Value:  s,
			Reason: "does not match the rule: " + p.rule.String(),
		}
	}

	value.Elem().SetString(s)

	return nil
}

func (p *IdentifierParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() != reflect.String {
		return nil, &packets.ErrInvalidKind{
			Kind:   value.Kind(),
			Wanted: reflect.String,
		}
	}

	if !p.rule.MatchString(value.String()) {
		return nil, &packets.ErrInvalidValue{
			Value:  value.String(),
			Reason: "does not match the rule: " + p.rule.String(),
		}
	}

	return p.stringParser.Marshal(value)
}

package parsers

import (
	"bufio"
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

func (p *IdentifierParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectKind("identifier", value, reflect.String); err != nil {
		return err
	}

	var s string
	if err := p.stringParser.Unmarshal(data, reflect.ValueOf(&s)); err != nil {
		return err
	}

	if !p.rule.MatchString(s) {
		return &ErrInvalidValue{
			Value:  s,
			Reason: "does not match the rule: " + p.rule.String(),
		}
	}

	value.SetString(s)

	return nil
}

func (p *IdentifierParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectKind("identifier", value, reflect.String); err != nil {
		return nil, err
	}

	if !p.rule.MatchString(value.String()) {
		return nil, &ErrInvalidValue{
			Value:  value.String(),
			Reason: "does not match the rule: " + p.rule.String(),
		}
	}

	return p.stringParser.Marshal(value)
}

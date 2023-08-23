package parsers

import (
	"bufio"
	"encoding/json"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
)

func init() {
	packets.RegisterParser("chat", &ChatParser{
		&StringParser{
			&VarintParser{
				bits: 32,
			},
		},
		reflect.TypeOf(datatypes.Chat{}),
	})
}

type ChatParser struct {
	stringParser *StringParser
	chatType     reflect.Type
}

func (p *ChatParser) Unmarshal(data *bufio.Reader, value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("chat", value, p.chatType); err != nil {
		return err
	}

	var s string
	if err := p.stringParser.Unmarshal(data, reflect.ValueOf(&s)); err != nil {
		return err
	}

	var chat datatypes.Chat
	if err := json.Unmarshal([]byte(s), &chat); err != nil {
		return err
	}

	value.Set(reflect.ValueOf(chat))

	return nil
}

func (p *ChatParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if err := expectType("chat", value, p.chatType); err != nil {
		return nil, err
	}

	var chat datatypes.Chat
	if value.CanAddr() {
		chat = value.Addr().Interface().(datatypes.Chat)
	} else {
		chat = value.Interface().(datatypes.Chat)
	}

	b, err := json.Marshal(chat)
	if err != nil {
		return nil, err
	}

	return p.stringParser.Marshal(reflect.ValueOf(string(b)))
}

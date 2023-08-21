package parsers

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/meir/mc1.20/pkg/packets"
	"github.com/meir/mc1.20/pkg/packets/objects"
)

func init() {
	packets.RegisterParser("chat", &ChatParser{
		&StringParser{
			&VarintParser{
				bits: 32,
			},
		},
		reflect.TypeOf(objects.Chat{}),
	})
}

type ChatParser struct {
	stringParser *StringParser
	chatType     reflect.Type
}

func (p *ChatParser) Unmarshal(data *bytes.Reader, value reflect.Value) error {
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

	if value.Elem().Type() != p.chatType {
		return &packets.ErrInvalidType{
			value.Elem().Type(),
			p.chatType,
		}
	}

	var s string
	if err := p.stringParser.Unmarshal(data, reflect.ValueOf(&s)); err != nil {
		return err
	}

	var chat objects.Chat
	if err := json.Unmarshal([]byte(s), &chat); err != nil {
		return err
	}

	value.Elem().Set(reflect.ValueOf(chat))

	return nil
}

func (p *ChatParser) Marshal(value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, &packets.ErrInvalidKind{
			value.Kind(),
			reflect.Struct,
		}
	}

	if value.Type() != p.chatType {
		return nil, &packets.ErrInvalidType{
			value.Type(),
			p.chatType,
		}
	}

	var chat objects.Chat
	if value.CanAddr() {
		chat = value.Addr().Interface().(objects.Chat)
	} else {
		chat = value.Interface().(objects.Chat)
	}

	b, err := json.Marshal(chat)
	if err != nil {
		return nil, err
	}

	return p.stringParser.Marshal(reflect.ValueOf(string(b)))
}

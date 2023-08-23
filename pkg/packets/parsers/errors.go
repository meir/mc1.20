package parsers

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/slices"
)

func expectKind(parser string, v reflect.Value, kinds ...reflect.Kind) error {
	if slices.Contains(kinds, v.Kind()) {
		return nil
	}

	return &ErrInvalidKind{
		parser,
		v.Kind(),
		kinds,
	}
}

var ErrVarintTooBig = fmt.Errorf("varint is too big")

type ErrInvalidKind struct {
	Parser string
	Kind   reflect.Kind
	Wanted []reflect.Kind
}

func (e *ErrInvalidKind) Error() string {
	return fmt.Sprintf("invalid kind on parser for %s: %s, wanted one of: %s", e.Parser, e.Kind, e.Wanted)
}

func expectType(parser string, v reflect.Value, types ...reflect.Type) error {
	if slices.Contains(types, v.Type()) {
		return nil
	}

	return &ErrInvalidType{
		parser,
		v.Type(),
		types,
	}
}

type ErrInvalidType struct {
	Parser string
	Type   reflect.Type
	Wanted []reflect.Type
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("invalid type on parser for %s: %s, wanted one of: %s", e.Parser, e.Type, e.Wanted)
}

type ErrInvalidLength struct {
	Length int
}

func (e *ErrInvalidLength) Error() string {
	return fmt.Sprintf("invalid length: %d", e.Length)
}

type ErrInvalidValue struct {
	Value  interface{}
	Reason string
}

func (e *ErrInvalidValue) Error() string {
	return fmt.Sprintf("invalid value: %v, reason: %s", e.Value, e.Reason)
}

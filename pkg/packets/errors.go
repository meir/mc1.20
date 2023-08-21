package packets

import (
	"fmt"
	"reflect"
)

var ErrVarintTooBig = fmt.Errorf("varint is too big")

type ErrInvalidKind struct {
	Kind   reflect.Kind
	Wanted reflect.Kind
}

func (e *ErrInvalidKind) Error() string {
	return fmt.Sprintf("invalid kind: %s, wanted: %s", e.Kind, e.Wanted)
}

type ErrInvalidType struct {
	Type   reflect.Type
	Wanted reflect.Type
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("invalid type: %s, wanted: %s", e.Type, e.Wanted)
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

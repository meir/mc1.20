package packets

import (
	"fmt"
	"reflect"
)

// ErrPacketMustBePointer is returned when the packet passed to Unmarshal is not a pointer.
type ErrPacketMustBePointer struct {
	kind reflect.Kind
}

func (e ErrPacketMustBePointer) Error() string {
	return fmt.Sprintf("packet must be a pointer, got %s", e.kind)
}

// ErrOptFieldCouldNotBeParsed is returned when an optional field could not be parsed.
type ErrOptFieldCouldNotBeParsed struct {
	field  string
	kind   string
	parser string
}

func (e ErrOptFieldCouldNotBeParsed) Error() string {
	return fmt.Sprintf("optional field %s (%s) could not be parsed using %s parser", e.field, e.kind, e.parser)
}

// ErrFieldCouldNotBeParsed is returned when a field could not be parsed.
type ErrFieldCouldNotBeParsed struct {
	field  string
	kind   string
	parser string
}

func (e ErrFieldCouldNotBeParsed) Error() string {
	return fmt.Sprintf("field %s (%s) could not be parsed using %s parser", e.field, e.kind, e.parser)
}

// ErrOptFieldCouldNotBeEncoded is returned when an optional field could not be encoded.
type ErrOptFieldCouldNotBeEncoded struct {
	field  string
	kind   string
	parser string
}

func (e ErrOptFieldCouldNotBeEncoded) Error() string {
	return fmt.Sprintf("optional field %s (%s) could not be encoded using %s parser", e.field, e.kind, e.parser)
}

// ErrFieldCouldNotBeEncoded is returned when a field could not be encoded.
type ErrFieldCouldNotBeEncoded struct {
	field  string
	kind   string
	parser string
}

func (e ErrFieldCouldNotBeEncoded) Error() string {
	return fmt.Sprintf("field %s (%s) could not be encoded using %s parser", e.field, e.kind, e.parser)
}

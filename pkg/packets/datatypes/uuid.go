package datatypes

import (
	"crypto/rand"
	"encoding/hex"
)

type UUID [16]byte

// for xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// aka 8-4-4-4-12 where each x is 2 hex digits for one byte
var uuidMap = [16]int{
	0, 2, 4, 6,
	9, 11,
	14, 16,
	19, 21,
	24, 26, 28, 30, 32, 34,
}

// ErrInvalidUUID is returned when a UUID is invalid
type ErrInvalidUUID string

// Error returns the error message for ErrInvalidUUID
func (e ErrInvalidUUID) Error() string {
	return "invalid uuid: " + string(e)
}

// NewUUID generates a new UUID
func NewUUID() UUID {
	r := rand.Reader
	var u UUID
	r.Read(u[:])
	return u
}

// Parse parses a UUID from a string
// It returns an error if the string is not a valid UUID
// It is not case sensitive
// The UUID must be in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// where x is a hex digit
func (u UUID) Parse(s string) (UUID, error) {
	//must be xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	//where x is a hex digit
	if len(s) != 36 {
		return u, ErrInvalidUUID(s)
	}

	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return u, ErrInvalidUUID(s)
	}

	for i, x := range uuidMap {
		h := s[x : x+2]
		b, err := hex.DecodeString(h)
		if err != nil {
			return u, ErrInvalidUUID(s)
		}

		u[i] = b[0]
	}

	return u, nil
}

// String returns the UUID as a string in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (u UUID) String() string {
	var buf []byte = make([]byte, 36)

	hex.Encode(buf, u[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

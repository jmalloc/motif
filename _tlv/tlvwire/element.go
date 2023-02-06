package tlvwire

import "bytes"

// Element represents a TLV element.
type Element struct {
	Tag   Tag
	Value Value
}

// Marshal returns the TLV encoded form of the given element.
func Marshal(e Element) ([]byte, error) {
	w := &bytes.Buffer{}

	w.WriteByte(
		byte(e.Tag.Form<<5) | byte(e.Value.Type),
	)

	if err := marshalTag(w, e.Tag); err != nil {
		return nil, err
	}

	if err := marshalValue(w, e.Value); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

const (
	// ElementTypeMask is the bitmask for the element type enumeration within
	// the "Control Octet".
	ElementTypeMask byte = 0b000_11111

	// TagFormMask is the bitmask for the tag form enumeration within the
	// "Control Octet".
	TagFormMask byte = 0b111_00000

	// EndOfContainer is the value of the "Control Octet" that indicates the
	// end of a container element.
	EndOfContainer byte = 0b000_11000
)

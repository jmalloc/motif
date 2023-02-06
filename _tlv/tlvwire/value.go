package tlvwire

import (
	"bytes"
	"fmt"
)

type Value interface {
	acceptVisitor(ValueVisitor) error
}

type ValueVisitor interface {
}

type XValue struct {
	Type        Type
	Signed      int64
	Unsigned    uint64
	Boolean     bool
	Float       float64
	UTF8String  string
	OctetString []byte
	Members     []Element
}

type Type byte

const (
	Signed1Type      Type = 0b00000
	Signed2Type      Type = 0b00001
	Signed4Type      Type = 0b00010
	Signed8Type      Type = 0b00011
	Unsigned1Type    Type = 0b00100
	Unsigned2Type    Type = 0b00101
	Unsigned4Type    Type = 0b00110
	Unsigned8Type    Type = 0b00111
	BooleanFalseType Type = 0b01000
	BooleanTrueType  Type = 0b01001
	Float4Type       Type = 0b01010
	Float8Type       Type = 0b01011
	UTF8String1Type  Type = 0b01100
	UTF8String2Type  Type = 0b01101
	UTF8String4Type  Type = 0b01110
	UTF8String8Type  Type = 0b01111
	OctetString1Type Type = 0b10000
	OctetString2Type Type = 0b10001
	OctetString4Type Type = 0b10010
	OctetString8Type Type = 0b10011
	NullType         Type = 0b10100
	StructureType    Type = 0b10101
	ArrayType        Type = 0b10110
	ListType         Type = 0b10111
)

func marshalValue(w *bytes.Buffer, v Value) error {
	switch v.Type {
	case Signed1Type:
		return marshalSigned1(w, v)
	case Signed2Type:
		return marshalSigned2(w, v)
	case Signed4Type:
		return marshalSigned4(w, v)
	case Signed8Type:
		return marshalSigned8(w, v)
	case Unsigned1Type:
		return marshalUnsigned1(w, v)
	case Unsigned2Type:
		return marshalUnsigned2(w, v)
	case Unsigned4Type:
		return marshalUnsigned4(w, v)
	case Unsigned8Type:
		return marshalUnsigned8(w, v)
	case BooleanFalseType:
		return marshalBooleanFalse(w, v)
	case BooleanTrueType:
		return marshalBooleanTrue(w, v)
	case Float4Type:
		return marshalFloat4(w, v)
	case Float8Type:
		return marshalFloat8(w, v)
	case UTF8String1Type:
		return marshalUTF8String1(w, v)
	case UTF8String2Type:
		return marshalUTF8String2(w, v)
	case UTF8String4Type:
		return marshalUTF8String4(w, v)
	case UTF8String8Type:
		return marshalUTF8String8(w, v)
	case OctetString1Type:
		return marshalOctetString1(w, v)
	case OctetString2Type:
		return marshalOctetString2(w, v)
	case OctetString4Type:
		return marshalOctetString4(w, v)
	case OctetString8Type:
		return marshalOctetString8(w, v)
	case NullType:
		return nil
	case StructureType:
		return marshalStructure(w, v)
	case ArrayType:
		return marshalArray(w, v)
	case ListType:
		return marshalList(w, v)
	default:
		return fmt.Errorf("unrecognized element type (%b)", v.Type)
	}
}

func marshalSigned1(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalSigned2(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalSigned4(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalSigned8(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUnsigned1(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUnsigned2(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUnsigned4(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUnsigned8(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalBooleanFalse(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalBooleanTrue(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalFloat4(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalFloat8(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUTF8String1(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUTF8String2(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUTF8String4(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalUTF8String8(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalOctetString1(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalOctetString2(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalOctetString4(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalOctetString8(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalStructure(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalArray(w *bytes.Buffer, v Value) error {
	panic("ni")
}

func marshalList(w *bytes.Buffer, v Value) error {
	panic("ni")
}

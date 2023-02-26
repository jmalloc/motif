package tlv

import (
	"bytes"
	"fmt"
)

// Value is an interface for a TLV value.
type Value interface {
	acceptVisitor(ValueVisitor) error
}

// ValueVisitor is an interface for visiting TLV values.
type ValueVisitor interface {
	VisitSigned1(Signed1) error
	VisitSigned2(Signed2) error
	VisitSigned4(Signed4) error
	VisitSigned8(Signed8) error
	VisitUnsigned1(Unsigned1) error
	VisitUnsigned2(Unsigned2) error
	VisitUnsigned4(Unsigned4) error
	VisitUnsigned8(Unsigned8) error
	VisitBool(Bool) error
	VisitSingle(Single) error
	VisitDouble(Double) error
	VisitNull() error
	VisitStruct(Struct) error
	VisitArray(Array) error
	VisitList(List) error
	VisitUTF8String1(UTF8String1) error
	VisitUTF8String2(UTF8String2) error
	VisitUTF8String4(UTF8String4) error
	VisitUTF8String8(UTF8String8) error
	VisitOctetString1(OctetString1) error
	VisitOctetString2(OctetString2) error
	VisitOctetString4(OctetString4) error
	VisitOctetString8(OctetString8) error
}

// VisitValue visits a value with a visitor.
func VisitValue(v Value, vis ValueVisitor) error {
	return v.acceptVisitor(vis)
}

func unmarshalValue(r *bytes.Reader, c byte) (Value, error) {
	switch c & typeMask {
	case signed1Type:
		return readInt[Signed1](r)
	case signed2Type:
		return readInt[Signed2](r)
	case signed4Type:
		return readInt[Signed4](r)
	case signed8Type:
		return readInt[Signed8](r)
	case unsigned1Type:
		return readInt[Unsigned1](r)
	case unsigned2Type:
		return readInt[Unsigned2](r)
	case unsigned4Type:
		return readInt[Unsigned4](r)
	case unsigned8Type:
		return readInt[Unsigned8](r)
	case boolFalseType:
		return False, nil
	case boolTrueType:
		return True, nil
	case singleType:
		return unmarshalSingle(r)
	case doubleType:
		return unmarshalDouble(r)
	case nullType:
		return Null, nil
	case structType:
		return unmarshalStruct(r)
	case arrayType:
		return unmarshalArray(r)
	case listType:
		return unmarshalList(r)
	case utf8String1Type:
		return readString[uint8, UTF8String1](r)
	case utf8String2Type:
		return readString[uint16, UTF8String2](r)
	case utf8String4Type:
		return readString[uint32, UTF8String4](r)
	case utf8String8Type:
		return readString[uint64, UTF8String8](r)
	case octetString1Type:
		return readString[uint8, OctetString1](r)
	case octetString2Type:
		return readString[uint16, OctetString2](r)
	case octetString4Type:
		return readString[uint32, OctetString4](r)
	case octetString8Type:
		return readString[uint64, OctetString8](r)
	default:
		return nil, fmt.Errorf("unrecognized type (%x)", c&typeMask)
	}
}

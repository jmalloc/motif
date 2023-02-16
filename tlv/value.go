package tlv

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

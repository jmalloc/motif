package tlv

// Value is an interface for a TLV value.
type Value interface {
	AcceptVisitor(ValueVisitor) error
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
	VisitFloat4(Float4) error
	VisitFloat8(Float8) error
	VisitNull() error
	VisitStruct(Struct) error
	VisitArray(Array) error
	VisitList(List) error
	VisitString1(String1) error
	VisitString2(String2) error
	VisitString4(String4) error
	VisitString8(String8) error
	VisitBytes1(Bytes1) error
	VisitBytes2(Bytes2) error
	VisitBytes4(Bytes4) error
	VisitBytes8(Bytes8) error
}

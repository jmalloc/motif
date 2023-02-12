package tlv

// Value is an interface for a TLV value.
type Value interface {
	AcceptVisitor(ValueVisitor)
}

// ValueVisitor is an interface for visiting TLV values.
type ValueVisitor interface {
	VisitSigned1(Signed1)
	VisitSigned2(Signed2)
	VisitSigned4(Signed4)
	VisitSigned8(Signed8)
	VisitUnsigned1(Unsigned1)
	VisitUnsigned2(Unsigned2)
	VisitUnsigned4(Unsigned4)
	VisitUnsigned8(Unsigned8)
	VisitBool(Bool)
	VisitFloat4(Float4)
	VisitFloat8(Float8)
	VisitNull()
	VisitStruct(Struct)
	VisitString1(String1)
	VisitString2(String2)
	VisitString4(String4)
	VisitString8(String8)
	VisitBytes1(Bytes1)
	VisitBytes2(Bytes2)
	VisitBytes4(Bytes4)
	VisitBytes8(Bytes8)
}

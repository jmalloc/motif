package tlv

// Value is an interface for a TLV value.
//
// A value is also an "anonymous" element, that is, an element without a tag.
type Value interface {
	Element

	AcceptValueVisitor(ValueVisitor)
}

// ValueVisitor is an interface for visiting TLV values.
type ValueVisitor interface {
	VisitSigned1(Signed1)
	VisitSigned2(Signed2)
	VisitSigned4(Signed4)
	VisitSigned8(Signed8)
	VisitNull()
	VisitStruct(Struct)
	VisitString1(String1)
	VisitString2(String2)
	VisitString4(String4)
	VisitString8(String8)
}

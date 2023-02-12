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
	VisitNull()
	VisitString1(s string)
	VisitString2(s string)
	VisitString4(s string)
	VisitString8(s string)
}

package tlv

// Struct is a TLV structure element.
type Struct []StructMember

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Struct) AcceptVisitor(v ValueVisitor) error {
	return v.VisitStruct(s)
}

// StructMember is an element that is a member of a structure.
type StructMember struct {
	T NonAnonymousTag
	V Value
}

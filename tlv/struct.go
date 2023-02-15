package tlv

// Struct is a TLV structure element.
type Struct []StructMember

// Members returns the elements that are members of the structure.
func (s Struct) Members() []Element {
	elements := make([]Element, len(s))
	for i, m := range s {
		elements[i] = m
	}
	return elements
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Struct) AcceptVisitor(v ValueVisitor) {
	v.VisitStruct(s)
}

// StructMember is an element that is a member of a structure.
type StructMember struct {
	Tag   NonAnonymousTag
	Value Value
}

// Components returns the tag and value of the element.
func (m StructMember) Components() (Tag, Value) {
	return m.Tag, m.Value
}

package tlv

// Array is a TLV array element.
type Array []ArrayMember

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (a Array) AcceptVisitor(v ValueVisitor) {
	v.VisitArray(a)
}

// ArrayMember is an element that is a member of an array.
type ArrayMember struct {
	V Value
}

// Tag returns the element's tag.
func (m ArrayMember) Tag() Tag {
	return AnonymousTag
}

// Value returns the element's value.
func (m ArrayMember) Value() Value {
	return m.V
}

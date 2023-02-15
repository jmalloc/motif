package tlv

// Array is a TLV array element.
type Array []ArrayMember

// NewArray returns an array with the given values as members.
func NewArray(values ...Value) Array {
	a := make(Array, len(values))
	for i, v := range values {
		a[i] = ArrayMember{v}
	}
	return a
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (a Array) AcceptVisitor(v ValueVisitor) {
	v.VisitArray(a)
}

// ArrayMember is an element that is a member of an array.
type ArrayMember struct {
	Value Value
}

// Components returns the tag and value of the element.
func (m ArrayMember) Components() (Tag, Value) {
	return AnonymousTag, m.Value
}

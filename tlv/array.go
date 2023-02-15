package tlv

// Array is a TLV array element.
type Array []Value

// Members returns the elements that are members of the array.
func (a Array) Members() []Element {
	elements := make([]Element, len(a))
	for i, v := range a {
		elements[i] = arrayMember{v}
	}
	return elements
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (a Array) AcceptVisitor(v ValueVisitor) {
	v.VisitArray(a)
}

type arrayMember struct {
	Value Value
}

func (e arrayMember) Components() (Tag, Value) {
	return AnonymousTag, e.Value
}

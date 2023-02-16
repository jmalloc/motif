package tlv

// Array is a TLV array element.
type Array []Value

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (a Array) AcceptVisitor(v ValueVisitor) error {
	return v.VisitArray(a)
}

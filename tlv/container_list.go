package tlv

// List is a TLV list element.
type List []ListMember

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (l List) AcceptVisitor(v ValueVisitor) error {
	return v.VisitList(l)
}

// ListMember is an element that is a member of a list.
type ListMember struct {
	T Tag
	V Value
}

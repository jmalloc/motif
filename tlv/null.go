package tlv

// Null is the TLV null value.
const Null null = 0

type null uint8

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (null) AcceptVisitor(v ValueVisitor) error {
	return v.VisitNull()
}

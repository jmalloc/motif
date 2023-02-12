package tlv

// Null is the TLV null value.
var Null Value = null{}

type null struct{}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (null) AcceptVisitor(v ValueVisitor) {
	v.VisitNull()
}

func (m marshaler) VisitNull() {
	m.WriteControl(nullType)
}

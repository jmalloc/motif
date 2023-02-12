package tlv

const (
	// False is the TLV false value.
	False Bool = false

	// True is the TLV true value.
	True Bool = true
)

// Bool is a TLV boolean value.
type Bool bool

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bool) AcceptVisitor(v ValueVisitor) {
	v.VisitBool(b)
}

func (m marshaler) VisitBool(b Bool) {
	if b {
		m.WriteControl(boolTrueType)
	} else {
		m.WriteControl(boolFalseType)
	}
}

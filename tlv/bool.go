package tlv

type (
	// Bool is a TLV boolean value.
	Bool bool
)

const (
	// False is the TLV false value.
	False Bool = false

	// True is the TLV true value.
	True Bool = true
)

func (v Bool) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitBool(v)
}

const (
	boolFalseType = 0b000_01000
	boolTrueType  = 0b000_01001
)

func (w *controlWriter) VisitBool(v Bool) error {
	if v {
		return w.write(boolTrueType)
	}
	return w.write(boolFalseType)
}

func (w *payloadWriter) VisitBool(v Bool) error {
	return nil
}

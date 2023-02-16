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

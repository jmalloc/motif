package tlv

type (
	// Unsigned1 is a signed 1 octet unsigned integer.
	Unsigned1 uint8

	// Unsigned2 is a signed 2 octet unsigned integer.
	Unsigned2 uint16

	// Unsigned4 is a signed 4 octet unsigned integer.
	Unsigned4 uint32

	// Unsigned8 is a signed 8 octet unsigned integer.
	Unsigned8 uint64
)

func (v Unsigned1) acceptVisitor(vis ValueVisitor) error { return vis.VisitUnsigned1(v) }
func (v Unsigned2) acceptVisitor(vis ValueVisitor) error { return vis.VisitUnsigned2(v) }
func (v Unsigned4) acceptVisitor(vis ValueVisitor) error { return vis.VisitUnsigned4(v) }
func (v Unsigned8) acceptVisitor(vis ValueVisitor) error { return vis.VisitUnsigned8(v) }

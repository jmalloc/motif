package tlv

type (
	// Signed1 is a signed 1 octet signed integer.
	Signed1 int8

	// Signed2 is a signed 2 octet signed integer.
	Signed2 int16

	// Signed4 is a signed 4 octet signed integer.
	Signed4 int32

	// Signed8 is a signed 8 octet signed integer.
	Signed8 int64
)

func (v Signed1) acceptVisitor(vis ValueVisitor) error { return vis.VisitSigned1(v) }
func (v Signed2) acceptVisitor(vis ValueVisitor) error { return vis.VisitSigned2(v) }
func (v Signed4) acceptVisitor(vis ValueVisitor) error { return vis.VisitSigned4(v) }
func (v Signed8) acceptVisitor(vis ValueVisitor) error { return vis.VisitSigned8(v) }

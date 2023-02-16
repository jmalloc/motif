package tlv

type (
	// Float4 is a single-precision (4 octet) TLV floating-point value.
	Float4 float32

	// Float8 is a double-precision (8 octet) TLV floating-point value.
	Float8 float64
)

func (v Float4) acceptVisitor(vis ValueVisitor) error { return vis.VisitFloat4(v) }
func (v Float8) acceptVisitor(vis ValueVisitor) error { return vis.VisitFloat8(v) }

package tlv

type (
	// Single is a single-precision (4 octet) TLV floating-point value.
	Single float32

	// Double is a double-precision (8 octet) TLV floating-point value.
	Double float64
)

func (v Single) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitSingle(v)
}

func (v Double) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitDouble(v)
}

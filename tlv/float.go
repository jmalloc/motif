package tlv

type (
	// Float4 is a single-precision (4 octet) TLV floating-point value.
	Float4 float32

	// Float8 is a double-precision (8 octet) TLV floating-point value.
	Float8 float64
)

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (f Float4) AcceptVisitor(v ValueVisitor) error {
	return v.VisitFloat4(f)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (f Float8) AcceptVisitor(v ValueVisitor) error {
	return v.VisitFloat8(f)
}

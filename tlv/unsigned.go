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

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned1) AcceptVisitor(v ValueVisitor) error {
	return v.VisitUnsigned1(u)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned2) AcceptVisitor(v ValueVisitor) error {
	return v.VisitUnsigned2(u)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned4) AcceptVisitor(v ValueVisitor) error {
	return v.VisitUnsigned4(u)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned8) AcceptVisitor(v ValueVisitor) error {
	return v.VisitUnsigned8(u)
}

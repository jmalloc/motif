package tlv

type (
	// Bytes1 is an octet-string with a 1 octet length.
	Bytes1 []byte

	// Bytes2 is an octet-string with a 2 octet length.
	Bytes2 []byte

	// Bytes4 is an octet-string with a 4 octet length.
	Bytes4 []byte

	// Bytes8 is an octet-string with an 8 octet length.
	Bytes8 []byte
)

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes1) AcceptVisitor(v ValueVisitor) error {
	return v.VisitBytes1(b)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes2) AcceptVisitor(v ValueVisitor) error {
	return v.VisitBytes2(b)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes4) AcceptVisitor(v ValueVisitor) error {
	return v.VisitBytes4(b)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes8) AcceptVisitor(v ValueVisitor) error {
	return v.VisitBytes8(b)
}

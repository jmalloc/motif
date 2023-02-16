package tlv

type (
	// String1 is a UTF-8 string with a 1 octet length.
	String1 string

	// String2 is a UTF-8 string with a 2 octet length.
	String2 string

	// String4 is a UTF-8 string with a 4 octet length.
	String4 string

	// String8 is a UTF-8 string with an 8 octet length.
	String8 string
)

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String1) AcceptVisitor(v ValueVisitor) error {
	return v.VisitString1(s)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String2) AcceptVisitor(v ValueVisitor) error {
	return v.VisitString2(s)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String4) AcceptVisitor(v ValueVisitor) error {
	return v.VisitString4(s)
}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String8) AcceptVisitor(v ValueVisitor) error {
	return v.VisitString8(s)
}

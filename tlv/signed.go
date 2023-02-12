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

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Signed1) AcceptVisitor(v ValueVisitor) { v.VisitSigned1(s) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Signed2) AcceptVisitor(v ValueVisitor) { v.VisitSigned2(s) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Signed4) AcceptVisitor(v ValueVisitor) { v.VisitSigned4(s) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Signed8) AcceptVisitor(v ValueVisitor) { v.VisitSigned8(s) }

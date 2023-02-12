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

func (m marshaler) VisitSigned1(s Signed1) {
	m.WriteControl(signed1Type)
	m.WriteByte(byte(s))
}

func (m marshaler) VisitSigned2(s Signed2) {
	m.WriteControl(signed2Type)
	m.Grow(2)
	m.WriteByte(byte(s))
	m.WriteByte(byte(s >> 8))
}

func (m marshaler) VisitSigned4(s Signed4) {
	m.WriteControl(signed4Type)
	m.Grow(4)
	m.WriteByte(byte(s))
	m.WriteByte(byte(s >> 8))
	m.WriteByte(byte(s >> 16))
	m.WriteByte(byte(s >> 24))
}

func (m marshaler) VisitSigned8(s Signed8) {
	m.WriteControl(signed8Type)
	m.Grow(8)
	m.WriteByte(byte(s))
	m.WriteByte(byte(s >> 8))
	m.WriteByte(byte(s >> 16))
	m.WriteByte(byte(s >> 24))
	m.WriteByte(byte(s >> 32))
	m.WriteByte(byte(s >> 40))
	m.WriteByte(byte(s >> 48))
	m.WriteByte(byte(s >> 56))
}

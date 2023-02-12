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
func (u Unsigned1) AcceptVisitor(v ValueVisitor) { v.VisitUnsigned1(u) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned2) AcceptVisitor(v ValueVisitor) { v.VisitUnsigned2(u) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned4) AcceptVisitor(v ValueVisitor) { v.VisitUnsigned4(u) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (u Unsigned8) AcceptVisitor(v ValueVisitor) { v.VisitUnsigned8(u) }

func (m marshaler) VisitUnsigned1(u Unsigned1) {
	m.WriteControl(unsigned1Type)
	m.WriteByte(byte(u))
}

func (m marshaler) VisitUnsigned2(u Unsigned2) {
	m.WriteControl(unsigned2Type)
	m.Grow(2)
	m.WriteByte(byte(u))
	m.WriteByte(byte(u >> 8))
}

func (m marshaler) VisitUnsigned4(u Unsigned4) {
	m.WriteControl(unsigned4Type)
	m.Grow(4)
	m.WriteByte(byte(u))
	m.WriteByte(byte(u >> 8))
	m.WriteByte(byte(u >> 16))
	m.WriteByte(byte(u >> 24))
}

func (m marshaler) VisitUnsigned8(u Unsigned8) {
	m.WriteControl(unsigned8Type)
	m.Grow(8)
	m.WriteByte(byte(u))
	m.WriteByte(byte(u >> 8))
	m.WriteByte(byte(u >> 16))
	m.WriteByte(byte(u >> 24))
	m.WriteByte(byte(u >> 32))
	m.WriteByte(byte(u >> 40))
	m.WriteByte(byte(u >> 48))
	m.WriteByte(byte(u >> 56))
}

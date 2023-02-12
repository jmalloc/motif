package tlv

type (
	// String1 is a string with a 1 octet length.
	String1 string

	// String2 is a string with a 2 octet length.
	String2 string

	// String4 is a string with a 4 octet length.
	String4 string

	// String8 is a string with an 8 octet length.
	String8 string
)

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String1) AcceptVisitor(v ValueVisitor) { v.VisitString1(s) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String2) AcceptVisitor(v ValueVisitor) { v.VisitString2(s) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String4) AcceptVisitor(v ValueVisitor) { v.VisitString4(s) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s String8) AcceptVisitor(v ValueVisitor) { v.VisitString8(s) }

func (m marshaler) VisitString1(s String1) {
	m.WriteControl(string1Type)
	n := len(s)
	m.Grow(1 + n)
	m.WriteByte(byte(n))
	m.WriteString(string(s))
}

func (m marshaler) VisitString2(s String2) {
	m.WriteControl(string2Type)
	n := len(s)
	m.Grow(2 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteString(string(s))
}

func (m marshaler) VisitString4(s String4) {
	m.WriteControl(string4Type)
	n := len(s)
	m.Grow(4 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteByte(byte(n >> 16))
	m.WriteByte(byte(n >> 24))
	m.WriteString(string(s))
}

func (m marshaler) VisitString8(s String8) {
	m.WriteControl(string8Type)
	n := len(s)
	m.Grow(8 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteByte(byte(n >> 16))
	m.WriteByte(byte(n >> 24))
	m.WriteByte(byte(n >> 32))
	m.WriteByte(byte(n >> 40))
	m.WriteByte(byte(n >> 48))
	m.WriteByte(byte(n >> 56))
	m.WriteString(string(s))
}

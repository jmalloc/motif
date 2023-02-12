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
func (b Bytes1) AcceptVisitor(v ValueVisitor) { v.VisitBytes1(b) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes2) AcceptVisitor(v ValueVisitor) { v.VisitBytes2(b) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes4) AcceptVisitor(v ValueVisitor) { v.VisitBytes4(b) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (b Bytes8) AcceptVisitor(v ValueVisitor) { v.VisitBytes8(b) }

func (m marshaler) VisitBytes1(b Bytes1) {
	m.WriteControl(octetString1Type)
	n := len(b)
	m.Grow(1 + n)
	m.WriteByte(byte(n))
	m.WriteString(string(b))
}

func (m marshaler) VisitBytes2(b Bytes2) {
	m.WriteControl(octetString2Type)
	n := len(b)
	m.Grow(2 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteString(string(b))
}

func (m marshaler) VisitBytes4(b Bytes4) {
	m.WriteControl(octetString4Type)
	n := len(b)
	m.Grow(4 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteByte(byte(n >> 16))
	m.WriteByte(byte(n >> 24))
	m.WriteString(string(b))
}

func (m marshaler) VisitBytes8(b Bytes8) {
	m.WriteControl(octetString8Type)
	n := len(b)
	m.Grow(8 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteByte(byte(n >> 16))
	m.WriteByte(byte(n >> 24))
	m.WriteByte(byte(n >> 32))
	m.WriteByte(byte(n >> 40))
	m.WriteByte(byte(n >> 48))
	m.WriteByte(byte(n >> 56))
	m.WriteString(string(b))
}

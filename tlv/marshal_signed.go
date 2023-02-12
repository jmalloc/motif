package tlv

const (
	signed1Type = 0b000_00000
	signed2Type = 0b000_00001
	signed4Type = 0b000_00010
	signed8Type = 0b000_00011
)

func (m *marshaler) VisitSigned1(v Signed1) {
	m.control |= signed1Type
	m.payload.WriteByte(byte(v))
}

func (m *marshaler) VisitSigned2(v Signed2) {
	m.control |= signed2Type
	m.payload.Grow(2)
	m.payload.WriteByte(byte(v))
	m.payload.WriteByte(byte(v >> 8))
}

func (m *marshaler) VisitSigned4(v Signed4) {
	m.control |= signed4Type
	m.payload.Grow(4)
	m.payload.WriteByte(byte(v))
	m.payload.WriteByte(byte(v >> 8))
	m.payload.WriteByte(byte(v >> 16))
	m.payload.WriteByte(byte(v >> 24))
}

func (m *marshaler) VisitSigned8(v Signed8) {
	m.control |= signed8Type
	m.payload.Grow(8)
	m.payload.WriteByte(byte(v))
	m.payload.WriteByte(byte(v >> 8))
	m.payload.WriteByte(byte(v >> 16))
	m.payload.WriteByte(byte(v >> 24))
	m.payload.WriteByte(byte(v >> 32))
	m.payload.WriteByte(byte(v >> 40))
	m.payload.WriteByte(byte(v >> 48))
	m.payload.WriteByte(byte(v >> 56))
}

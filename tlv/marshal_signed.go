package tlv

const (
	signed1Type = 0b000_00000
	signed2Type = 0b000_00001
	signed4Type = 0b000_00010
	signed8Type = 0b000_00011
)

func (m marshaler) VisitSigned1(v Signed1) {
	m.WriteControl(signed1Type)
	m.WriteByte(byte(v))
}

func (m marshaler) VisitSigned2(v Signed2) {
	m.WriteControl(signed2Type)
	m.Grow(2)
	m.WriteByte(byte(v))
	m.WriteByte(byte(v >> 8))
}

func (m marshaler) VisitSigned4(v Signed4) {
	m.WriteControl(signed4Type)
	m.Grow(4)
	m.WriteByte(byte(v))
	m.WriteByte(byte(v >> 8))
	m.WriteByte(byte(v >> 16))
	m.WriteByte(byte(v >> 24))
}

func (m marshaler) VisitSigned8(v Signed8) {
	m.WriteControl(signed8Type)
	m.Grow(8)
	m.WriteByte(byte(v))
	m.WriteByte(byte(v >> 8))
	m.WriteByte(byte(v >> 16))
	m.WriteByte(byte(v >> 24))
	m.WriteByte(byte(v >> 32))
	m.WriteByte(byte(v >> 40))
	m.WriteByte(byte(v >> 48))
	m.WriteByte(byte(v >> 56))
}

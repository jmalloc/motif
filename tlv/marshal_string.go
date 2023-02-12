package tlv

const (
	string1Type = 0b000_01100
	string2Type = 0b000_01101
	string4Type = 0b000_01110
	string8Type = 0b000_01111
)

func (m *marshaler) VisitString1(v String1) {
	m.control |= string1Type
	n := len(v)
	m.payload.Grow(1 + n)
	m.payload.WriteByte(byte(n))
	m.payload.WriteString(string(v))
}

func (m *marshaler) VisitString2(v String2) {
	m.control |= string2Type
	n := len(v)
	m.payload.Grow(2 + n)
	m.payload.WriteByte(byte(n))
	m.payload.WriteByte(byte(n >> 8))
	m.payload.WriteString(string(v))
}

func (m *marshaler) VisitString4(v String4) {
	m.control |= string4Type
	n := len(v)
	m.payload.Grow(4 + n)
	m.payload.WriteByte(byte(n))
	m.payload.WriteByte(byte(n >> 8))
	m.payload.WriteByte(byte(n >> 16))
	m.payload.WriteByte(byte(n >> 24))
	m.payload.WriteString(string(v))
}

func (m *marshaler) VisitString8(v String8) {
	m.control |= string8Type
	n := len(v)
	m.payload.Grow(8 + n)
	m.payload.WriteByte(byte(n))
	m.payload.WriteByte(byte(n >> 8))
	m.payload.WriteByte(byte(n >> 16))
	m.payload.WriteByte(byte(n >> 24))
	m.payload.WriteByte(byte(n >> 32))
	m.payload.WriteByte(byte(n >> 40))
	m.payload.WriteByte(byte(n >> 48))
	m.payload.WriteByte(byte(n >> 56))
	m.payload.WriteString(string(v))
}

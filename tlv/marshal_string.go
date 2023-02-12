package tlv

const (
	string1Type = 0b000_01100
	string2Type = 0b000_01101
	string4Type = 0b000_01110
	string8Type = 0b000_01111
)

func (m marshaler) VisitString1(v String1) {
	m.WriteControl(string1Type)
	n := len(v)
	m.Grow(1 + n)
	m.WriteByte(byte(n))
	m.WriteString(string(v))
}

func (m marshaler) VisitString2(v String2) {
	m.WriteControl(string2Type)
	n := len(v)
	m.Grow(2 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteString(string(v))
}

func (m marshaler) VisitString4(v String4) {
	m.WriteControl(string4Type)
	n := len(v)
	m.Grow(4 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteByte(byte(n >> 16))
	m.WriteByte(byte(n >> 24))
	m.WriteString(string(v))
}

func (m marshaler) VisitString8(v String8) {
	m.WriteControl(string8Type)
	n := len(v)
	m.Grow(8 + n)
	m.WriteByte(byte(n))
	m.WriteByte(byte(n >> 8))
	m.WriteByte(byte(n >> 16))
	m.WriteByte(byte(n >> 24))
	m.WriteByte(byte(n >> 32))
	m.WriteByte(byte(n >> 40))
	m.WriteByte(byte(n >> 48))
	m.WriteByte(byte(n >> 56))
	m.WriteString(string(v))
}

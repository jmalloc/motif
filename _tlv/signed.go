package tlv

import (
	"math"
)

// Signed represents a Matter TLV "signed integer" value.
type Signed int64

// AcceptVisitor calls v.VisitSigned().
func (e Signed) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitSigned(ctx, e)
}

const (
	Signed1ElementType = 0b00000
	Signed2ElementType = 0b00001
	Signed4ElementType = 0b00010
	Signed8ElementType = 0b00011
)

func (m marshaler) VisitSigned(ctx VisitContext, e Signed) error {
	if t, ok := ctx.Tag(); ok {
		t.AcceptVisitor(ctx, m)
	}

	var data [9]byte
	data[0] = m.control

	if e >= math.MinInt8 && e <= math.MaxInt8 {
		data[0] |= Signed1ElementType
		data[1] = byte(e)
		m.data.Write(data[:2])
		return nil
	}

	if e >= math.MinInt16 && e <= math.MaxInt16 {
		data[0] |= Signed2ElementType
		m.data.WriteByte(byte(e))
		m.data.WriteByte(byte(e >> 8))
		return nil
	}

	if e >= math.MinInt32 && e <= math.MaxInt32 {
		data[0] |= Signed4ElementType
		m.data.WriteByte(byte(e))
		m.data.WriteByte(byte(e >> 8))
		m.data.WriteByte(byte(e >> 16))
		m.data.WriteByte(byte(e >> 24))
		return nil
	}

	data[0] |= Signed8ElementType
	m.data.WriteByte(byte(e))
	m.data.WriteByte(byte(e >> 8))
	m.data.WriteByte(byte(e >> 16))
	m.data.WriteByte(byte(e >> 24))
	m.data.WriteByte(byte(e >> 32))
	m.data.WriteByte(byte(e >> 40))
	m.data.WriteByte(byte(e >> 48))
	m.data.WriteByte(byte(e >> 56))
}

// func MarshalSigned1(v int8) []byte {
// 	return []byte{
// 		Signed1ElementType,
// 		byte(v),
// 	}
// }

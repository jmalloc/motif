package tlv

import (
	"math"
)

const (
	string1Type = 0b000_01100
	string2Type = 0b000_01101
	string4Type = 0b000_01110
	string8Type = 0b000_01111
)

type (
	string1 string
	string2 string
	string4 string
	string8 string
)

func (v string1) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }
func (v string2) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }
func (v string4) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }
func (v string8) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

func (v string1) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString1(string(v)) }
func (v string2) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString2(string(v)) }
func (v string4) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString4(string(v)) }
func (v string8) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString8(string(v)) }

// String returns an anonymous string element.
//
// The string's length is encoded using the smallest possible integer type.
func String(v string) Value {
	switch {
	case len(v) <= math.MaxUint8:
		return string1(v)
	case len(v) <= math.MaxUint16:
		return string2(v)
	case len(v) <= math.MaxUint32:
		return string4(v)
	default:
		return string8(v)
	}
}

// String1 returns a string element with a 1-octet length.
//
// It panics if the string's length overflows the uint8 type.
func String1(v string) Value {
	if len(v) > math.MaxUint8 {
		panic("string length overflows uint8")
	}

	return string1(v)
}

// String2 returns a string element with a 2-octet length.
//
// It panics if the string's length overflows the uint16 type.
func String2(v string) Value {
	if len(v) > math.MaxUint16 {
		panic("string length overflows uint16")
	}

	return string2(v)
}

// String4 returns a string element with a 4-octet length.
//
// It panics if the string's length overflows the uint32 type.
func String4(v string) Value {
	if len(v) > math.MaxUint32 {
		panic("string length overflows uint32")
	}

	return string4(v)
}

// String8 returns a string element with an 8-octet length.
func String8(v string) Value {
	return string8(v)
}

func (m *marshaler) VisitString1(v string) {
	m.control |= string1Type
	n := len(v)
	m.data.Grow(1 + n)
	m.data.WriteByte(byte(n))
	m.data.WriteString(v)
}

func (m *marshaler) VisitString2(v string) {
	m.control |= string2Type
	n := len(v)
	m.data.Grow(2 + n)
	m.data.WriteByte(byte(n))
	m.data.WriteByte(byte(n >> 8))
	m.data.WriteString(v)
}

func (m *marshaler) VisitString4(v string) {
	m.control |= string4Type
	n := len(v)
	m.data.Grow(4 + n)
	m.data.WriteByte(byte(n))
	m.data.WriteByte(byte(n >> 8))
	m.data.WriteByte(byte(n >> 16))
	m.data.WriteByte(byte(n >> 24))
	m.data.WriteString(v)
}

func (m *marshaler) VisitString8(v string) {
	m.control |= string8Type
	n := len(v)
	m.data.Grow(8 + n)
	m.data.WriteByte(byte(n))
	m.data.WriteByte(byte(n >> 8))
	m.data.WriteByte(byte(n >> 16))
	m.data.WriteByte(byte(n >> 24))
	m.data.WriteByte(byte(n >> 32))
	m.data.WriteByte(byte(n >> 40))
	m.data.WriteByte(byte(n >> 48))
	m.data.WriteByte(byte(n >> 56))
	m.data.WriteString(v)
}

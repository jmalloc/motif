package tlv

import (
	"math"
)

type (
	// Float4 is a single-precision (4 octet) TLV floating-point value.
	Float4 float32

	// Float8 is a double-precision (8 octet) TLV floating-point value.
	Float8 float64
)

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (f Float4) AcceptVisitor(v ValueVisitor) { v.VisitFloat4(f) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (f Float8) AcceptVisitor(v ValueVisitor) { v.VisitFloat8(f) }

func (m marshaler) VisitFloat4(f Float4) {
	bits := math.Float32bits(float32(f))

	m.WriteControl(float4Type)
	m.Grow(4)
	m.WriteByte(byte(bits))
	m.WriteByte(byte(bits >> 8))
	m.WriteByte(byte(bits >> 16))
	m.WriteByte(byte(bits >> 24))
}

func (m marshaler) VisitFloat8(f Float8) {
	bits := math.Float64bits(float64(f))

	m.WriteControl(float8Type)
	m.Grow(8)
	m.WriteByte(byte(bits))
	m.WriteByte(byte(bits >> 8))
	m.WriteByte(byte(bits >> 16))
	m.WriteByte(byte(bits >> 24))
	m.WriteByte(byte(bits >> 32))
	m.WriteByte(byte(bits >> 40))
	m.WriteByte(byte(bits >> 48))
	m.WriteByte(byte(bits >> 56))
}

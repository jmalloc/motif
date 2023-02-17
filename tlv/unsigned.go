package tlv

import "github.com/jmalloc/motif/internal/wire"

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

func (v Unsigned1) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUnsigned1(v)
}

func (v Unsigned2) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUnsigned2(v)
}

func (v Unsigned4) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUnsigned4(v)
}

func (v Unsigned8) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUnsigned8(v)
}

const (
	unsigned1Type = 0b000_00100
	unsigned2Type = 0b000_00101
	unsigned4Type = 0b000_00110
	unsigned8Type = 0b000_00111
)

func (m payloadWriter) VisitUnsigned1(u Unsigned1) error {
	return wire.WriteInt(m, u)
}

func (m payloadWriter) VisitUnsigned2(u Unsigned2) error {
	return wire.WriteInt(m, u)
}

func (m payloadWriter) VisitUnsigned4(u Unsigned4) error {
	return wire.WriteInt(m, u)
}

func (m payloadWriter) VisitUnsigned8(u Unsigned8) error {
	return wire.WriteInt(m, u)
}

func (c *controlWriter) VisitUnsigned1(u Unsigned1) error {
	return c.write(unsigned1Type)
}

func (c *controlWriter) VisitUnsigned2(u Unsigned2) error {
	return c.write(unsigned2Type)
}

func (c *controlWriter) VisitUnsigned4(u Unsigned4) error {
	return c.write(unsigned4Type)
}

func (c *controlWriter) VisitUnsigned8(u Unsigned8) error {
	return c.write(unsigned8Type)
}

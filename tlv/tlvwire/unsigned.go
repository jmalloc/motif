package tlvwire

import (
	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

const (
	unsigned1Type = 0b000_00100
	unsigned2Type = 0b000_00101
	unsigned4Type = 0b000_00110
	unsigned8Type = 0b000_00111
)

func (m payloadWriter) VisitUnsigned1(u tlv.Unsigned1) error {
	return wire.WriteInt(m, u)
}

func (m payloadWriter) VisitUnsigned2(u tlv.Unsigned2) error {
	return wire.WriteInt(m, u)
}

func (m payloadWriter) VisitUnsigned4(u tlv.Unsigned4) error {
	return wire.WriteInt(m, u)
}

func (m payloadWriter) VisitUnsigned8(u tlv.Unsigned8) error {
	return wire.WriteInt(m, u)
}

func (c *controlWriter) VisitUnsigned1(u tlv.Unsigned1) error {
	return c.write(unsigned1Type)
}

func (c *controlWriter) VisitUnsigned2(u tlv.Unsigned2) error {
	return c.write(unsigned2Type)
}

func (c *controlWriter) VisitUnsigned4(u tlv.Unsigned4) error {
	return c.write(unsigned4Type)
}

func (c *controlWriter) VisitUnsigned8(u tlv.Unsigned8) error {
	return c.write(unsigned8Type)
}

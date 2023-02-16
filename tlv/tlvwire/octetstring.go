package tlvwire

import (
	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

const (
	octetString1Type = 0b000_10000
	octetString2Type = 0b000_10001
	octetString4Type = 0b000_10010
	octetString8Type = 0b000_10011
)

func (m *payloadWriter) VisitOctetString1(v tlv.OctetString1) error {
	return wire.WriteString[uint8](m, v)
}

func (m *payloadWriter) VisitOctetString2(v tlv.OctetString2) error {
	return wire.WriteString[uint16](m, v)
}

func (m *payloadWriter) VisitOctetString4(v tlv.OctetString4) error {
	return wire.WriteString[uint32](m, v)
}

func (m *payloadWriter) VisitOctetString8(v tlv.OctetString8) error {
	return wire.WriteString[uint64](m, v)
}

func (c *controlWriter) VisitOctetString1(v tlv.OctetString1) error {
	return c.write(octetString1Type)
}

func (c *controlWriter) VisitOctetString2(v tlv.OctetString2) error {
	return c.write(octetString2Type)
}

func (c *controlWriter) VisitOctetString4(v tlv.OctetString4) error {
	return c.write(octetString4Type)
}

func (c *controlWriter) VisitOctetString8(v tlv.OctetString8) error {
	return c.write(octetString8Type)
}

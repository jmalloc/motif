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

func (m *payloadWriter) VisitBytes1(v tlv.Bytes1) error { return wire.WriteString[uint8](m, v) }
func (m *payloadWriter) VisitBytes2(v tlv.Bytes2) error { return wire.WriteString[uint16](m, v) }
func (m *payloadWriter) VisitBytes4(v tlv.Bytes4) error { return wire.WriteString[uint32](m, v) }
func (m *payloadWriter) VisitBytes8(v tlv.Bytes8) error { return wire.WriteString[uint64](m, v) }
func (c *controlWriter) VisitBytes1(v tlv.Bytes1) error { return c.set(octetString1Type) }
func (c *controlWriter) VisitBytes2(v tlv.Bytes2) error { return c.set(octetString2Type) }
func (c *controlWriter) VisitBytes4(v tlv.Bytes4) error { return c.set(octetString4Type) }
func (c *controlWriter) VisitBytes8(v tlv.Bytes8) error { return c.set(octetString8Type) }

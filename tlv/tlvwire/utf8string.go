package tlvwire

import (
	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

const (
	utf8String1Type = 0b000_01100
	utf8String2Type = 0b000_01101
	utf8String4Type = 0b000_01110
	utf8String8Type = 0b000_01111
)

func (m payloadWriter) VisitString1(s tlv.String1) error { return wire.WriteString[uint8](m, s) }
func (m payloadWriter) VisitString2(s tlv.String2) error { return wire.WriteString[uint16](m, s) }
func (m payloadWriter) VisitString4(s tlv.String4) error { return wire.WriteString[uint32](m, s) }
func (m payloadWriter) VisitString8(s tlv.String8) error { return wire.WriteString[uint64](m, s) }

func (c *controlWriter) VisitString1(s tlv.String1) error { return c.set(utf8String1Type) }
func (c *controlWriter) VisitString2(s tlv.String2) error { return c.set(utf8String2Type) }
func (c *controlWriter) VisitString4(s tlv.String4) error { return c.set(utf8String4Type) }
func (c *controlWriter) VisitString8(s tlv.String8) error { return c.set(utf8String8Type) }

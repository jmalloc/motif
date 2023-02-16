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

func (m payloadWriter) VisitUTF8String1(s tlv.UTF8String1) error {
	return wire.WriteString[uint8](m, s)
}
func (m payloadWriter) VisitUTF8String2(s tlv.UTF8String2) error {
	return wire.WriteString[uint16](m, s)
}
func (m payloadWriter) VisitUTF8String4(s tlv.UTF8String4) error {
	return wire.WriteString[uint32](m, s)
}
func (m payloadWriter) VisitUTF8String8(s tlv.UTF8String8) error {
	return wire.WriteString[uint64](m, s)
}

func (c *controlWriter) VisitUTF8String1(s tlv.UTF8String1) error {
	return c.write(utf8String1Type)
}

func (c *controlWriter) VisitUTF8String2(s tlv.UTF8String2) error {
	return c.write(utf8String2Type)
}

func (c *controlWriter) VisitUTF8String4(s tlv.UTF8String4) error {
	return c.write(utf8String4Type)
}

func (c *controlWriter) VisitUTF8String8(s tlv.UTF8String8) error {
	return c.write(utf8String8Type)
}

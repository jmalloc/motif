package tlv

import "github.com/jmalloc/motif/internal/wire"

type (
	// UTF8String1 is a UTF-8 string with a length that can be represented by a
	// 1 octet integer.
	UTF8String1 string

	// UTF8String2 is a UTF-8 string with a length that can be represented by a
	// 2 octet integer.
	UTF8String2 string

	// UTF8String4 is a UTF-8 string with a length that can be represented by a
	// 4 octet integer.
	UTF8String4 string

	// UTF8String8 is a UTF-8 string with a length that can be represented by an
	// 8 octet integer.
	UTF8String8 string
)

func (v UTF8String1) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUTF8String1(v)
}

func (v UTF8String2) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUTF8String2(v)
}

func (v UTF8String4) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUTF8String4(v)
}

func (v UTF8String8) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitUTF8String8(v)
}

const (
	utf8String1Type = 0b000_01100
	utf8String2Type = 0b000_01101
	utf8String4Type = 0b000_01110
	utf8String8Type = 0b000_01111
)

func (m payloadWriter) VisitUTF8String1(s UTF8String1) error {
	return wire.WriteString[uint8](m, s)
}

func (m payloadWriter) VisitUTF8String2(s UTF8String2) error {
	return wire.WriteString[uint16](m, s)
}

func (m payloadWriter) VisitUTF8String4(s UTF8String4) error {
	return wire.WriteString[uint32](m, s)
}

func (m payloadWriter) VisitUTF8String8(s UTF8String8) error {
	return wire.WriteString[uint64](m, s)
}

func (c *controlWriter) VisitUTF8String1(s UTF8String1) error {
	return c.write(utf8String1Type)
}

func (c *controlWriter) VisitUTF8String2(s UTF8String2) error {
	return c.write(utf8String2Type)
}

func (c *controlWriter) VisitUTF8String4(s UTF8String4) error {
	return c.write(utf8String4Type)
}

func (c *controlWriter) VisitUTF8String8(s UTF8String8) error {
	return c.write(utf8String8Type)
}

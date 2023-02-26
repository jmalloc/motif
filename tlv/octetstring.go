package tlv

type (
	// OctetString1 is an octet string with a length that can be represented by
	// a 1 octet integer.
	OctetString1 []byte

	// OctetString2 is an octet string with a length that can be represented by
	// a 2 octet integer.
	OctetString2 []byte

	// OctetString4 is an octet string with a length that can be represented by
	// a 4 octet integer.
	OctetString4 []byte

	// OctetString8 is an octet string with a length that can be represented by
	// an 8 octet integer.
	OctetString8 []byte
)

func (v OctetString1) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitOctetString1(v)
}

func (v OctetString2) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitOctetString2(v)
}

func (v OctetString4) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitOctetString4(v)
}

func (v OctetString8) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitOctetString8(v)
}

const (
	octetString1Type = 0b000_10000
	octetString2Type = 0b000_10001
	octetString4Type = 0b000_10010
	octetString8Type = 0b000_10011
)

func (m *payloadWriter) VisitOctetString1(v OctetString1) error {
	return writeString[uint8](m, v)
}

func (m *payloadWriter) VisitOctetString2(v OctetString2) error {
	return writeString[uint16](m, v)
}

func (m *payloadWriter) VisitOctetString4(v OctetString4) error {
	return writeString[uint32](m, v)
}

func (m *payloadWriter) VisitOctetString8(v OctetString8) error {
	return writeString[uint64](m, v)
}

func (c *controlWriter) VisitOctetString1(v OctetString1) error {
	return c.write(octetString1Type)
}

func (c *controlWriter) VisitOctetString2(v OctetString2) error {
	return c.write(octetString2Type)
}

func (c *controlWriter) VisitOctetString4(v OctetString4) error {
	return c.write(octetString4Type)
}

func (c *controlWriter) VisitOctetString8(v OctetString8) error {
	return c.write(octetString8Type)
}

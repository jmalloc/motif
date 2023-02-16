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

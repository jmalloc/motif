package tlv

type (
	// Bytes1 is an octet-string with a 1 octet length.
	Bytes1 []byte

	// Bytes2 is an octet-string with a 2 octet length.
	Bytes2 []byte

	// Bytes4 is an octet-string with a 4 octet length.
	Bytes4 []byte

	// Bytes8 is an octet-string with an 8 octet length.
	Bytes8 []byte
)

func (v Bytes1) acceptVisitor(vis ValueVisitor) error { return vis.VisitBytes1(v) }
func (v Bytes2) acceptVisitor(vis ValueVisitor) error { return vis.VisitBytes2(v) }
func (v Bytes4) acceptVisitor(vis ValueVisitor) error { return vis.VisitBytes4(v) }
func (v Bytes8) acceptVisitor(vis ValueVisitor) error { return vis.VisitBytes8(v) }

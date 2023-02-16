package tlv

type (
	// String1 is a UTF-8 string with a 1 octet length.
	String1 string

	// String2 is a UTF-8 string with a 2 octet length.
	String2 string

	// String4 is a UTF-8 string with a 4 octet length.
	String4 string

	// String8 is a UTF-8 string with an 8 octet length.
	String8 string
)

func (v String1) acceptVisitor(vis ValueVisitor) error { return vis.VisitString1(v) }
func (v String2) acceptVisitor(vis ValueVisitor) error { return vis.VisitString2(v) }
func (v String4) acceptVisitor(vis ValueVisitor) error { return vis.VisitString4(v) }
func (v String8) acceptVisitor(vis ValueVisitor) error { return vis.VisitString8(v) }

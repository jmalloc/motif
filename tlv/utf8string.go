package tlv

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

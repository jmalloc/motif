package tlv

const (
	signed1Type = 0b000_00000
	signed2Type = 0b000_00001
	signed4Type = 0b000_00010
	signed8Type = 0b000_00011

	unsigned1Type = 0b000_00100
	unsigned2Type = 0b000_00101
	unsigned4Type = 0b000_00110
	unsigned8Type = 0b000_00111

	boolFalseType = 0b000_01000
	boolTrueType  = 0b000_01001
	nullType      = 0b000_10100
	structType    = 0b000_10101

	utf8String1Type = 0b000_01100
	utf8String2Type = 0b000_01101
	utf8String4Type = 0b000_01110
	utf8String8Type = 0b000_01111

	octetString1Type = 0b000_10000
	octetString2Type = 0b000_10001
	octetString4Type = 0b000_10010
	octetString8Type = 0b000_10011
)

const (
	anonymousTagForm       = 0b000_00000
	contextSpecificTagForm = 0b001_00000
)

const endOfContainer = 0b000_11000

package tlv

const (
	signed1Type   = 0b000_00000
	signed2Type   = 0b000_00001
	signed4Type   = 0b000_00010
	signed8Type   = 0b000_00011
	unsigned1Type = 0b000_00100
	unsigned2Type = 0b000_00101
	unsigned4Type = 0b000_00110
	unsigned8Type = 0b000_00111
	nullType      = 0b000_10100
	structType    = 0b000_10101
	string1Type   = 0b000_01100
	string2Type   = 0b000_01101
	string4Type   = 0b000_01110
	string8Type   = 0b000_01111
)

const (
	anonymousTagForm       = 0b000_00000
	contextSpecificTagForm = 0b001_00000
)

const endOfContainer = 0b000_11000

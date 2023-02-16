package tlv

const (
	typeMask    = 0b000_11111
	tagFormMask = 0b111_00000
)

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

	float4Type = 0b000_01010
	float8Type = 0b000_01011

	nullType = 0b000_10100

	structType = 0b000_10101
	arrayType  = 0b000_10110
	listType   = 0b000_10111

	utf8String1Type = 0b000_01100
	utf8String2Type = 0b000_01101
	utf8String4Type = 0b000_01110
	utf8String8Type = 0b000_01111

	octetString1Type = 0b000_10000
	octetString2Type = 0b000_10001
	octetString4Type = 0b000_10010
	octetString8Type = 0b000_10011

	endOfContainer = 0b000_11000

	anonymousTagForm        = 0b000_00000
	contextSpecificTagForm  = 0b001_00000
	commonProfileTag2Form   = 0b010_00000
	commonProfileTag4Form   = 0b011_00000
	implicitProfileTag2Form = 0b100_00000
	implicitProfileTag4Form = 0b101_00000
	fullyQualifiedTag6Form  = 0b110_00000
	fullyQualifiedTag8Form  = 0b111_00000
)

// controlOctetBuilder is an implementation of ValueVisitor and TagVisitor that
// builds a control octet that describes the tag form and element type.
type controlOctetBuilder struct {
	Value byte
}

func (c *controlOctetBuilder) update(v byte) error {
	c.Value |= v
	return nil
}

func (c *controlOctetBuilder) VisitSigned1(s Signed1) error {
	return c.update(signed1Type)
}

func (c *controlOctetBuilder) VisitSigned2(s Signed2) error {
	return c.update(signed2Type)
}

func (c *controlOctetBuilder) VisitSigned4(s Signed4) error {
	return c.update(signed4Type)
}

func (c *controlOctetBuilder) VisitSigned8(s Signed8) error {
	return c.update(signed8Type)
}

func (c *controlOctetBuilder) VisitUnsigned1(u Unsigned1) error {
	return c.update(unsigned1Type)
}

func (c *controlOctetBuilder) VisitUnsigned2(u Unsigned2) error {
	return c.update(unsigned2Type)
}

func (c *controlOctetBuilder) VisitUnsigned4(u Unsigned4) error {
	return c.update(unsigned4Type)
}

func (c *controlOctetBuilder) VisitUnsigned8(u Unsigned8) error {
	return c.update(unsigned8Type)
}

func (c *controlOctetBuilder) VisitBool(b Bool) error {
	if b {
		return c.update(boolTrueType)
	}
	return c.update(boolFalseType)
}

func (c *controlOctetBuilder) VisitFloat4(f Float4) error {
	return c.update(float4Type)
}

func (c *controlOctetBuilder) VisitFloat8(f Float8) error {
	return c.update(float8Type)
}

func (c *controlOctetBuilder) VisitNull() error {
	return c.update(nullType)
}

func (c *controlOctetBuilder) VisitStruct(s Struct) error {
	return c.update(structType)
}

func (c *controlOctetBuilder) VisitArray(a Array) error {
	return c.update(arrayType)
}

func (c *controlOctetBuilder) VisitList(l List) error {
	return c.update(listType)
}

func (c *controlOctetBuilder) VisitString1(s String1) error {
	return c.update(utf8String1Type)
}

func (c *controlOctetBuilder) VisitString2(s String2) error {
	return c.update(utf8String2Type)
}

func (c *controlOctetBuilder) VisitString4(s String4) error {
	return c.update(utf8String4Type)
}

func (c *controlOctetBuilder) VisitString8(s String8) error {
	return c.update(utf8String8Type)
}

func (c *controlOctetBuilder) VisitBytes1(b Bytes1) error {
	return c.update(octetString1Type)
}

func (c *controlOctetBuilder) VisitBytes2(b Bytes2) error {
	return c.update(octetString2Type)
}

func (c *controlOctetBuilder) VisitBytes4(b Bytes4) error {
	return c.update(octetString4Type)
}

func (c *controlOctetBuilder) VisitBytes8(b Bytes8) error {
	return c.update(octetString8Type)
}

func (c *controlOctetBuilder) VisitAnonymousTag() error {
	return c.update(anonymousTagForm)
}

func (c *controlOctetBuilder) VisitContextSpecificTag(t ContextSpecificTag) error {
	return c.update(contextSpecificTagForm)
}

func (c *controlOctetBuilder) VisitCommonProfileTag2(t CommonProfileTag2) error {
	return c.update(commonProfileTag2Form)
}

func (c *controlOctetBuilder) VisitCommonProfileTag4(t CommonProfileTag4) error {
	return c.update(commonProfileTag4Form)
}

func (c *controlOctetBuilder) VisitImplicitProfileTag2(t ImplicitProfileTag2) error {
	return c.update(implicitProfileTag2Form)
}

func (c *controlOctetBuilder) VisitImplicitProfileTag4(t ImplicitProfileTag4) error {
	return c.update(implicitProfileTag4Form)
}

func (c *controlOctetBuilder) VisitFullyQualifiedTag6(t FullyQualifiedTag6) error {
	return c.update(fullyQualifiedTag6Form)
}

func (c *controlOctetBuilder) VisitFullyQualifiedTag8(t FullyQualifiedTag8) error {
	return c.update(fullyQualifiedTag8Form)
}

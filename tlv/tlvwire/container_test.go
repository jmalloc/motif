package tlvwire_test

import (
	"github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"structure containers",
	testContainer[tlv.Struct],
	Entry(
		"nil structure",
		tlv.Struct(nil),
		[]byte{0x15, 0x18},
	),
	Entry(
		"empty structure",
		tlv.Struct{},
		[]byte{0x15, 0x18},
	),
	Entry(
		"two context specific tags, signed inte­ger, 1 octet values {0 = 42, 1 = -17}",
		tlv.Struct{
			{
				T: tlv.ContextSpecificTag(0),
				V: tlv.Signed1(42),
			},
			{
				T: tlv.ContextSpecificTag(1),
				V: tlv.Signed1(-17),
			},
		},
		[]byte{
			0x15, 0x20, 0x00, 0x2a, 0x20, 0x01, 0xef, 0x18,
		},
	),
)

var _ = DescribeTable(
	"array containers",
	testContainer[tlv.Array],
	Entry(
		"nil array",
		tlv.Array(nil),
		[]byte{0x16, 0x18},
	),
	Entry(
		"empty array",
		tlv.Array{},
		[]byte{0x16, 0x18},
	),
	Entry(
		"signed integer members, 1-octet values, [0, 1, 2, 3, 4]",
		tlv.Array{
			tlv.Signed1(0),
			tlv.Signed1(1),
			tlv.Signed1(2),
			tlv.Signed1(3),
			tlv.Signed1(4),
		},
		[]byte{
			0x16, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x00,
			0x03, 0x00, 0x04, 0x18,
		},
	),
	Entry(
		`mix of element types, [42, -170000, {}, 17.9, "Hello!"]`,
		tlv.Array{
			tlv.Signed1(42),
			tlv.Signed4(-170000),
			tlv.Struct(nil),
			tlv.Float4(17.9),
			tlv.String1("Hello!"),
		},
		[]byte{
			0x16, 0x00, 0x2a, 0x02, 0xf0, 0x67, 0xfd, 0xff,
			0x15, 0x18, 0x0a, 0x33, 0x33, 0x8f, 0x41, 0x0c,
			0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21, 0x18,
		},
	),
)

var _ = DescribeTable(
	"list containers",
	testContainer[tlv.List],
	Entry(
		"nil list",
		tlv.List(nil),
		[]byte{0x17, 0x18},
	),
	Entry(
		"empty list",
		tlv.List{},
		[]byte{0x17, 0x18},
	),
	Entry(
		"mix of anonymous and context tags, signed integer, 1 octet values, [1, 0 = 42, 2, 3, 0 = -17]",
		tlv.List{
			{
				T: tlv.AnonymousTag,
				V: tlv.Signed1(1),
			},
			{
				T: tlv.ContextSpecificTag(0),
				V: tlv.Signed1(42),
			},
			{
				T: tlv.AnonymousTag,
				V: tlv.Signed1(2),
			},
			{
				T: tlv.AnonymousTag,
				V: tlv.Signed1(3),
			},
			{
				T: tlv.ContextSpecificTag(0),
				V: tlv.Signed1(-17),
			},
		},
		[]byte{
			0x17, 0x00, 0x01, 0x20, 0x00, 0x2a, 0x00, 0x02,
			0x00, 0x03, 0x20, 0x00, 0xef, 0x18,
		},
	),
)

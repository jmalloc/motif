package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"structure containers",
	testContainer[Struct],
	Entry(
		"nil structure",
		Struct(nil),
		[]byte{0x15, 0x18},
	),
	Entry(
		"empty structure",
		Struct{},
		[]byte{0x15, 0x18},
	),
	Entry(
		"two context specific tags, signed inte­ger, 1 octet values {0 = 42, 1 = -17}",
		Struct{
			{
				T: ContextSpecificTag(0),
				V: Signed1(42),
			},
			{
				T: ContextSpecificTag(1),
				V: Signed1(-17),
			},
		},
		[]byte{
			0x15, 0x20, 0x00, 0x2a, 0x20, 0x01, 0xef, 0x18,
		},
	),
)

var _ = DescribeTable(
	"array containers",
	testContainer[Array],
	Entry(
		"nil array",
		Array(nil),
		[]byte{0x16, 0x18},
	),
	Entry(
		"empty array",
		Array{},
		[]byte{0x16, 0x18},
	),
	Entry(
		"signed integer members, 1-octet values, [0, 1, 2, 3, 4]",
		Array{
			Signed1(0),
			Signed1(1),
			Signed1(2),
			Signed1(3),
			Signed1(4),
		},
		[]byte{
			0x16, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x00,
			0x03, 0x00, 0x04, 0x18,
		},
	),
	Entry(
		`mix of element types, [42, -170000, {}, 17.9, "Hello!"]`,
		Array{
			Signed1(42),
			Signed4(-170000),
			Struct(nil),
			Single(17.9),
			UTF8String1("Hello!"),
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
	testContainer[List],
	Entry(
		"nil list",
		List(nil),
		[]byte{0x17, 0x18},
	),
	Entry(
		"empty list",
		List{},
		[]byte{0x17, 0x18},
	),
	Entry(
		"mix of anonymous and context tags, signed integer, 1 octet values, [1, 0 = 42, 2, 3, 0 = -17]",
		List{
			{
				T: AnonymousTag,
				V: Signed1(1),
			},
			{
				T: ContextSpecificTag(0),
				V: Signed1(42),
			},
			{
				T: AnonymousTag,
				V: Signed1(2),
			},
			{
				T: AnonymousTag,
				V: Signed1(3),
			},
			{
				T: ContextSpecificTag(0),
				V: Signed1(-17),
			},
		},
		[]byte{
			0x17, 0x00, 0x01, 0x20, 0x00, 0x2a, 0x00, 0x02,
			0x00, 0x03, 0x20, 0x00, 0xef, 0x18,
		},
	),
)

package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("func Marshal() and Unmarshal()", func() {
	DescribeTable(
		"it encodes/decodes lists correctly",
		testContainerEncoding[List],
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
				{AnonymousTag, Signed1(1)},
				{ContextSpecificTag(0), Signed1(42)},
				{AnonymousTag, Signed1(2)},
				{AnonymousTag, Signed1(3)},
				{ContextSpecificTag(0), Signed1(-17)},
			},
			[]byte{
				0x17, 0x00, 0x01, 0x20, 0x00, 0x2a, 0x00, 0x02,
				0x00, 0x03, 0x20, 0x00, 0xef, 0x18,
			},
		),
	)
})

package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("func Marshal() and Unmarshal()", func() {
	DescribeTable(
		"it encodes structures correctly",
		testContainerEncoding[Struct],
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
				{ContextSpecificTag(0), Signed1(42)},
				{ContextSpecificTag(1), Signed1(-17)},
			},
			[]byte{
				0x15, 0x20, 0x00, 0x2a, 0x20, 0x01, 0xef, 0x18,
			},
		),
	)
})

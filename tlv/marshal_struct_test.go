package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes structures correctly",
		func(v Value, expect []byte) {
			data, err := Marshal(Root{Value: v})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expect))
		},
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
			[]byte{0x15, 0x20, 0x00, 0x2a, 0x20, 0x01, 0xef, 0x18},
		),
	)
})

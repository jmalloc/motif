package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes arrays correctly",
		func(v Value, expect []byte) {
			data, err := Marshal(Root{Value: v})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expect))
		},
		Entry(
			"empty array",
			Array{},
			[]byte{0x16, 0x18},
		),
		Entry(
			"signed integer members, 1-octet values, [0, 1, 2, 3, 4]",
			Array{
				{Signed1(0)},
				{Signed1(1)},
				{Signed1(2)},
				{Signed1(3)},
				{Signed1(4)},
			},
			[]byte{
				0x16, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x00,
				0x03, 0x00, 0x04, 0x18,
			},
		),
		Entry(
			`mix of element types, [42, -170000, {}, 17.9, "Hello!"]`,
			Array{
				{Signed1(42)},
				{Signed4(-170000)},
				{Struct{}},
				{Float4(17.9)},
				{String1("Hello!")},
			},
			[]byte{
				0x16, 0x00, 0x2a, 0x02, 0xf0, 0x67, 0xfd, 0xff,
				0x15, 0x18, 0x0a, 0x33, 0x33, 0x8f, 0x41, 0x0c,
				0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21, 0x18,
			},
		),
	)
})

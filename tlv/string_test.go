package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type String", func() {
	Describe("func Marshal()", func() {
		DescribeTable(
			"it encodes the string correctly",
			func(s string, expect []byte) {
				data, err := Marshal(String(s))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(data).To(Equal(expect))
			},
			Entry("empty", "", []byte{0x0c, 0x00}),
			Entry("ASCII", "Hello!", []byte{0x0c, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21}),
			Entry("UTF-8", "Tschüs", []byte{0x0c, 0x07, 0x54, 0x73, 0x63, 0x68, 0xc3, 0xbc, 0x73}),
		)
	})
})

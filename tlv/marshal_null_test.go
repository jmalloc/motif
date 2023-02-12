package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes NULL correctly",
		func(v Value, expect []byte) {
			data, err := Marshal(v)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expect))
		},
		Entry(
			"null (anonymous)",
			Null,
			[]byte{0x14},
		),
	)
})

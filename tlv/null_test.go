package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes/decodes NULL correctly",
		func(expectValue Value, expectData []byte) {
			data, err := Marshal(Root{Value: expectValue})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expectData))

			e, err := Unmarshal(data)
			Expect(err).ShouldNot(HaveOccurred())

			t, v := e.Components()
			Expect(t).To(Equal(AnonymousTag))
			Expect(v).To(Equal(expectValue))
		},
		Entry(
			"null",
			Null,
			[]byte{0x14},
		),
	)
})

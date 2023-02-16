package tlv_test

import (
	"bytes"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes/decodes NULL correctly",
		func(expectValue Value, expectData []byte) {
			data := &bytes.Buffer{}
			err := Marshal(data, Root{V: expectValue})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.Bytes()).To(Equal(expectData))

			e, err := Unmarshal(data)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(e.Value()).To(Equal(expectValue))
		},
		Entry(
			"null",
			Null,
			[]byte{0x14},
		),
	)
})

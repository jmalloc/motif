package tlv_test

import (
	"bytes"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal() and Unmarshal()", func() {
	DescribeTable(
		"it encodes/decodes booleans correctly",
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
			"false",
			False,
			[]byte{0x08},
		),
		Entry(
			"true",
			True,
			[]byte{0x09},
		),
	)
})

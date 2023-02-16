package tlv_test

import (
	"bytes"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes/decodes unsigned integers correctly",
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
			"1 octet",
			Unsigned1(42),
			[]byte{0x04, 0x2a},
		),
		Entry(
			"2 octet",
			Unsigned2(420),
			[]byte{0x05, 0xa4, 0x01},
		),
		Entry(
			"4 octet",
			Unsigned4(420000),
			[]byte{0x06, 0xa0, 0x68, 0x06, 0x00},
		),
		Entry(
			"8 octet",
			Unsigned8(40000000000),
			[]byte{0x07, 0x00, 0x90, 0x2f, 0x50, 0x09, 0x00, 0x00, 0x00},
		),
	)
})

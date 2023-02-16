package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes/decodes unsigned integers correctly",
		testScalarEncoding,
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

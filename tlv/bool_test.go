package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("func Marshal() and Unmarshal()", func() {
	DescribeTable(
		"it encodes/decodes booleans correctly",
		testScalarEncoding,
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

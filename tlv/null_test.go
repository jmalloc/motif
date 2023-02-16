package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("func Marshal() and Unmarshal()", func() {
	DescribeTable(
		"it encodes/decodes NULL correctly",
		testScalarEncoding,
		Entry(
			"null",
			Null,
			[]byte{0x14},
		),
	)
})

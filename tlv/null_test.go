package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"null",
	testScalar,
	Entry(
		"null",
		Null,
		[]byte{0x14},
	),
)

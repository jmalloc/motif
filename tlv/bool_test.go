package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"booleans",
	testScalar,
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

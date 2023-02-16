package tlvwire_test

import (
	"github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"signed integers",
	testScalar,
	Entry(
		"1 octet, positive",
		tlv.Signed1(42),
		[]byte{0x00, 0x2a},
	),
	Entry(
		"1 octet, negative",
		tlv.Signed1(-17),
		[]byte{0x00, 0xef},
	),
	Entry(
		"2 octet, positive",
		tlv.Signed2(420),
		[]byte{0x01, 0xa4, 0x01},
	),
	Entry(
		"2 octet, negative",
		tlv.Signed2(-1700),
		[]byte{0x01, 0x5c, 0xf9},
	),
	Entry(
		"4 octet, positive",
		tlv.Signed4(42000),
		[]byte{0x02, 0x10, 0xa4, 0x00, 0x00},
	),
	Entry(
		"4 octet, negative",
		tlv.Signed4(-170000),
		[]byte{0x02, 0xf0, 0x67, 0xfd, 0xff},
	),
	Entry(
		"8 octet, positive",
		tlv.Signed8(40000000000),
		[]byte{0x03, 0x00, 0x90, 0x2f, 0x50, 0x09, 0x00, 0x00, 0x00},
	),
	Entry(
		"8 octet, negative",
		tlv.Signed8(-170000000000),
		[]byte{0x03, 0x00, 0xdc, 0x35, 0x6b, 0xd8, 0xff, 0xff, 0xff},
	),
)

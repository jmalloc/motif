package tlvwire_test

import (
	"github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"tags",
	testTag,
	Entry(
		"anonymous",
		tlv.AnonymousTag,
		[]byte{0x00, 0x00},
	),
	Entry(
		"context-specific",
		tlv.ContextSpecificTag(1),
		[]byte{0x20, 0x01, 0x00},
	),
	Entry(
		"common profile, 2 octet",
		tlv.CommonProfileTag2(1),
		[]byte{0x40, 0x01, 0x00, 0x00},
	),
	Entry(
		"common profile, 4 octet",
		tlv.CommonProfileTag4(100000),
		[]byte{0x60, 0xa0, 0x86, 0x01, 0x00, 0x00},
	),
	Entry(
		"implicit profile, 2 octet",
		tlv.ImplicitProfileTag2(1),
		[]byte{0x80, 0x01, 0x00, 0x00},
	),
	Entry(
		"implicit profile, 4 octet",
		tlv.ImplicitProfileTag4(100000),
		[]byte{0xA0, 0xa0, 0x86, 0x01, 0x00, 0x00},
	),
	Entry(
		"fully-qualified, 6 octet",
		tlv.FullyQualifiedTag6{
			VendorID: 0xfff1,
			Profile:  0xdeed,
			Tag:      0x01,
		},
		[]byte{0xc0, 0xf1, 0xff, 0xed, 0xde, 0x01, 0x00, 0x00},
	),
	Entry(
		"fully-qualified, 8 octet",
		tlv.FullyQualifiedTag8{
			VendorID: 0xfff1,
			Profile:  0xdeed,
			Tag:      0xaa55feed,
		},
		[]byte{0xe0, 0xf1, 0xff, 0xed, 0xde, 0xed, 0xfe, 0x55, 0xaa, 0x00},
	),
)

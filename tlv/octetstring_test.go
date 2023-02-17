package tlv_test

import (
	"bytes"
	"math"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
)

var _ = DescribeTable(
	"octet strings",
	testScalar,
	Entry(
		"1 octet length, empty",
		OctetString1{},
		[]byte{0x10, 0x00},
	),
	Entry(
		"1 octet length",
		OctetString1{0x00, 0x01, 0x02, 0x03, 0x04},
		[]byte{0x10, 0x05, 0x00, 0x01, 0x02, 0x03, 0x04},
	),
	Entry(
		"2 octet length, empty",
		OctetString2{},
		[]byte{0x11, 0x00, 0x00},
	),
	Entry(
		"2 octet length",
		OctetString2(bytes.Repeat([]byte{0x10}, math.MaxUint8+1)),
		append(
			[]byte{0x11, 0x00, 0x01},
			bytes.Repeat([]byte{0x10}, math.MaxUint8+1)...,
		),
	),
	Entry(
		"4 octet length, empty",
		OctetString4{},
		[]byte{0x12, 0x00, 0x00, 0x00, 0x00},
	),
	Entry(
		"4 octet length",
		OctetString4(bytes.Repeat([]byte{0x10}, math.MaxUint16+1)),
		append(
			[]byte{0x12, 0x00, 0x00, 0x01, 0x00},
			bytes.Repeat([]byte{0x10}, math.MaxUint16+1)...,
		),
	),
	Entry(
		"8 octet length, empty",
		OctetString8{},
		[]byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	),
	// Note, we don't check for lengths larger than 4 octets because
	// allocating a slice of that length takes too long (~50s on M1 MBP).
)

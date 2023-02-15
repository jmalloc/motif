package tlv_test

import (
	"bytes"
	"math"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes/decodes octet strings correctly",
		func(expectValue Value, expectData []byte) {
			data, err := Marshal(Root{V: expectValue})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expectData))

			e, err := Unmarshal(data)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(e.Value()).To(Equal(expectValue))
		},
		Entry(
			"1 octet length, empty",
			Bytes1{},
			[]byte{0x10, 0x00},
		),
		Entry(
			"1 octet length",
			Bytes1{0x00, 0x01, 0x02, 0x03, 0x04},
			[]byte{0x10, 0x05, 0x00, 0x01, 0x02, 0x03, 0x04},
		),
		Entry(
			"2 octet length, empty",
			Bytes2{},
			[]byte{0x11, 0x00, 0x00},
		),
		Entry(
			"2 octet length",
			Bytes2(bytes.Repeat([]byte{0x10}, math.MaxUint8+1)),
			append(
				[]byte{0x11, 0x00, 0x01},
				bytes.Repeat([]byte{0x10}, math.MaxUint8+1)...,
			),
		),
		Entry(
			"4 octet length, empty",
			Bytes4{},
			[]byte{0x12, 0x00, 0x00, 0x00, 0x00},
		),
		Entry(
			"4 octet length",
			Bytes4(bytes.Repeat([]byte{0x10}, math.MaxUint16+1)),
			append(
				[]byte{0x12, 0x00, 0x00, 0x01, 0x00},
				bytes.Repeat([]byte{0x10}, math.MaxUint16+1)...,
			),
		),
		Entry(
			"8 octet length, empty",
			Bytes8{},
			[]byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		),
		// Note, we don't check for lengths larger than 4 octets because
		// allocating a slice of that length takes too long (~50s on M1 MBP).
	)
})

package tlv_test

import (
	"bytes"
	"math"
	"strings"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes UTF-8 strings correctly",
		func(v Value, expect []byte) {
			data, err := Marshal(Root{Value: v})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expect))
		},
		Entry(
			"1 octet length, empty",
			String1(""),
			[]byte{0x0c, 0x00},
		),
		Entry(
			"1 octet length, ASCII",
			String1("Hello!"),
			[]byte{0x0c, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21}),
		Entry(
			"1 octet length, UTF-8",
			String1("Tschüs"),
			[]byte{0x0c, 0x07, 0x54, 0x73, 0x63, 0x68, 0xc3, 0xbc, 0x73},
		),
		Entry(
			"2 octet length, empty",
			String2(""),
			[]byte{0x0d, 0x00, 0x00},
		),
		Entry(
			"2 octet length, ASCII",
			String2("Hello!"),
			[]byte{0x0d, 0x06, 0x00, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21}),
		Entry(
			"2 octet length, UTF-8",
			String2("Tschüs"),
			[]byte{0x0d, 0x07, 0x00, 0x54, 0x73, 0x63, 0x68, 0xc3, 0xbc, 0x73},
		),
		Entry(
			"2 octet length, length larger than 1 octet",
			String2(strings.Repeat(" ", math.MaxUint8+1)),
			append(
				[]byte{0x0d, 0x00, 0x01},
				bytes.Repeat([]byte{' '}, math.MaxUint8+1)...,
			),
		),
		Entry(
			"4 octet length, empty",
			String4(""),
			[]byte{0x0e, 0x00, 0x00, 0x00, 0x00},
		),
		Entry(
			"4 octet length, ASCII",
			String4("Hello!"),
			[]byte{0x0e, 0x06, 0x00, 0x00, 0x00, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21}),
		Entry(
			"4 octet length, UTF-8",
			String4("Tschüs"),
			[]byte{0x0e, 0x07, 0x00, 0x00, 0x00, 0x54, 0x73, 0x63, 0x68, 0xc3, 0xbc, 0x73},
		),
		Entry(
			"4 octet length, length larger than 2 octets",
			String4(strings.Repeat(" ", math.MaxUint16+1)),
			append(
				[]byte{0x0e, 0x00, 0x00, 0x01, 0x00},
				bytes.Repeat([]byte{' '}, math.MaxUint16+1)...,
			),
		),
		Entry(
			"8 octet length, empty",
			String8(""),
			[]byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		),
		Entry(
			"8 octet length, ASCII",
			String8("Hello!"),
			[]byte{0x0f, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x21}),
		Entry(
			"8 octet length, UTF-8",
			String8("Tschüs"),
			[]byte{0x0f, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x73, 0x63, 0x68, 0xc3, 0xbc, 0x73},
		),
		// Note, we don't check for lengths larger than 4 octets because
		// allocating a string of that length takes too long (~50s on M1 MBP).
	)
})

package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes signed ingegers correctly",
		func(v Value, expect []byte) {
			data, err := Marshal(Root{Value: v})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expect))
		},
		Entry(
			"1 octet, positive",
			Signed1(42),
			[]byte{0x00, 0x2a},
		),
		Entry(
			"1 octet, negative",
			Signed1(-17),
			[]byte{0x00, 0xef},
		),
		Entry(
			"2 octet, positive",
			Signed2(420),
			[]byte{0x01, 0xa4, 0x01},
		),
		Entry(
			"2 octet, negative",
			Signed2(-1700),
			[]byte{0x01, 0x5c, 0xf9},
		),
		Entry(
			"4 octet, positive",
			Signed4(42000),
			[]byte{0x02, 0x10, 0xa4, 0x00, 0x00},
		),
		Entry(
			"4 octet, negative",
			Signed4(-170000),
			[]byte{0x02, 0xf0, 0x67, 0xfd, 0xff},
		),
		Entry(
			"8 octet, positive",
			Signed8(40000000000),
			[]byte{0x03, 0x00, 0x90, 0x2f, 0x50, 0x09, 0x00, 0x00, 0x00},
		),
		Entry(
			"8 octet, negative",
			Signed8(-170000000000),
			[]byte{0x03, 0x00, 0xdc, 0x35, 0x6b, 0xd8, 0xff, 0xff, 0xff},
		),
	)
})

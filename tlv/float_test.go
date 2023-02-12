package tlv_test

import (
	"math"

	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Marshal()", func() {
	DescribeTable(
		"it encodes floating-point values correctly",
		func(v Value, expect []byte) {
			data, err := Marshal(Root{Value: v})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal(expect))
		},
		Entry(
			"single-precision, zero",
			Float4(0),
			[]byte{0x0a, 0x00, 0x00, 0x00, 0x00},
		),
		Entry(
			"single-precision, 1.0 / 3.0",
			Float4(1.0/3.0),
			[]byte{0x0a, 0xab, 0xaa, 0xaa, 0x3e},
		),
		Entry(
			"single-precision, 17.9",
			Float4(17.9),
			[]byte{0x0a, 0x33, 0x33, 0x8f, 0x41},
		),
		Entry(
			"single-precision, positive infinity",
			Float4(math.Float32frombits(0x7f800000)),
			[]byte{0x0a, 0x00, 0x00, 0x80, 0x7f},
		),
		Entry(
			"single-precision, negative infinity",
			Float4(math.Float32frombits(0xff800000)),
			[]byte{0x0a, 0x00, 0x00, 0x80, 0xff},
		),
		Entry(
			"double-precision, zero",
			Float8(0),
			[]byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		),
		Entry(
			"double-precision, 1.0 / 3.0",
			Float8(1.0/3.0),
			[]byte{0x0b, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0xd5, 0x3f},
		),
		Entry(
			"double-precision, 17.9",
			Float8(17.9),
			[]byte{0x0b, 0x66, 0x66, 0x66, 0x66, 0x66, 0xe6, 0x31, 0x40},
		),
		Entry(
			"double-precision, positive infinity",
			Float8(math.Inf(+1)),
			[]byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x7f},
		),
		Entry(
			"double-precision, negative infinity",
			Float8(math.Inf(-1)),
			[]byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xff},
		),
	)
})
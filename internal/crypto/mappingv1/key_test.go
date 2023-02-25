package mappingv1_test

import (
	. "github.com/jmalloc/motif/internal/crypto/mappingv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func DeriveKey()", func() {
	It("derives the key using NIST SP 800-56C", func() {
		var (
			inputKey = SymmetricKey{0x6e, 0xe6, 0xc0, 0x0d, 0x70, 0xa6, 0xcd, 0x14, 0xbd, 0x5a, 0x4e, 0x8f, 0xcf, 0xec, 0x83, 0x86}
			salt     = []byte{0x53, 0x2f, 0x51, 0x31, 0xe0, 0xa2, 0xfe, 0xcc, 0x72, 0x2f, 0x87, 0xe5, 0xaa, 0x20, 0x62, 0xcb}
			info     = []byte{0x86, 0x1a, 0xa2, 0x88, 0x67, 0x98, 0x23, 0x12, 0x59, 0xbd, 0x03, 0x14}
			expected = SymmetricKey{0x13, 0x47, 0x9e, 0x9a, 0x91, 0xdd, 0x20, 0xfd, 0xd7, 0x57, 0xd6, 0x8f, 0xfe, 0x88, 0x69, 0xfb}
		)

		actual := DeriveKey(
			inputKey,
			salt,
			info,
		)
		Expect(actual).To(Equal(expected))
	})
})
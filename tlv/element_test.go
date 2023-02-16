package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Element", func() {
	It("has a meaningful zero-value (anonymous, null)", func() {
		var e Element
		d, err := e.MarshalBinary()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(d).To(Equal([]byte{0b000_10100}))
	})
})

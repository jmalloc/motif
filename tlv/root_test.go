package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Root", func() {
	It("has a meaningful zero-value", func() {
		var r Root
		data, err := r.MarshalBinary()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(data).To(Equal([]byte{0b000_10100}))
	})
})

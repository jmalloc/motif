package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Root", func() {
	It("has a meaningful zero-value", func() {
		var r Root
		Expect(r.Tag()).To(Equal(AnonymousTag))
		Expect(r.Value()).To(Equal(Null))
	})
})
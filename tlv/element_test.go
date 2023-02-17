package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Element", func() {
	Describe("MarshalBinary()", func() {
		It("encodes the zero-value as an anonymous null TLV element", func() {
			var e Element
			data, err := e.MarshalBinary()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).To(Equal([]byte{0x14}))
		})
	})
})

func testScalar(v Value, data []byte) {
	m := Element{
		T: AnonymousTag,
		V: v,
	}
	d, err := m.MarshalBinary()
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	var u Element
	err = u.UnmarshalBinary(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u).To(Equal(m))
}

func testContainer[
	T interface {
		Value
		~[]M
	},
	M any,
](v T, data []byte) {
	m := Element{
		T: AnonymousTag,
		V: v,
	}
	d, err := m.MarshalBinary()
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	var u Element
	err = u.UnmarshalBinary(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u.T).To(Equal(m.T))

	if len(v) == 0 {
		Expect(u.V).To(HaveLen(0))
	} else {
		Expect(u.V).To(Equal(m.V))
	}
}

func testTag(t Tag, data []byte) {
	m := Element{
		T: t,
		V: Signed1(0),
	}
	d, err := m.MarshalBinary()
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	var u Element
	err = u.UnmarshalBinary(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u).To(Equal(m))
}

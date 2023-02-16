package tlv_test

import (
	. "github.com/jmalloc/motif/tlv"
	. "github.com/onsi/gomega"
)

func testScalarEncoding(v Value, data []byte) {
	m := Root{
		T: AnonymousTag,
		V: v,
	}
	d, err := m.MarshalBinary()
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	u := Root{}
	err = u.UnmarshalBinary(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u).To(Equal(m))
}

func testContainerEncoding[
	T interface {
		Value
		~[]E
	},
	E any,
](v T, data []byte) {
	m := Root{
		T: AnonymousTag,
		V: v,
	}
	d, err := m.MarshalBinary()
	Expect(err).ShouldNot(HaveOccurred())
	Expect(data).To(Equal(data))

	u := Root{}
	err = u.UnmarshalBinary(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u.T).To(Equal(m.T))

	if len(v) == 0 {
		Expect(u.V).To(HaveLen(0))
	} else {
		Expect(u.V).To(Equal(m.V))
	}
}

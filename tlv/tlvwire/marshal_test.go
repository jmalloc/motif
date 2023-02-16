package tlvwire_test

import (
	"github.com/jmalloc/motif/tlv"
	. "github.com/jmalloc/motif/tlv/tlvwire"
	. "github.com/onsi/gomega"
)

func testScalar(v tlv.Value, data []byte) {
	m := tlv.Element{
		T: tlv.AnonymousTag,
		V: v,
	}
	d, err := Marshal(m)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	u, err := Unmarshal(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u).To(Equal(m))
}

func testContainer[
	T interface {
		tlv.Value
		~[]M
	},
	M any,
](v T, data []byte) {
	m := tlv.Element{
		T: tlv.AnonymousTag,
		V: v,
	}
	d, err := Marshal(m)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	u, err := Unmarshal(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u.T).To(Equal(m.T))

	if len(v) == 0 {
		Expect(u.V).To(HaveLen(0))
	} else {
		Expect(u.V).To(Equal(m.V))
	}
}

func testTag(t tlv.Tag, data []byte) {
	m := tlv.Element{
		T: t,
		V: tlv.Signed1(0),
	}
	d, err := Marshal(m)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(d).To(Equal(data))

	u, err := Unmarshal(d)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(u).To(Equal(m))
}

package tlv

import (
	"bytes"
	"math"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Marshal returns the binary representation of e.
func Marshal(e Element) ([]byte, error) {
	w := &bytes.Buffer{}
	marshal(w, e)
	return w.Bytes(), nil
}

func marshal(w *bytes.Buffer, e Element) {
	start := w.Len()
	w.WriteByte(0)

	m := marshaler{w, start}
	t, v := e.Components()
	t.AcceptVisitor(m)
	v.AcceptVisitor(m)
}

type marshaler struct {
	*bytes.Buffer
	start int
}

func (m marshaler) WriteControl(c byte) {
	m.Bytes()[m.start] |= c
}

func (m marshaler) VisitSigned1(s Signed1) {
	m.WriteControl(signed1Type)
	writeInt(m, s)
}

func (m marshaler) VisitSigned2(s Signed2) {
	m.WriteControl(signed2Type)
	writeInt(m, s)
}

func (m marshaler) VisitSigned4(s Signed4) {
	m.WriteControl(signed4Type)
	writeInt(m, s)
}

func (m marshaler) VisitSigned8(s Signed8) {
	m.WriteControl(signed8Type)
	writeInt(m, s)
}

func (m marshaler) VisitUnsigned1(u Unsigned1) {
	m.WriteControl(unsigned1Type)
	writeInt(m, u)
}

func (m marshaler) VisitUnsigned2(u Unsigned2) {
	m.WriteControl(unsigned2Type)
	writeInt(m, u)
}

func (m marshaler) VisitUnsigned4(u Unsigned4) {
	m.WriteControl(unsigned4Type)
	writeInt(m, u)
}

func (m marshaler) VisitUnsigned8(u Unsigned8) {
	m.WriteControl(unsigned8Type)
	writeInt(m, u)
}

func (m marshaler) VisitBool(b Bool) {
	if b {
		m.WriteControl(boolTrueType)
	} else {
		m.WriteControl(boolFalseType)
	}
}

func (m marshaler) VisitFloat4(f Float4) {
	m.WriteControl(float4Type)
	writeInt(m, math.Float32bits(float32(f)))

}

func (m marshaler) VisitFloat8(f Float8) {
	m.WriteControl(float8Type)
	writeInt(m, math.Float64bits(float64(f)))
}

func (m marshaler) VisitNull() {
	m.WriteControl(nullType)
}

func (m marshaler) VisitStruct(s Struct) {
	m.WriteControl(structType)
	for _, x := range s {
		marshal(m.Buffer, x)
	}
	m.WriteByte(endOfContainer)
}

func (m marshaler) VisitArray(a Array) {
	m.WriteControl(arrayType)
	for _, x := range a {
		marshal(m.Buffer, x)
	}
	m.WriteByte(endOfContainer)
}

func (m marshaler) VisitList(l List) {
	m.WriteControl(listType)
	for _, x := range l {
		marshal(m.Buffer, x)
	}
	m.WriteByte(endOfContainer)
}

func (m marshaler) VisitString1(s String1) {
	m.WriteControl(utf8String1Type)
	writeInt(m, uint8(len(s)))
	m.WriteString(string(s))
}

func (m marshaler) VisitString2(s String2) {
	m.WriteControl(utf8String2Type)
	writeInt(m, uint16(len(s)))
	m.WriteString(string(s))
}

func (m marshaler) VisitString4(s String4) {
	m.WriteControl(utf8String4Type)
	writeInt(m, uint32(len(s)))
	m.WriteString(string(s))
}

func (m marshaler) VisitString8(s String8) {
	m.WriteControl(utf8String8Type)
	writeInt(m, uint64(len(s)))
	m.WriteString(string(s))
}

func (m marshaler) VisitBytes1(b Bytes1) {
	m.WriteControl(octetString1Type)
	writeInt(m, uint8(len(b)))
	m.Write(b)
}

func (m marshaler) VisitBytes2(b Bytes2) {
	m.WriteControl(octetString2Type)
	writeInt(m, uint16(len(b)))
	m.Write(b)
}

func (m marshaler) VisitBytes4(b Bytes4) {
	m.WriteControl(octetString4Type)
	writeInt(m, uint32(len(b)))
	m.Write(b)
}

func (m marshaler) VisitBytes8(b Bytes8) {
	m.WriteControl(octetString8Type)
	writeInt(m, uint64(len(b)))
	m.Write(b)
}

func (m marshaler) VisitAnonymousTag() {
}

func (m marshaler) VisitContextSpecificTag(t ContextSpecificTag) {
	m.WriteControl(contextSpecificTagForm)
	writeInt(m, t)
}

func (m marshaler) VisitCommonProfileTag2(t CommonProfileTag2) {
	m.WriteControl(commonProfileTag2Form)
	writeInt(m, t)
}

func (m marshaler) VisitCommonProfileTag4(t CommonProfileTag4) {
	m.WriteControl(commonProfileTag4Form)
	writeInt(m, t)
}

func (m marshaler) VisitImplicitProfileTag2(t ImplicitProfileTag2) {
	m.WriteControl(implicitProfileTag2Form)
	writeInt(m, t)
}

func (m marshaler) VisitImplicitProfileTag4(t ImplicitProfileTag4) {
	m.WriteControl(implicitProfileTag4Form)
	writeInt(m, t)
}

func (m marshaler) VisitFullyQualifiedTag6(t FullyQualifiedTag6) {
	m.WriteControl(fullyQualifiedTag6Form)
	writeInt(m, t.VendorID)
	writeInt(m, t.Profile)
	writeInt(m, t.Tag)
}

func (m marshaler) VisitFullyQualifiedTag8(t FullyQualifiedTag8) {
	m.WriteControl(fullyQualifiedTag8Form)
	writeInt(m, t.VendorID)
	writeInt(m, t.Profile)
	writeInt(m, t.Tag)
}

func writeInt[T constraints.Integer](m marshaler, n T) {
	size := int(unsafe.Sizeof(n))
	for i := 0; i < size; i++ {
		m.WriteByte(byte(n >> (8 * i)))
	}
}

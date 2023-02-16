package tlv

import (
	"io"
	"math"

	"github.com/jmalloc/motif/internal/wire"
)

// marshal returns the binary representation of an element consisting of t and v.
func marshal(w io.Writer, t Tag, v Value) error {
	c := &controlOctetBuilder{}

	if err := t.AcceptVisitor(c); err != nil {
		return err
	}

	if err := v.AcceptVisitor(c); err != nil {
		return err
	}

	if err := wire.WriteInt(w, c.Value); err != nil {
		return err
	}

	m := marshaler{w}

	if err := t.AcceptVisitor(m); err != nil {
		return err
	}

	if err := v.AcceptVisitor(m); err != nil {
		return err
	}

	return nil
}

type marshaler struct {
	io.Writer
}

func (m marshaler) VisitSigned1(s Signed1) error {
	return wire.WriteInt(m, s)
}

func (m marshaler) VisitSigned2(s Signed2) error {
	return wire.WriteInt(m, s)
}

func (m marshaler) VisitSigned4(s Signed4) error {
	return wire.WriteInt(m, s)
}

func (m marshaler) VisitSigned8(s Signed8) error {
	return wire.WriteInt(m, s)
}

func (m marshaler) VisitUnsigned1(u Unsigned1) error {
	return wire.WriteInt(m, u)
}

func (m marshaler) VisitUnsigned2(u Unsigned2) error {
	return wire.WriteInt(m, u)
}

func (m marshaler) VisitUnsigned4(u Unsigned4) error {
	return wire.WriteInt(m, u)
}

func (m marshaler) VisitUnsigned8(u Unsigned8) error {
	return wire.WriteInt(m, u)
}

func (m marshaler) VisitBool(b Bool) error {
	return nil
}

func (m marshaler) VisitFloat4(f Float4) error {
	return wire.WriteInt(
		m,
		math.Float32bits(float32(f)),
	)
}

func (m marshaler) VisitFloat8(f Float8) error {
	return wire.WriteInt(
		m,
		math.Float64bits(float64(f)),
	)
}

func (m marshaler) VisitNull() error {
	return nil
}

func (m marshaler) VisitStruct(s Struct) error {
	for _, e := range s {
		if err := marshal(m, e.T, e.V); err != nil {
			return err
		}
	}

	_, err := m.Write([]byte{endOfContainer})
	return err
}

func (m marshaler) VisitArray(a Array) error {
	for _, e := range a {
		if err := marshal(m, AnonymousTag, e.V); err != nil {
			return err
		}
	}

	_, err := m.Write([]byte{endOfContainer})
	return err
}

func (m marshaler) VisitList(l List) error {
	for _, e := range l {
		if err := marshal(m, e.T, e.V); err != nil {
			return err
		}
	}

	_, err := m.Write([]byte{endOfContainer})
	return err
}

func (m marshaler) VisitString1(s String1) error {
	return wire.WriteString[uint8](m, s)
}

func (m marshaler) VisitString2(s String2) error {
	return wire.WriteString[uint16](m, s)
}

func (m marshaler) VisitString4(s String4) error {
	return wire.WriteString[uint32](m, s)
}

func (m marshaler) VisitString8(s String8) error {
	return wire.WriteString[uint64](m, s)
}

func (m marshaler) VisitBytes1(b Bytes1) error {
	return wire.WriteString[uint8](m, b)
}

func (m marshaler) VisitBytes2(b Bytes2) error {
	return wire.WriteString[uint16](m, b)
}

func (m marshaler) VisitBytes4(b Bytes4) error {
	return wire.WriteString[uint32](m, b)
}

func (m marshaler) VisitBytes8(b Bytes8) error {
	return wire.WriteString[uint64](m, b)
}

func (m marshaler) VisitAnonymousTag() error {
	return nil
}

func (m marshaler) VisitContextSpecificTag(t ContextSpecificTag) error {
	return wire.WriteInt(m, t)
}

func (m marshaler) VisitCommonProfileTag2(t CommonProfileTag2) error {
	return wire.WriteInt(m, t)
}

func (m marshaler) VisitCommonProfileTag4(t CommonProfileTag4) error {
	return wire.WriteInt(m, t)
}

func (m marshaler) VisitImplicitProfileTag2(t ImplicitProfileTag2) error {
	return wire.WriteInt(m, t)
}

func (m marshaler) VisitImplicitProfileTag4(t ImplicitProfileTag4) error {
	return wire.WriteInt(m, t)
}

func (m marshaler) VisitFullyQualifiedTag6(t FullyQualifiedTag6) error {
	if err := wire.WriteInt(m, t.VendorID); err != nil {
		return err
	}

	if err := wire.WriteInt(m, t.Profile); err != nil {
		return err
	}

	if err := wire.WriteInt(m, t.Tag); err != nil {
		return err
	}

	return nil
}

func (m marshaler) VisitFullyQualifiedTag8(t FullyQualifiedTag8) error {
	if err := wire.WriteInt(m, t.VendorID); err != nil {
		return err
	}

	if err := wire.WriteInt(m, t.Profile); err != nil {
		return err
	}

	if err := wire.WriteInt(m, t.Tag); err != nil {
		return err
	}

	return nil
}

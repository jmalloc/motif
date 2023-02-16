package tlv

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/jmalloc/motif/internal/wire"
)

func unmarshal(r *bytes.Reader) (Tag, Value, error) {
	c, err := r.ReadByte()
	if err != nil {
		return nil, nil, err
	}

	t, err := unmarshalTag(r, c)
	if err != nil {
		return nil, nil, err
	}

	v, err := unmarshalValue(r, c)
	if err != nil {
		return nil, nil, err
	}

	return t, v, nil
}

func unmarshalValue(r *bytes.Reader, c byte) (Value, error) {
	switch c & typeMask {
	case signed1Type:
		return wire.ReadInt[Signed1](r)
	case signed2Type:
		return wire.ReadInt[Signed2](r)
	case signed4Type:
		return wire.ReadInt[Signed4](r)
	case signed8Type:
		return wire.ReadInt[Signed8](r)
	case unsigned1Type:
		return wire.ReadInt[Unsigned1](r)
	case unsigned2Type:
		return wire.ReadInt[Unsigned2](r)
	case unsigned4Type:
		return wire.ReadInt[Unsigned4](r)
	case unsigned8Type:
		return wire.ReadInt[Unsigned8](r)
	case boolFalseType:
		return False, nil
	case boolTrueType:
		return True, nil
	case float4Type:
		n, err := wire.ReadInt[uint32](r)
		return Float4(math.Float32frombits(n)), err
	case float8Type:
		n, err := wire.ReadInt[uint64](r)
		return Float8(math.Float64frombits(n)), err
	case nullType:
		return Null, nil
	case structType:
		return unmarshalStruct(r)
	case arrayType:
		return unmarshalArray(r)
	case listType:
		return unmarshalList(r)
	case utf8String1Type:
		return wire.ReadString[uint8, String1](r)
	case utf8String2Type:
		return wire.ReadString[uint16, String2](r)
	case utf8String4Type:
		return wire.ReadString[uint32, String4](r)
	case utf8String8Type:
		return wire.ReadString[uint64, String8](r)
	case octetString1Type:
		return wire.ReadString[uint8, Bytes1](r)
	case octetString2Type:
		return wire.ReadString[uint16, Bytes2](r)
	case octetString4Type:
		return wire.ReadString[uint32, Bytes4](r)
	case octetString8Type:
		return wire.ReadString[uint64, Bytes8](r)
	default:
		return nil, fmt.Errorf("unrecognized type (%x)", c&typeMask)
	}
}

func unmarshalStruct(r *bytes.Reader) (Struct, error) {
	var s Struct

	return s, readMembers(
		r,
		func(t Tag, v Value) error {
			x, ok := t.(NonAnonymousTag)
			if !ok {
				return errors.New("struct members cannot be anonymous")
			}

			s = append(s, StructMember{x, v})
			return nil
		},
	)
}

func unmarshalArray(r *bytes.Reader) (Array, error) {
	var a Array

	return a, readMembers(
		r,
		func(t Tag, v Value) error {
			if t != AnonymousTag {
				return errors.New("array members must be anonymous")
			}

			a = append(a, ArrayMember{v})
			return nil
		},
	)
}

func unmarshalList(r *bytes.Reader) (List, error) {
	var l List

	return l, readMembers(
		r,
		func(t Tag, v Value) error {
			l = append(l, ListMember{t, v})
			return nil
		},
	)
}

func unmarshalTag(r *bytes.Reader, c byte) (Tag, error) {
	switch c & tagFormMask {
	default: // anonymous
		return AnonymousTag, nil
	case contextSpecificTagForm:
		return wire.ReadInt[ContextSpecificTag](r)
	case commonProfileTag2Form:
		return wire.ReadInt[CommonProfileTag2](r)
	case commonProfileTag4Form:
		return wire.ReadInt[CommonProfileTag4](r)
	case implicitProfileTag2Form:
		return wire.ReadInt[ImplicitProfileTag2](r)
	case implicitProfileTag4Form:
		return wire.ReadInt[ImplicitProfileTag4](r)
	case fullyQualifiedTag6Form:
		return unmarshalFullyQualifiedTag6(r)
	case fullyQualifiedTag8Form:
		return unmarshalFullyQualifiedTag8(r)
	}
}

func unmarshalFullyQualifiedTag6(r *bytes.Reader) (FullyQualifiedTag6, error) {
	var (
		t   FullyQualifiedTag6
		err error
	)

	t.VendorID, err = wire.ReadInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Profile, err = wire.ReadInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Tag, err = wire.ReadInt[uint16](r)
	if err != nil {
		return t, err
	}

	return t, nil
}

func unmarshalFullyQualifiedTag8(r *bytes.Reader) (FullyQualifiedTag8, error) {
	var (
		t   FullyQualifiedTag8
		err error
	)

	t.VendorID, err = wire.ReadInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Profile, err = wire.ReadInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Tag, err = wire.ReadInt[uint32](r)
	if err != nil {
		return t, err
	}

	return t, nil
}

func readMembers(
	r *bytes.Reader,
	fn func(Tag, Value) error,
) error {
	for {
		c, err := r.ReadByte()
		if err != nil {
			return err
		}

		if c == endOfContainer {
			return nil
		}

		if err := r.UnreadByte(); err != nil {
			return err
		}

		t, v, err := unmarshal(r)
		if err != nil {
			return err
		}

		if err := fn(t, v); err != nil {
			return err
		}
	}
}

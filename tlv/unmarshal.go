package tlv

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Unmarshal returns the element represented by the given data.
func Unmarshal(data []byte) (Element, error) {
	r := bytes.NewReader(data)

	t, v, err := unmarshal(r)
	if err != nil {
		return nil, err
	}

	if r.Len() != 0 {
		return nil, errors.New("unexpected data remaining")
	}

	return Root{t, v}, nil
}

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
		return readInt[Signed1](r)
	case signed2Type:
		return readInt[Signed2](r)
	case signed4Type:
		return readInt[Signed4](r)
	case signed8Type:
		return readInt[Signed8](r)
	case unsigned1Type:
		return readInt[Unsigned1](r)
	case unsigned2Type:
		return readInt[Unsigned2](r)
	case unsigned4Type:
		return readInt[Unsigned4](r)
	case unsigned8Type:
		return readInt[Unsigned8](r)
	case boolFalseType:
		return False, nil
	case boolTrueType:
		return True, nil
	case float4Type:
		n, err := readInt[uint32](r)
		return Float4(math.Float32frombits(n)), err
	case float8Type:
		n, err := readInt[uint64](r)
		return Float8(math.Float64frombits(n)), err
	case nullType:
		return Null, nil
	case structType:
		return unmarshalStruct(r)
	// case arrayType:
	// 	return unmarshalArray(r)
	// case listType:
	// 	return unmarshalList(r)
	case utf8String1Type:
		return unmarshalString[String1, uint8](r)
	case utf8String2Type:
		return unmarshalString[String2, uint16](r)
	case utf8String4Type:
		return unmarshalString[String4, uint32](r)
	case utf8String8Type:
		return unmarshalString[String8, uint64](r)
	case octetString1Type:
		return unmarshalString[Bytes1, uint8](r)
	case octetString2Type:
		return unmarshalString[Bytes2, uint16](r)
	case octetString4Type:
		return unmarshalString[Bytes4, uint32](r)
	case octetString8Type:
		return unmarshalString[Bytes8, uint64](r)
	default:
		return nil, fmt.Errorf("unrecognized type (%x)", c&typeMask)
	}
}

func unmarshalStruct(r *bytes.Reader) (Struct, error) {
	var s Struct

	for {
		c, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		if c == endOfContainer {
			return s, nil
		}

		if err := r.UnreadByte(); err != nil {
			return nil, err
		}

		t, v, err := unmarshal(r)
		if err != nil {
			return nil, err
		}

		s = append(s, StructMember{t, v})
	}
}

// func unmarshalArray(r *bytes.Reader) (Array, error) {
// 	panic("unimplemented")
// }

// func unmarshalList(r *bytes.Reader) (List, error) {
// 	panic("unimplemented")
// }

func unmarshalString[S ~string | ~[]byte, L constraints.Unsigned](r *bytes.Reader) (S, error) {
	var data []byte

	len, err := readInt[L](r)
	if err != nil {
		return S(data), err
	}

	data = make([]byte, len)
	_, err = io.ReadFull(r, data)

	return S(data), err
}

func unmarshalTag(r *bytes.Reader, c byte) (Tag, error) {
	switch c & tagFormMask {
	default: // anonymous
		return AnonymousTag, nil
	case contextSpecificTagForm:
		return readInt[ContextSpecificTag](r)
	case commonProfileTag2Form:
		return readInt[CommonProfileTag2](r)
	case commonProfileTag4Form:
		return readInt[CommonProfileTag4](r)
	case implicitProfileTag2Form:
		return readInt[ImplicitProfileTag2](r)
	case implicitProfileTag4Form:
		return readInt[ImplicitProfileTag4](r)
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

	t.VendorID, err = readInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Profile, err = readInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Tag, err = readInt[uint16](r)
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

	t.VendorID, err = readInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Profile, err = readInt[uint16](r)
	if err != nil {
		return t, err
	}

	t.Tag, err = readInt[uint32](r)
	if err != nil {
		return t, err
	}

	return t, nil
}

func readInt[T constraints.Integer](r *bytes.Reader) (T, error) {
	var v T
	size := int(unsafe.Sizeof(v))

	for i := 0; i < size; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return v, err
		}

		v |= T(b) << (8 * i)
	}

	return v, nil
}

package tlv

import (
	"bytes"
	"errors"
	"io"
)

type (
	// Struct is a TLV structure element.
	Struct []StructMember

	// StructMember is an element that is a member of a structure.
	StructMember struct {
		T NonAnonymousTag
		V Value
	}

	// Array is a TLV array element.
	Array []Value

	// List is a TLV list element.
	List []ListMember

	// ListMember is an element that is a member of a list.
	ListMember struct {
		T Tag
		V Value
	}
)

func (v Struct) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitStruct(v)
}

func (v Array) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitArray(v)
}

func (v List) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitList(v)
}

const (
	structType     = 0b000_10101
	arrayType      = 0b000_10110
	listType       = 0b000_10111
	endOfContainer = 0b000_11000
)

var (
	endOfContainerBuffer = []byte{endOfContainer}
)

func (w *controlWriter) VisitStruct(v Struct) error {
	return w.write(structType)
}

func (w *controlWriter) VisitArray(v Array) error {
	return w.write(arrayType)
}

func (w *controlWriter) VisitList(v List) error {
	return w.write(listType)
}

func (w *payloadWriter) VisitStruct(v Struct) error {
	return marshalContainer(
		w,
		v,
		func(m StructMember) (Tag, Value) {
			return m.T, m.V
		},
	)
}

func (w *payloadWriter) VisitArray(v Array) error {
	return marshalContainer(
		w,
		v,
		func(m Value) (Tag, Value) {
			return AnonymousTag, m
		},
	)
}

func (w *payloadWriter) VisitList(v List) error {
	return marshalContainer(
		w,
		v,
		func(m ListMember) (Tag, Value) {
			return m.T, m.V
		},
	)
}

func unmarshalStruct(r *bytes.Reader) (Struct, error) {
	return unmarshalContainer[Struct](
		r,
		func(t Tag, v Value) (StructMember, error) {
			if t, ok := t.(NonAnonymousTag); ok {
				return StructMember{T: t, V: v}, nil
			}
			return StructMember{}, errors.New("struct members cannot be anonymous")
		},
	)
}

func unmarshalArray(r *bytes.Reader) (Array, error) {
	return unmarshalContainer[Array](
		r,
		func(t Tag, v Value) (Value, error) {
			if t == AnonymousTag {
				return v, nil
			}
			return nil, errors.New("array members must be anonymous")
		},
	)
}

func unmarshalList(r *bytes.Reader) (List, error) {
	return unmarshalContainer[List](
		r,
		func(t Tag, v Value) (ListMember, error) {
			return ListMember{T: t, V: v}, nil
		},
	)
}

func marshalContainer[T ~[]M, M any](
	w io.Writer,
	v T,
	fn func(M) (Tag, Value),
) error {
	for _, m := range v {
		mt, mv := fn(m)

		if err := marshalElement(w, mt, mv); err != nil {
			return err
		}
	}

	_, err := w.Write(endOfContainerBuffer)
	return err
}

func unmarshalContainer[T ~[]M, M any](
	r *bytes.Reader,
	fn func(Tag, Value) (M, error),
) (T, error) {
	var v T

	for {
		octet, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		if octet == endOfContainer {
			return v, nil
		}

		if err := r.UnreadByte(); err != nil {
			return nil, err
		}

		mt, mv, err := unmarshalElement(r)
		if err != nil {
			return nil, err
		}

		m, err := fn(mt, mv)
		if err != nil {
			return nil, err
		}

		v = append(v, m)
	}
}

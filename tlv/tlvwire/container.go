package tlvwire

import (
	"bytes"
	"errors"
	"io"

	"github.com/jmalloc/motif/tlv"
)

const (
	structType     = 0b000_10101
	arrayType      = 0b000_10110
	listType       = 0b000_10111
	endOfContainer = 0b000_11000
)

var (
	endOfContainerBuffer = []byte{endOfContainer}
)

func (w *controlWriter) VisitStruct(v tlv.Struct) error {
	return w.write(structType)
}

func (w *controlWriter) VisitArray(v tlv.Array) error {
	return w.write(arrayType)
}

func (w *controlWriter) VisitList(v tlv.List) error {
	return w.write(listType)
}

func (w *payloadWriter) VisitStruct(v tlv.Struct) error {
	return marshalContainer(
		w,
		v,
		func(m tlv.StructMember) (tlv.Tag, tlv.Value) {
			return m.T, m.V
		},
	)
}

func (w *payloadWriter) VisitArray(v tlv.Array) error {
	return marshalContainer(
		w,
		v,
		func(m tlv.Value) (tlv.Tag, tlv.Value) {
			return tlv.AnonymousTag, m
		},
	)
}

func (w *payloadWriter) VisitList(v tlv.List) error {
	return marshalContainer(
		w,
		v,
		func(m tlv.ListMember) (tlv.Tag, tlv.Value) {
			return m.T, m.V
		},
	)
}

func unmarshalStruct(r *bytes.Reader) (tlv.Struct, error) {
	return unmarshalContainer[tlv.Struct](
		r,
		func(t tlv.Tag, v tlv.Value) (tlv.StructMember, error) {
			if t, ok := t.(tlv.NonAnonymousTag); ok {
				return tlv.StructMember{T: t, V: v}, nil
			}
			return tlv.StructMember{}, errors.New("struct members cannot be anonymous")
		},
	)
}

func unmarshalArray(r *bytes.Reader) (tlv.Array, error) {
	return unmarshalContainer[tlv.Array](
		r,
		func(t tlv.Tag, v tlv.Value) (tlv.Value, error) {
			if t == tlv.AnonymousTag {
				return v, nil
			}
			return nil, errors.New("array members must be anonymous")
		},
	)
}

func unmarshalList(r *bytes.Reader) (tlv.List, error) {
	return unmarshalContainer[tlv.List](
		r,
		func(t tlv.Tag, v tlv.Value) (tlv.ListMember, error) {
			return tlv.ListMember{T: t, V: v}, nil
		},
	)
}

func marshalContainer[T ~[]M, M any](
	w io.Writer,
	v T,
	fn func(M) (tlv.Tag, tlv.Value),
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
	fn func(tlv.Tag, tlv.Value) (M, error),
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

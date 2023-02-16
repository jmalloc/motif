package tlvwire

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

// Marshal returns the binary representation of e.
func Marshal(e tlv.Element) ([]byte, error) {
	t := e.T
	if t == nil {
		t = tlv.AnonymousTag
	}

	v := e.V
	if v == nil {
		v = tlv.Null
	}

	w := &bytes.Buffer{}
	if err := marshalElement(w, t, v); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func marshalElement(w io.Writer, t tlv.Tag, v tlv.Value) error {
	var c controlWriter

	if err := tlv.VisitTag(t, &c); err != nil {
		return err
	}

	if err := tlv.VisitValue(v, &c); err != nil {
		return err
	}

	if err := wire.WriteInt(w, c); err != nil {
		return err
	}

	p := payloadWriter{w}

	if err := tlv.VisitTag(t, &p); err != nil {
		return err
	}

	if err := tlv.VisitValue(v, &p); err != nil {
		return err
	}

	return nil
}

type (
	controlWriter byte
	payloadWriter struct{ io.Writer }
)

func (w *controlWriter) write(v byte) error {
	*w |= controlWriter(v)
	return nil
}

// Unmarshal returns the TLV element represented by the given binary data.
func Unmarshal(data []byte) (tlv.Element, error) {
	r := bytes.NewReader(data)

	t, v, err := unmarshalElement(r)
	if err != nil {
		return tlv.Element{}, err
	}

	if r.Len() > 0 {
		return tlv.Element{}, errors.New("unexpected trailing data")
	}

	return tlv.Element{T: t, V: v}, nil
}

const (
	typeMask    = 0b000_11111
	tagFormMask = 0b111_00000
)

func unmarshalElement(r *bytes.Reader) (tlv.Tag, tlv.Value, error) {
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

func unmarshalValue(r *bytes.Reader, c byte) (tlv.Value, error) {
	switch c & typeMask {
	case signed1Type:
		return wire.ReadInt[tlv.Signed1](r)
	case signed2Type:
		return wire.ReadInt[tlv.Signed2](r)
	case signed4Type:
		return wire.ReadInt[tlv.Signed4](r)
	case signed8Type:
		return wire.ReadInt[tlv.Signed8](r)
	case unsigned1Type:
		return wire.ReadInt[tlv.Unsigned1](r)
	case unsigned2Type:
		return wire.ReadInt[tlv.Unsigned2](r)
	case unsigned4Type:
		return wire.ReadInt[tlv.Unsigned4](r)
	case unsigned8Type:
		return wire.ReadInt[tlv.Unsigned8](r)
	case boolFalseType:
		return tlv.False, nil
	case boolTrueType:
		return tlv.True, nil
	case singleType:
		return unmarshalSingle(r)
	case doubleType:
		return unmarshalDouble(r)
	case nullType:
		return tlv.Null, nil
	case structType:
		return unmarshalStruct(r)
	case arrayType:
		return unmarshalArray(r)
	case listType:
		return unmarshalList(r)
	case utf8String1Type:
		return wire.ReadString[uint8, tlv.UTF8String1](r)
	case utf8String2Type:
		return wire.ReadString[uint16, tlv.UTF8String2](r)
	case utf8String4Type:
		return wire.ReadString[uint32, tlv.UTF8String4](r)
	case utf8String8Type:
		return wire.ReadString[uint64, tlv.UTF8String8](r)
	case octetString1Type:
		return wire.ReadString[uint8, tlv.OctetString1](r)
	case octetString2Type:
		return wire.ReadString[uint16, tlv.OctetString2](r)
	case octetString4Type:
		return wire.ReadString[uint32, tlv.OctetString4](r)
	case octetString8Type:
		return wire.ReadString[uint64, tlv.OctetString8](r)
	default:
		return nil, fmt.Errorf("unrecognized type (%x)", c&typeMask)
	}
}

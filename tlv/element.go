package tlv

import (
	"bytes"
	"errors"
	"io"
)

// Element is a TLV element.
type Element struct {
	T Tag
	V Value
}

// MarshalBinary returns the binary representation of e.
func (e Element) MarshalBinary() ([]byte, error) {
	t := e.T
	if t == nil {
		t = AnonymousTag
	}

	v := e.V
	if v == nil {
		v = Null
	}

	w := &bytes.Buffer{}
	if err := marshalElement(w, t, v); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

// UnmarshalBinary sets e to the value represented by data.
func (e *Element) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	t, v, err := unmarshalElement(r)
	if err != nil {
		return err
	}

	if r.Len() > 0 {
		return errors.New("unexpected trailing data")
	}

	e.T = t
	e.V = v

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

func marshalElement(w io.Writer, t Tag, v Value) error {
	var c controlWriter

	if err := VisitTag(t, &c); err != nil {
		return err
	}

	if err := VisitValue(v, &c); err != nil {
		return err
	}

	if err := writeInt(w, c); err != nil {
		return err
	}

	p := payloadWriter{w}

	if err := VisitTag(t, &p); err != nil {
		return err
	}

	if err := VisitValue(v, &p); err != nil {
		return err
	}

	return nil
}

const (
	typeMask    = 0b000_11111
	tagFormMask = 0b111_00000
)

func unmarshalElement(r *bytes.Reader) (Tag, Value, error) {
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

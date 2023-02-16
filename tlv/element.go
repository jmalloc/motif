package tlv

import "bytes"

// Element is a TLV element.
type Element struct {
	T Tag
	V Value
}

// MarshalBinary returns the binary encoding of the element.
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
	if err := marshal(w, t, v); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

// UnmarshalBinary populates the element from its binary encoding.
func (e *Element) UnmarshalBinary(data []byte) error {
	t, v, err := unmarshal(bytes.NewReader(data))
	if err != nil {
		return err
	}

	e.T = t
	e.V = v

	return nil
}

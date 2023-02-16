package tlv

import "bytes"

// Root is the element at the root of a TLV element tree.
type Root struct {
	T Tag
	V Value
}

// MarshalBinary returns the binary encoding of the element.
func (r Root) MarshalBinary() ([]byte, error) {
	t := r.T
	if t == nil {
		t = AnonymousTag
	}

	v := r.V
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
func (r *Root) UnmarshalBinary(data []byte) error {
	t, v, err := unmarshal(bytes.NewReader(data))
	if err != nil {
		return err
	}

	r.T = t
	r.V = v

	return nil
}

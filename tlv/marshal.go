package tlv

import "bytes"

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

package tlv

import "bytes"

// Marshal returns the binary representation of e.
func Marshal(e Element) ([]byte, error) {
	m := &marshaler{}
	e.AcceptElementVisitor(m)
	return m.result.Bytes(), nil
}

type marshaler struct {
	result  bytes.Buffer
	control byte
	payload bytes.Buffer
}

func (m *marshaler) VisitAnonymousElement(v Value) {
	m.control = 0
	m.payload.Reset()

	v.AcceptValueVisitor(m)

	m.result.WriteByte(m.control)
	m.result.Write(m.payload.Bytes())
}

func (m *marshaler) VisitTaggedElement(t Tag, v Value) {
	m.control = 0
	m.payload.Reset()

	t.AcceptTagVisitor(m)
	v.AcceptValueVisitor(m)

	m.result.WriteByte(m.control)
	m.result.Write(m.payload.Bytes())
}

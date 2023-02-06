package tlv

// String is a TLV UTF-8 string.
//
// Matter v1.0 Core § A.11.2
type String string

// AcceptVisitor passes the element to the appropriate method of v.
func (e String) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitString(ctx, e)
}

func (m marshaler) VisitString(ctx VisitContext, e String) error {
	m.Data.WriteByte(0x0c)
	m.Data.WriteByte(byte(len(e)))
	m.Data.WriteString(string(e))
	return nil
}

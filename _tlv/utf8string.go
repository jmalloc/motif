package tlv

// UTF8String represents a Matter TLV "UTF-8 string" value.
type UTF8String uint64

// AcceptVisitor calls v.VisitUTF8String().
func (e UTF8String) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitUTF8String(ctx, e)
}

func (m marshaler) VisitUTF8String(ctx VisitContext, e UTF8String) error {
	panic("not implemented")
}

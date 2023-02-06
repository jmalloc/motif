package tlv

// Float represents a Matter TLV "float point" value.
type Float uint64

// AcceptVisitor calls v.VisitFloat().
func (e Float) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitFloat(ctx, e)
}

func (m marshaler) VisitFloat(ctx VisitContext, e Float) error {
	panic("not implemented")
}

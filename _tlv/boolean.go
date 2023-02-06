package tlv

// Boolean represents a Matter TLV "boolean" value.
type Boolean uint64

// AcceptVisitor calls v.VisitBoolean().
func (e Boolean) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitBoolean(ctx, e)
}

func (m marshaler) VisitBoolean(ctx VisitContext, e Boolean) error {
	panic("not implemented")
}

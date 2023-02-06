package tlv

// Unsigned represents a Matter TLV "unsigned integer" value.
type Unsigned uint64

// AcceptVisitor calls v.VisitUnsigned().
func (e Unsigned) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitUnsigned(ctx, e)
}

func (m marshaler) VisitUnsigned(ctx VisitContext, e Unsigned) error {
	panic("not implemented")
}

package tlv

// OctetString represents a Matter TLV "octet string" value.
type OctetString uint64

// AcceptVisitor calls v.VisitOctetString().
func (e OctetString) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitOctetString(ctx, e)
}

func (m marshaler) VisitOctetString(ctx VisitContext, e OctetString) error {
	panic("not implemented")
}

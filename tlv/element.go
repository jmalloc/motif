package tlv

// Element is an interface for a TLV element.
type Element interface {
	// AcceptVisitor passes the element to the appropriate method of v.
	AcceptVisitor(VisitContext, ElementVisitor) error
}

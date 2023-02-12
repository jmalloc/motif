package tlv

// Element is an interface for a TLV element.
type Element interface {
	AcceptElementVisitor(vis ElementVisitor)
}

// ElementVisitor is an interface for visiting TLV elements.
type ElementVisitor interface {
	VisitAnonymousElement(v Value)
	VisitTaggedElement(t Tag, v Value)
}

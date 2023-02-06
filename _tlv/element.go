package tlv

// Element is an interface for a TLV element.
type Element interface {
	AcceptVisitor(VisitContext, ElementVisitor) error
}

// AnonymousElement is an interface for a TLV element that lacks a tag value
// (also known as having an "anonymous" tag).
type AnonymousElement interface {
	Element

	// WithTag returns a copy of the element with the given tag.
	WithTag(Tag) TaggedElement
}

// TaggedElement is a TLV element that has a tag.
type TaggedElement struct {
	tag     Tag
	element AnonymousElement
}

// AcceptVisitor visits the inner element.
func (e TaggedElement) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return e.element.AcceptVisitor(
		visitContext{
			Context: ctx,
			tag:     e.tag,
		},
		v,
	)
}

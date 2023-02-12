package tlv

// Tag is an interface for a TLV tag.
type Tag interface {
	AcceptTagVisitor(TagVisitor)
}

// TagVisitor is an interface for visiting TLV tags.
type TagVisitor interface {
}

// ApplyTag tags an anonymous element.
func ApplyTag(t Tag, v Value) Element {
	return TaggedElement{t, v}
}

// TaggedElement is an element with a (non-anonymous) tag.
type TaggedElement struct {
	t Tag
	v Value
}

// AcceptElementVisitor invokes the appropriate method on vis.
func (e TaggedElement) AcceptElementVisitor(vis ElementVisitor) {
	vis.VisitTaggedElement(e.t, e.v)
}

package tlv

type Tag interface {
	AcceptTagVisitor(TagVisitor)
}

// TagVisitor is an interface for visiting TLV tags.
type TagVisitor interface {
}

// ApplyTag tags an anonymous element.
func ApplyTag(t Tag, v Value) Element {
	return taggedElement{t, v}
}

type taggedElement struct {
	t Tag
	v Value
}

func (e taggedElement) AcceptElementVisitor(vis ElementVisitor) {
	vis.VisitTaggedElement(e.t, e.v)
}

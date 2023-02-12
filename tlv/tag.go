package tlv

// Tag is an interface for a TLV tag.
type Tag interface {
	AcceptVisitor(TagVisitor)
}

// TagVisitor is an interface for visiting TLV tags.
type TagVisitor interface {
	VisitAnonymousTag()
	VisitContextSpecificTag(ContextSpecificTag)
}

// AnonymousTag is a specifical tag that identifies an element without any tag
// value.
var AnonymousTag Tag = anonymousTag{}

// anonymousTag is a specifical tag that identifies an element without any tag
// value.
type anonymousTag struct{}

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (anonymousTag) AcceptVisitor(v TagVisitor) {
	v.VisitAnonymousTag()
}

// ContextSpecificTag is a tag that identifies an element within the context of
// a specific structure.
type ContextSpecificTag uint8

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t ContextSpecificTag) AcceptVisitor(v TagVisitor) {
	v.VisitContextSpecificTag(t)
}

func (m marshaler) VisitAnonymousTag() {
}

func (m marshaler) VisitContextSpecificTag(t ContextSpecificTag) {
	m.WriteControl(contextSpecificTagForm)
	m.WriteByte(byte(t))
}

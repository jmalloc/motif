package tlv

// A Tag identifiers the purpose of a TLV element.
type Tag interface {
	AcceptVisitor(VisitContext, TagVisitor) error
}

type TagVisitor interface {
	VisitProfileSpecificTag(VisitContext, ProfileSpecificTag) error
	VisitCommonProfileTag(VisitContext, CommonProfileTag) error
	VisitContextSpecificTag(VisitContext, ContextSpecificTag) error
}

type ProfileSpecificTag struct {
	VendorID      uint16
	ProfileNumber uint16
	TagNumber     uint32
}

type CommonProfileTag struct {
	TagNumber uint32
}

type ContextSpecificTag struct {
	TagNumber uint8
}

func (m *marshaler) VisitProfileSpecificTag(ctx VisitContext, t ProfileSpecificTag) error {
	return nil
}

func (m *marshaler) VisitCommonProfileTag(ctx VisitContext, t CommonProfileTag) error {
	return nil
}

func (m *marshaler) VisitContextSpecificTag(ctx VisitContext, t ContextSpecificTag) error {
	return nil
}

// func marshalTag(ctx VisitContext, t Tag, e AnonymousElement) (byte, []byte) {
// 	if t == nil {
// 		return AnonymousTagForm, nil
// 	}

// 	var m tagMarshaler
// 	t.AcceptVisitor(ctx, &m)

// 	return m.form, m.data
// }

// type tagMarshaler struct {
// 	form byte
// 	data []byte
// }

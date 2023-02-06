package tlvwire

type Tag interface {
	acceptVisitor(TagVisitor) error
}

type TagVisitor interface {
	VisitAnonymousTag() error
	VisitContextSpecificTag(ContextSpecificTag) error
	VisitCommonProfileTag2(CommonProfileTag2) error
	VisitCommonProfileTag4(CommonProfileTag4) error
	VisitImplicitProfileTag2(ImplicitProfileTag2) error
	VisitImplicitProfileTag4(ImplicitProfileTag4) error
	VisitFullyQualifiedTag6(FullyQualifiedTag6) error
	VisitFullyQualifiedTag8(FullyQualifiedTag8) error
}

func VisitTag(t Tag, v TagVisitor) error {
	if t == nil {
		return v.VisitAnonymousTag()
	}

	return t.acceptVisitor(v)
}

type ContextSpecificTag struct {
	TagNumber uint8
}

func (t ContextSpecificTag) AcceptVisitor(v TagVisitor) error {
	return v.VisitContextSpecificTag(t)
}

type CommonProfileTag2 struct {
	TagNumber uint16
}

func (t CommonProfileTag2) AcceptVisitor(v TagVisitor) error {
	return v.VisitCommonProfileTag2(t)
}

type CommonProfileTag4 struct {
	TagNumber uint32
}

func (t CommonProfileTag4) AcceptVisitor(v TagVisitor) error {
	return v.VisitCommonProfileTag4(t)
}

type ImplicitProfileTag2 struct {
	TagNumber uint16
}

func (t ImplicitProfileTag2) AcceptVisitor(v TagVisitor) error {
	return v.VisitImplicitProfileTag2(t)
}

type ImplicitProfileTag4 struct {
	TagNumber uint32
}

func (t ImplicitProfileTag4) AcceptVisitor(v TagVisitor) error {
	return v.VisitImplicitProfileTag4(t)
}

type FullyQualifiedTag6 struct {
	VendorID      uint16
	ProfileNumber uint16
	TagNumber     uint16
}

func (t FullyQualifiedTag6) AcceptVisitor(v TagVisitor) error {
	return v.VisitFullyQualifiedTag6(t)
}

type FullyQualifiedTag8 struct {
	VendorID      uint16
	ProfileNumber uint16
	TagNumber     uint32
}

func (t FullyQualifiedTag8) AcceptVisitor(v TagVisitor) error {
	return v.VisitFullyQualifiedTag8(t)
}

// TagForm is an enumeration of the different forms in which a tag can be
// encoded.
type TagForm byte

const (
	AnonymousTagForm        TagForm = 0b000
	ContextSpecificTagForm  TagForm = 0b001
	CommonProfileTag2Form   TagForm = 0b010
	CommonProfileTag4Form   TagForm = 0b011
	ImplicitProfileTag2Form TagForm = 0b100
	ImplicitProfileTag4Form TagForm = 0b101
	FullyQualifiedTag6Form  TagForm = 0b110
	FullyQualifiedTag8Form  TagForm = 0b111
)

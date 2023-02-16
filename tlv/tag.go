package tlv

// Tag is an interface for a TLV tag.
type Tag interface {
	acceptVisitor(TagVisitor) error
}

// NonAnonymousTag is an interface for a TLV tag that is not AnonymousTag.
type NonAnonymousTag interface {
	Tag

	isNotAnonymous()
}

// TagVisitor is an interface for visiting TLV tags.
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

// VisitTag visits a tag with a visitor.
func VisitTag(t Tag, vis TagVisitor) error {
	return t.acceptVisitor(vis)
}

// AnonymousTag is a specifical tag that identifies an element without any tag
// value.
var AnonymousTag anonymousTag

type (
	// anonymousTag is a specifical tag that identifies an element without any
	// tag value.
	anonymousTag struct{}

	// ContextSpecificTag is a tag that identifies an element within the context
	// of a specific structure.
	ContextSpecificTag uint8

	// CommonProfileTag2 is a tag defined in the Matter Common Profile
	// represented as 2 octets.
	CommonProfileTag2 uint16

	// CommonProfileTag4 is a tag defined in the Matter Common Profile
	// represented as 4 octets.
	CommonProfileTag4 uint32

	// ImplicitProfileTag2 is a profile-specific tag with an implied vendor ID
	// and profile number based on the context in which it is used. It is
	// represented as 2 octets.
	ImplicitProfileTag2 uint16

	// ImplicitProfileTag4 is a profile-specific tag with an implied vendor ID
	// and profile number based on the context in which it is used. It is
	// represented as 4 octets.
	ImplicitProfileTag4 uint32

	// FullyQualifiedTag6 is a tag that is fully qualified with a vendor ID and
	// profile number. It is represented as 6 octets.
	FullyQualifiedTag6 struct {
		VendorID uint16
		Profile  uint16
		Tag      uint16
	}

	// FullyQualifiedTag8 is a tag that is fully qualified with a vendor ID and
	// profile number. It is represented as 8 octets.
	FullyQualifiedTag8 struct {
		VendorID uint16
		Profile  uint16
		Tag      uint32
	}
)

func (anonymousTag) acceptVisitor(vis TagVisitor) error {
	return vis.VisitAnonymousTag()
}

func (t ContextSpecificTag) acceptVisitor(vis TagVisitor) error {
	return vis.VisitContextSpecificTag(t)
}

func (t CommonProfileTag2) acceptVisitor(vis TagVisitor) error {
	return vis.VisitCommonProfileTag2(t)
}

func (t CommonProfileTag4) acceptVisitor(vis TagVisitor) error {
	return vis.VisitCommonProfileTag4(t)
}

func (t ImplicitProfileTag2) acceptVisitor(vis TagVisitor) error {
	return vis.VisitImplicitProfileTag2(t)
}

func (t ImplicitProfileTag4) acceptVisitor(vis TagVisitor) error {
	return vis.VisitImplicitProfileTag4(t)
}

func (t FullyQualifiedTag6) acceptVisitor(vis TagVisitor) error {
	return vis.VisitFullyQualifiedTag6(t)
}

func (t FullyQualifiedTag8) acceptVisitor(vis TagVisitor) error {
	return vis.VisitFullyQualifiedTag8(t)
}

func (ContextSpecificTag) isNotAnonymous()  {}
func (CommonProfileTag2) isNotAnonymous()   {}
func (CommonProfileTag4) isNotAnonymous()   {}
func (ImplicitProfileTag2) isNotAnonymous() {}
func (ImplicitProfileTag4) isNotAnonymous() {}
func (FullyQualifiedTag6) isNotAnonymous()  {}
func (FullyQualifiedTag8) isNotAnonymous()  {}

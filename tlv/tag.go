package tlv

// Tag is an interface for a TLV tag.
type Tag interface {
	AcceptVisitor(TagVisitor)
}

// NonAnonymousTag is an interface for a TLV tag that is not AnonymousTag.
type NonAnonymousTag interface {
	Tag
	isNotAnonymous()
}

// TagVisitor is an interface for visiting TLV tags.
type TagVisitor interface {
	VisitAnonymousTag()
	VisitContextSpecificTag(ContextSpecificTag)
	VisitCommonProfileTag2(CommonProfileTag2)
	VisitCommonProfileTag4(CommonProfileTag4)
	VisitImplicitProfileTag2(ImplicitProfileTag2)
	VisitImplicitProfileTag4(ImplicitProfileTag4)
	VisitFullyQualifiedTag6(FullyQualifiedTag6)
	VisitFullyQualifiedTag8(FullyQualifiedTag8)
}

// AnonymousTag is a specifical tag that identifies an element without any tag
// value.
const AnonymousTag anonymousTag = 0

type (
	// anonymousTag is a specifical tag that identifies an element without any
	// tag value.
	anonymousTag uint8

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

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (anonymousTag) AcceptVisitor(v TagVisitor) { v.VisitAnonymousTag() }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t ContextSpecificTag) AcceptVisitor(v TagVisitor) { v.VisitContextSpecificTag(t) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t CommonProfileTag2) AcceptVisitor(v TagVisitor) { v.VisitCommonProfileTag2(t) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t CommonProfileTag4) AcceptVisitor(v TagVisitor) { v.VisitCommonProfileTag4(t) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t ImplicitProfileTag2) AcceptVisitor(v TagVisitor) { v.VisitImplicitProfileTag2(t) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t ImplicitProfileTag4) AcceptVisitor(v TagVisitor) { v.VisitImplicitProfileTag4(t) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t FullyQualifiedTag6) AcceptVisitor(v TagVisitor) { v.VisitFullyQualifiedTag6(t) }

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (t FullyQualifiedTag8) AcceptVisitor(v TagVisitor) { v.VisitFullyQualifiedTag8(t) }

func (ContextSpecificTag) isNotAnonymous()  {}
func (CommonProfileTag2) isNotAnonymous()   {}
func (CommonProfileTag4) isNotAnonymous()   {}
func (ImplicitProfileTag2) isNotAnonymous() {}
func (ImplicitProfileTag4) isNotAnonymous() {}
func (FullyQualifiedTag6) isNotAnonymous()  {}
func (FullyQualifiedTag8) isNotAnonymous()  {}

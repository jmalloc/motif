package tlv

// Tag is an interface for a TLV tag.
type Tag interface {
	AcceptVisitor(TagVisitor)
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

func (m marshaler) VisitAnonymousTag() {
}

func (m marshaler) VisitContextSpecificTag(t ContextSpecificTag) {
	m.WriteControl(contextSpecificTagForm)
	m.WriteByte(byte(t))
}

func (m marshaler) VisitCommonProfileTag2(t CommonProfileTag2) {
	m.WriteControl(commonProfileTag2Form)
	m.Grow(2)
	m.WriteByte(byte(t))
	m.WriteByte(byte(t >> 8))
}

func (m marshaler) VisitCommonProfileTag4(t CommonProfileTag4) {
	m.WriteControl(commonProfileTag4Form)
	m.Grow(4)
	m.WriteByte(byte(t))
	m.WriteByte(byte(t >> 8))
	m.WriteByte(byte(t >> 16))
	m.WriteByte(byte(t >> 24))
}

func (m marshaler) VisitImplicitProfileTag2(t ImplicitProfileTag2) {
	m.WriteControl(implicitProfileTag2Form)
	m.Grow(2)
	m.WriteByte(byte(t))
	m.WriteByte(byte(t >> 8))
}

func (m marshaler) VisitImplicitProfileTag4(t ImplicitProfileTag4) {
	m.WriteControl(implicitProfileTag4Form)
	m.Grow(4)
	m.WriteByte(byte(t))
	m.WriteByte(byte(t >> 8))
	m.WriteByte(byte(t >> 16))
	m.WriteByte(byte(t >> 24))
}

func (m marshaler) VisitFullyQualifiedTag6(t FullyQualifiedTag6) {
	m.WriteControl(fullyQualifiedTag6Form)
	m.Grow(6)
	m.WriteByte(byte(t.VendorID))
	m.WriteByte(byte(t.VendorID >> 8))
	m.WriteByte(byte(t.Profile))
	m.WriteByte(byte(t.Profile >> 8))
	m.WriteByte(byte(t.Tag))
	m.WriteByte(byte(t.Tag >> 8))
}

func (m marshaler) VisitFullyQualifiedTag8(t FullyQualifiedTag8) {
	m.WriteControl(fullyQualifiedTag8Form)
	m.Grow(8)
	m.WriteByte(byte(t.VendorID))
	m.WriteByte(byte(t.VendorID >> 8))
	m.WriteByte(byte(t.Profile))
	m.WriteByte(byte(t.Profile >> 8))
	m.WriteByte(byte(t.Tag))
	m.WriteByte(byte(t.Tag >> 8))
	m.WriteByte(byte(t.Tag >> 16))
	m.WriteByte(byte(t.Tag >> 24))
}

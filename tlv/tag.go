package tlv

import (
	"bytes"
	"io"
)

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

const (
	anonymousTagForm        = 0b000_00000
	contextSpecificTagForm  = 0b001_00000
	commonProfileTag2Form   = 0b010_00000
	commonProfileTag4Form   = 0b011_00000
	implicitProfileTag2Form = 0b100_00000
	implicitProfileTag4Form = 0b101_00000
	fullyQualifiedTag6Form  = 0b110_00000
	fullyQualifiedTag8Form  = 0b111_00000
)

func (c *controlWriter) VisitAnonymousTag() error {
	return c.write(anonymousTagForm)
}

func (c *controlWriter) VisitContextSpecificTag(t ContextSpecificTag) error {
	return c.write(contextSpecificTagForm)
}

func (c *controlWriter) VisitCommonProfileTag2(t CommonProfileTag2) error {
	return c.write(commonProfileTag2Form)
}

func (c *controlWriter) VisitCommonProfileTag4(t CommonProfileTag4) error {
	return c.write(commonProfileTag4Form)
}

func (c *controlWriter) VisitImplicitProfileTag2(t ImplicitProfileTag2) error {
	return c.write(implicitProfileTag2Form)
}

func (c *controlWriter) VisitImplicitProfileTag4(t ImplicitProfileTag4) error {
	return c.write(implicitProfileTag4Form)
}

func (c *controlWriter) VisitFullyQualifiedTag6(t FullyQualifiedTag6) error {
	return c.write(fullyQualifiedTag6Form)
}

func (c *controlWriter) VisitFullyQualifiedTag8(t FullyQualifiedTag8) error {
	return c.write(fullyQualifiedTag8Form)
}

func (w *payloadWriter) VisitAnonymousTag() error {
	return nil
}

func (w *payloadWriter) VisitContextSpecificTag(t ContextSpecificTag) error {
	return writeInt(w, t)
}

func (w *payloadWriter) VisitCommonProfileTag2(t CommonProfileTag2) error {
	return writeInt(w, t)
}

func (w *payloadWriter) VisitCommonProfileTag4(t CommonProfileTag4) error {
	return writeInt(w, t)
}

func (w *payloadWriter) VisitImplicitProfileTag2(t ImplicitProfileTag2) error {
	return writeInt(w, t)
}

func (w *payloadWriter) VisitImplicitProfileTag4(t ImplicitProfileTag4) error {
	return writeInt(w, t)
}

func (w *payloadWriter) VisitFullyQualifiedTag6(t FullyQualifiedTag6) error {
	return marshalFullyQualifiedTag(w, t.VendorID, t.Profile, t.Tag)
}

func (w *payloadWriter) VisitFullyQualifiedTag8(t FullyQualifiedTag8) error {
	return marshalFullyQualifiedTag(w, t.VendorID, t.Profile, t.Tag)
}

func marshalFullyQualifiedTag[T uint16 | uint32](
	w io.Writer,
	vendor, profile uint16,
	tag T,
) error {
	if err := writeInt(w, vendor); err != nil {
		return err
	}

	if err := writeInt(w, profile); err != nil {
		return err
	}

	if err := writeInt(w, tag); err != nil {
		return err
	}

	return nil
}

func unmarshalTag(r *bytes.Reader, c byte) (Tag, error) {
	switch c & tagFormMask {
	default:
		return AnonymousTag, nil
	case contextSpecificTagForm:
		return readInt[ContextSpecificTag](r)
	case commonProfileTag2Form:
		return readInt[CommonProfileTag2](r)
	case commonProfileTag4Form:
		return readInt[CommonProfileTag4](r)
	case implicitProfileTag2Form:
		return readInt[ImplicitProfileTag2](r)
	case implicitProfileTag4Form:
		return readInt[ImplicitProfileTag4](r)
	case fullyQualifiedTag6Form:
		return unmarshalFullyQualifiedTag6(r)
	case fullyQualifiedTag8Form:
		return unmarshalFullyQualifiedTag8(r)
	}
}

func unmarshalFullyQualifiedTag6(r *bytes.Reader) (FullyQualifiedTag6, error) {
	var t FullyQualifiedTag6
	return t, unmarshalFullyQualifiedTag(
		r,
		&t.VendorID,
		&t.Profile,
		&t.Tag,
	)
}

func unmarshalFullyQualifiedTag8(r *bytes.Reader) (FullyQualifiedTag8, error) {
	var t FullyQualifiedTag8
	return t, unmarshalFullyQualifiedTag(
		r,
		&t.VendorID,
		&t.Profile,
		&t.Tag,
	)
}

func unmarshalFullyQualifiedTag[T uint16 | uint32](
	r *bytes.Reader,
	vendor, profile *uint16,
	tag *T,
) error {
	if err := assignInt(r, vendor); err != nil {
		return err
	}

	if err := assignInt(r, profile); err != nil {
		return err
	}

	if err := assignInt(r, tag); err != nil {
		return err
	}

	return nil
}

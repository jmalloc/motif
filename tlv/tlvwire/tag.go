package tlvwire

import (
	"bytes"
	"io"

	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

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

func (c *controlWriter) VisitContextSpecificTag(t tlv.ContextSpecificTag) error {
	return c.write(contextSpecificTagForm)
}
func (c *controlWriter) VisitCommonProfileTag2(t tlv.CommonProfileTag2) error {
	return c.write(commonProfileTag2Form)
}
func (c *controlWriter) VisitCommonProfileTag4(t tlv.CommonProfileTag4) error {
	return c.write(commonProfileTag4Form)
}
func (c *controlWriter) VisitImplicitProfileTag2(t tlv.ImplicitProfileTag2) error {
	return c.write(implicitProfileTag2Form)
}
func (c *controlWriter) VisitImplicitProfileTag4(t tlv.ImplicitProfileTag4) error {
	return c.write(implicitProfileTag4Form)
}
func (c *controlWriter) VisitFullyQualifiedTag6(t tlv.FullyQualifiedTag6) error {
	return c.write(fullyQualifiedTag6Form)
}
func (c *controlWriter) VisitFullyQualifiedTag8(t tlv.FullyQualifiedTag8) error {
	return c.write(fullyQualifiedTag8Form)
}

func (w *payloadWriter) VisitAnonymousTag() error {
	return nil
}

func (w *payloadWriter) VisitContextSpecificTag(t tlv.ContextSpecificTag) error {
	return wire.WriteInt(w, t)
}

func (w *payloadWriter) VisitCommonProfileTag2(t tlv.CommonProfileTag2) error {
	return wire.WriteInt(w, t)
}

func (w *payloadWriter) VisitCommonProfileTag4(t tlv.CommonProfileTag4) error {
	return wire.WriteInt(w, t)
}

func (w *payloadWriter) VisitImplicitProfileTag2(t tlv.ImplicitProfileTag2) error {
	return wire.WriteInt(w, t)
}

func (w *payloadWriter) VisitImplicitProfileTag4(t tlv.ImplicitProfileTag4) error {
	return wire.WriteInt(w, t)
}

func (w *payloadWriter) VisitFullyQualifiedTag6(t tlv.FullyQualifiedTag6) error {
	return marshalFullyQualifiedTag(w, t.VendorID, t.Profile, t.Tag)
}

func (w *payloadWriter) VisitFullyQualifiedTag8(t tlv.FullyQualifiedTag8) error {
	return marshalFullyQualifiedTag(w, t.VendorID, t.Profile, t.Tag)
}

func marshalFullyQualifiedTag[T uint16 | uint32](
	w io.Writer,
	vendor, profile uint16,
	tag T,
) error {
	if err := wire.WriteInt(w, vendor); err != nil {
		return err
	}

	if err := wire.WriteInt(w, profile); err != nil {
		return err
	}

	if err := wire.WriteInt(w, tag); err != nil {
		return err
	}

	return nil
}

func unmarshalTag(r *bytes.Reader, c byte) (tlv.Tag, error) {
	switch c & tagFormMask {
	default:
		return tlv.AnonymousTag, nil
	case contextSpecificTagForm:
		return wire.ReadInt[tlv.ContextSpecificTag](r)
	case commonProfileTag2Form:
		return wire.ReadInt[tlv.CommonProfileTag2](r)
	case commonProfileTag4Form:
		return wire.ReadInt[tlv.CommonProfileTag4](r)
	case implicitProfileTag2Form:
		return wire.ReadInt[tlv.ImplicitProfileTag2](r)
	case implicitProfileTag4Form:
		return wire.ReadInt[tlv.ImplicitProfileTag4](r)
	case fullyQualifiedTag6Form:
		return unmarshalFullyQualifiedTag6(r)
	case fullyQualifiedTag8Form:
		return unmarshalFullyQualifiedTag8(r)
	}
}

func unmarshalFullyQualifiedTag6(r *bytes.Reader) (tlv.FullyQualifiedTag6, error) {
	var t tlv.FullyQualifiedTag6
	return t, unmarshalFullyQualifiedTag(
		r,
		&t.VendorID,
		&t.Profile,
		&t.Tag,
	)
}

func unmarshalFullyQualifiedTag8(r *bytes.Reader) (tlv.FullyQualifiedTag8, error) {
	var t tlv.FullyQualifiedTag8
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
	if err := wire.AssignInt(r, vendor); err != nil {
		return err
	}

	if err := wire.AssignInt(r, profile); err != nil {
		return err
	}

	if err := wire.AssignInt(r, tag); err != nil {
		return err
	}

	return nil
}

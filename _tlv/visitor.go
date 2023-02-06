package tlv

import "context"

// ElementVisitor encapsulates element-type-specific logic.
type ElementVisitor interface {
	VisitSigned(VisitContext, Signed) error
	VisitUnsigned(VisitContext, Unsigned) error
	VisitBoolean(VisitContext, Boolean) error
	VisitFloat(VisitContext, Float) error
	VisitUTF8String(VisitContext, UTF8String) error
	VisitOctetString(VisitContext, OctetString) error
	VisitStructure(VisitContext, Structure) error
	VisitArray(VisitContext, Array) error
	VisitList(VisitContext, List) error
}

// VisitContext provides information about an element being visited.
type VisitContext interface {
	context.Context

	// Tag returns the element's tag, if any.
	Tag() (Tag, bool)
}

type visitContext struct {
	context.Context
	tag Tag
}

func (c visitContext) Tag() (Tag, bool) {
	return c.tag, c.tag != nil
}

package tlv

import "context"

// ElementVisitor is an interface for visiting a tree of TLV elements.
type ElementVisitor interface {
	VisitString(VisitContext, String) error
}

// VisitContext carries TLV-specific information when visiting a tree of TLV
// elements.
type VisitContext interface {
	context.Context
}

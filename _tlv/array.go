package tlv

import "golang.org/x/exp/slices"

// Array represents a Matter TLV "array" value.
type Array struct {
	members []AnonymousElement
}

// WithMembers returns a copy of e with the given members.
func (e Array) WithMembers(members ...AnonymousElement) Array {
	return Array{
		append(
			slices.Clone(e.members),
			members...,
		),
	}
}

// AcceptVisitor calls v.VisitArray().
func (e Array) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitArray(ctx, e)
}

func (m marshaler) VisitArray(ctx VisitContext, e Array) error {
	panic("not implemented")
}

// func VisitArray(ctx VisitContext, v ElementVisitor, e Array) error {
// 	for _, m := range e.members {
// 		if err := m.AcceptVisitor(ctx, v); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

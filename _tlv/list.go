package tlv

import "golang.org/x/exp/slices"

// List represents a Matter TLV "list" value.
type List struct {
	members []Element
}

// WithMembers returns a copy of e with the given members.
func (e List) WithMembers(members ...Element) List {
	return List{
		append(
			slices.Clone(e.members),
			members...,
		),
	}
}

// AcceptVisitor calls v.VisitList().
func (e List) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitList(ctx, e)
}

func (m marshaler) VisitList(ctx VisitContext, e List) error {
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

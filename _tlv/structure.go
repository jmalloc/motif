package tlv

import "golang.org/x/exp/slices"

// Structure represents a Matter TLV "structure" value.
type Structure struct {
	members []TaggedElement
}

// WithMembers returns a copy of e with the given members.
func (e Structure) WithMembers(members ...TaggedElement) Structure {
	e.members = slices.Clone(e.members)

next:
	for _, m := range members {
		for i, x := range e.members {
			if x.tag == m.tag {
				e.members[i] = m
				continue next
			}
		}

		e.members = append(e.members, m)
	}

	return e
}

// AcceptVisitor calls v.VisitStructure().
func (e Structure) AcceptVisitor(ctx VisitContext, v ElementVisitor) error {
	return v.VisitStructure(ctx, e)
}

func (m marshaler) VisitStructure(ctx VisitContext, e Structure) error {
	panic("not implemented")
}

// func VisitStructure(ctx VisitContext, v ElementVisitor, e Structure) error {
// 	for _, m := range e.members {
// 		if err := m.AcceptVisitor(ctx, v); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

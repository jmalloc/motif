package tlv

import "golang.org/x/exp/slices"

// Struct is a TLV structure element.
type Struct struct {
	members []TaggedElement
}

// With returns a copy of s with v as a member.
func (s Struct) With(t Tag, v Value) Struct {
	e := TaggedElement{t, v}
	s.members = slices.Clone(s.members)

	for i, m := range s.members {
		if m.t == t {
			s.members[i] = e
			return s
		}
	}

	s.members = append(s.members, e)
	return s
}

// AcceptElementVisitor invokes the appropriate method on vis.
func (s Struct) AcceptElementVisitor(vis ElementVisitor) {
	vis.VisitAnonymousElement(s)
}

// AcceptValueVisitor invokes the appropriate method on vis.
func (s Struct) AcceptValueVisitor(vis ValueVisitor) {
	vis.VisitStruct(s)
}

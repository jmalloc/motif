package tlv

// Null is the null value.
var Null Value = null{}

type null struct{}

func (v null) AcceptElementVisitor(vis ElementVisitor) {
	vis.VisitAnonymousElement(v)
}

func (null) AcceptValueVisitor(vis ValueVisitor) {
	vis.VisitNull()
}

package tlv

// Null is the null value.
var Null Value = null{}

const nullType = 0b000_10100

type null struct{}

func (v null) AcceptElementVisitor(vis ElementVisitor) {
	vis.VisitAnonymousElement(v)
}

func (null) AcceptValueVisitor(vis ValueVisitor) {
	vis.VisitNull()
}

func (m *marshaler) VisitNull() {
	m.control |= nullType
}

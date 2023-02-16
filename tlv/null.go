package tlv

type (
	null struct{}
)

var (
	// Null is the TLV null value.
	Null null
)

func (null) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitNull()
}

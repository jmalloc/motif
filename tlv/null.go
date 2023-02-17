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

const (
	nullType = 0b000_10100
)

func (w *controlWriter) VisitNull() error {
	return w.write(nullType)
}

func (w *payloadWriter) VisitNull() error {
	return nil
}

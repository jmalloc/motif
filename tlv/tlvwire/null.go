package tlvwire

const (
	nullType = 0b000_10100
)

func (w *controlWriter) VisitNull() error { return w.set(nullType) }
func (w *payloadWriter) VisitNull() error { return nil }

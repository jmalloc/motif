package tlvwire

import "github.com/jmalloc/motif/tlv"

const (
	boolFalseType = 0b000_01000
	boolTrueType  = 0b000_01001
)

func (w *controlWriter) VisitBool(v tlv.Bool) error {
	if v {
		return w.write(boolTrueType)
	}

	return w.write(boolFalseType)
}

func (w *payloadWriter) VisitBool(v tlv.Bool) error {
	return nil
}

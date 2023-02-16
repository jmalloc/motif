package tlvwire

import "github.com/jmalloc/motif/tlv"

const (
	boolFalseType = 0b000_01000
	boolTrueType  = 0b000_01001
)

func (w *controlWriter) VisitBool(v tlv.Bool) error { return w.set(boolType(v)) }
func (w *payloadWriter) VisitBool(v tlv.Bool) error { return nil }

func boolType(v tlv.Bool) byte {
	if v {
		return boolTrueType
	}

	return boolFalseType
}

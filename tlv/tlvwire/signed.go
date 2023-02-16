package tlvwire

import (
	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

const (
	signed1Type = 0b000_00000
	signed2Type = 0b000_00001
	signed4Type = 0b000_00010
	signed8Type = 0b000_00011
)

func (w *controlWriter) VisitSigned1(v tlv.Signed1) error {
	return w.write(signed1Type)
}

func (w *controlWriter) VisitSigned2(v tlv.Signed2) error {
	return w.write(signed2Type)
}

func (w *controlWriter) VisitSigned4(v tlv.Signed4) error {
	return w.write(signed4Type)
}

func (w *controlWriter) VisitSigned8(v tlv.Signed8) error {
	return w.write(signed8Type)
}

func (w *payloadWriter) VisitSigned1(v tlv.Signed1) error {
	return wire.WriteInt(w, v)
}

func (w *payloadWriter) VisitSigned2(v tlv.Signed2) error {
	return wire.WriteInt(w, v)
}

func (w *payloadWriter) VisitSigned4(v tlv.Signed4) error {
	return wire.WriteInt(w, v)
}

func (w *payloadWriter) VisitSigned8(v tlv.Signed8) error {
	return wire.WriteInt(w, v)
}

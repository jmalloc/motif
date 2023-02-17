package tlv

import "github.com/jmalloc/motif/internal/wire"

type (
	// Signed1 is a signed 1 octet signed integer.
	Signed1 int8

	// Signed2 is a signed 2 octet signed integer.
	Signed2 int16

	// Signed4 is a signed 4 octet signed integer.
	Signed4 int32

	// Signed8 is a signed 8 octet signed integer.
	Signed8 int64
)

func (v Signed1) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitSigned1(v)
}

func (v Signed2) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitSigned2(v)
}

func (v Signed4) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitSigned4(v)
}

func (v Signed8) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitSigned8(v)
}

const (
	signed1Type = 0b000_00000
	signed2Type = 0b000_00001
	signed4Type = 0b000_00010
	signed8Type = 0b000_00011
)

func (w *controlWriter) VisitSigned1(v Signed1) error {
	return w.write(signed1Type)
}

func (w *controlWriter) VisitSigned2(v Signed2) error {
	return w.write(signed2Type)
}

func (w *controlWriter) VisitSigned4(v Signed4) error {
	return w.write(signed4Type)
}

func (w *controlWriter) VisitSigned8(v Signed8) error {
	return w.write(signed8Type)
}

func (w *payloadWriter) VisitSigned1(v Signed1) error {
	return wire.WriteInt(w, v)
}

func (w *payloadWriter) VisitSigned2(v Signed2) error {
	return wire.WriteInt(w, v)
}

func (w *payloadWriter) VisitSigned4(v Signed4) error {
	return wire.WriteInt(w, v)
}

func (w *payloadWriter) VisitSigned8(v Signed8) error {
	return wire.WriteInt(w, v)
}

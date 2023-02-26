package tlv

import (
	"bytes"
	"math"
)

type (
	// Single is a single-precision (4 octet) TLV floating-point value.
	Single float32

	// Double is a double-precision (8 octet) TLV floating-point value.
	Double float64
)

func (v Single) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitSingle(v)
}

func (v Double) acceptVisitor(vis ValueVisitor) error {
	return vis.VisitDouble(v)
}

const (
	singleType = 0b000_01010
	doubleType = 0b000_01011
)

func (c *controlWriter) VisitSingle(v Single) error {
	return c.write(singleType)
}

func (c *controlWriter) VisitDouble(v Double) error {
	return c.write(doubleType)
}

func (w *payloadWriter) VisitSingle(v Single) error {
	return writeInt(
		w,
		math.Float32bits(
			float32(v),
		),
	)
}

func (w *payloadWriter) VisitDouble(v Double) error {
	return writeInt(
		w,
		math.Float64bits(
			float64(v),
		),
	)
}

func unmarshalSingle(r *bytes.Reader) (Single, error) {
	n, err := readInt[uint32](r)
	if err != nil {
		return 0, err
	}

	return Single(
		math.Float32frombits(n),
	), nil
}

func unmarshalDouble(r *bytes.Reader) (Double, error) {
	n, err := readInt[uint64](r)
	if err != nil {
		return 0, err
	}

	return Double(
		math.Float64frombits(n),
	), nil
}

package tlvwire

import (
	"bytes"
	"math"

	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

const (
	singleType = 0b000_01010
	doubleType = 0b000_01011
)

func (c *controlWriter) VisitSingle(v tlv.Single) error {
	return c.write(singleType)
}

func (c *controlWriter) VisitDouble(v tlv.Double) error {
	return c.write(doubleType)
}

func (w *payloadWriter) VisitSingle(v tlv.Single) error {
	return wire.WriteInt(
		w,
		math.Float32bits(
			float32(v),
		),
	)
}

func (w *payloadWriter) VisitDouble(v tlv.Double) error {
	return wire.WriteInt(
		w,
		math.Float64bits(
			float64(v),
		),
	)
}

func unmarshalSingle(r *bytes.Reader) (tlv.Single, error) {
	n, err := wire.ReadInt[uint32](r)
	if err != nil {
		return 0, err
	}

	return tlv.Single(
		math.Float32frombits(n),
	), nil
}

func unmarshalDouble(r *bytes.Reader) (tlv.Double, error) {
	n, err := wire.ReadInt[uint64](r)
	if err != nil {
		return 0, err
	}

	return tlv.Double(
		math.Float64frombits(n),
	), nil
}

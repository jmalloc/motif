package tlvwire

import (
	"bytes"
	"io"
	"math"

	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/tlv"
)

const (
	float4Type = 0b000_01010
	float8Type = 0b000_01011
)

func (c *controlWriter) VisitFloat4(v tlv.Float4) error { return c.set(float4Type) }
func (c *controlWriter) VisitFloat8(v tlv.Float8) error { return c.set(float8Type) }
func (w *payloadWriter) VisitFloat4(v tlv.Float4) error { return marshalFloat4(w, v) }
func (w *payloadWriter) VisitFloat8(v tlv.Float8) error { return marshalFloat8(w, v) }

func marshalFloat4(w io.Writer, v tlv.Float4) error {
	return wire.WriteInt(
		w,
		math.Float32bits(
			float32(v),
		),
	)
}

func unmarshalFloat4(r *bytes.Reader) (tlv.Float4, error) {
	n, err := wire.ReadInt[uint32](r)
	if err != nil {
		return 0, err
	}

	return tlv.Float4(
		math.Float32frombits(n),
	), nil
}

func marshalFloat8(w io.Writer, v tlv.Float8) error {
	return wire.WriteInt(
		w,
		math.Float64bits(
			float64(v),
		),
	)
}

func unmarshalFloat8(r *bytes.Reader) (tlv.Float8, error) {
	n, err := wire.ReadInt[uint64](r)
	if err != nil {
		return 0, err
	}

	return tlv.Float8(
		math.Float64frombits(n),
	), nil
}

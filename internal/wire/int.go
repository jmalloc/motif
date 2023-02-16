package wire

import (
	"io"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// WriteInt writes an integer to w in little-endian order.
func WriteInt[T constraints.Integer](w io.Writer, n T) error {
	size := int(unsafe.Sizeof(n))
	data := make([]byte, size)

	for i := range data {
		shift := 8 * i
		data[i] = byte(n >> shift)
	}

	_, err := w.Write(data)
	return err
}

// ReadInt reads an integer from r in little-endian order.
func ReadInt[T constraints.Integer](r io.Reader) (T, error) {
	var v T
	return v, AssignInt(r, &v)
}

// AssignInt reads an integer from r in little-endian order and assigns it to
// *v.
func AssignInt[T constraints.Integer](r io.Reader, v *T) error {
	size := int(unsafe.Sizeof(*v))
	data := make([]byte, size)

	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}

	for i := range data {
		shift := 8 * i
		*v |= T(data[i]) << shift
	}

	return nil
}

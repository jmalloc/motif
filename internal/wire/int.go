package wire

import (
	"io"
	"unsafe"

	"github.com/jmalloc/motif/optional"
	"golang.org/x/exp/constraints"
)

// AppendInt appends an integer to data in little-endian order and returns the
// result.
func AppendInt[T constraints.Integer](data []byte, n T) []byte {
	size := int(unsafe.Sizeof(n))

	for i := 0; i < size; i++ {
		shift := 8 * i
		data = append(data, byte(n>>shift))
	}

	return data
}

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

// AssignOptionalInt reads an integer from r in little-endian order and assigns
// it to *v.
//
// If p is false, *v is set to None and no value is read from r.
func AssignOptionalInt[T constraints.Integer](
	r io.Reader,
	p bool,
	v *optional.Value[T],
) error {
	if !p {
		*v = optional.None[T]()
		return nil
	}

	x, err := ReadInt[T](r)
	if err != nil {
		return err
	}

	*v = optional.Some(x)

	return nil
}

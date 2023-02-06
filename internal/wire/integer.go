package wire

import (
	"io"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func MarshalLittleEndian[T constraints.Integer](
	w io.Writer,
	v T,
) error {
	n := int(unsafe.Sizeof(v))
	data := make([]byte, n)

	for i := range data {
		shift := (n - i - 1) * 8
		data[i] = byte(v >> shift)
	}

	_, err := w.Write(data)
	return err
}

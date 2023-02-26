package tlv

import (
	"fmt"
	"io"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// writeInt writes an integer to w in little-endian order.
func writeInt[T constraints.Integer](w io.Writer, n T) error {
	size := int(unsafe.Sizeof(n))
	data := make([]byte, size)

	for i := range data {
		shift := 8 * i
		data[i] = byte(n >> shift)
	}

	_, err := w.Write(data)
	return err
}

// readInt reads an integer from r in little-endian order.
func readInt[T constraints.Integer](r io.Reader) (T, error) {
	var v T
	return v, assignInt(r, &v)
}

// assignInt reads an integer from r in little-endian order and assigns it to
// *v.
func assignInt[T constraints.Integer](r io.Reader, v *T) error {
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

// writeString writes an (octet or UTF-8) string to w.
//
// The length is written as type L in little-endian order.
func writeString[
	L constraints.Unsigned,
	T ~string | ~[]byte,
](w io.Writer, s T) error {
	lenInt := len(s)
	lenL := L(lenInt)

	if int(lenL) != lenInt {
		return fmt.Errorf("string too long to encode length as %T", lenL)
	}

	if err := writeInt(w, lenL); err != nil {
		return err
	}

	_, err := w.Write([]byte(s))
	return err
}

// readString reads an (octet or UTF-8) string from readString.
//
// L is the type used to encode the string's length (in little-endian order).
func readString[
	L constraints.Unsigned,
	T ~string | ~[]byte,
](r io.Reader) (T, error) {
	var v T
	return v, assignString[L](r, &v)
}

// assignString reads an (octet or UTF-8) string from r and assigns it to *s.
//
// L is the type used to encode the string's length (in little-endian order).
func assignString[
	L constraints.Unsigned,
	T ~string | ~[]byte,
](r io.Reader, s *T) error {
	var data []byte

	n, err := readInt[L](r)
	if err != nil {
		return err
	}

	data = make([]byte, n)

	if _, err = io.ReadFull(r, data); err != nil {
		return err
	}

	*s = T(data)

	return nil
}

package wire

import (
	"io"

	"golang.org/x/exp/constraints"
)

// WriteString writes an (octet or UTF-8) string to w.
//
// The length is written as type L in little-endian order.
func WriteString[
	L constraints.Unsigned,
	T ~string | ~[]byte,
](w io.Writer, s T) error {
	n := len(s)
	if err := WriteInt(w, L(n)); err != nil {
		return err
	}

	_, err := w.Write([]byte(s))
	return err
}

// ReadString reads an (octet or UTF-8) string from r.
//
// L is the type used to encode the string's length (in little-endian order).
func ReadString[
	L constraints.Unsigned,
	T ~string | ~[]byte,
](r io.Reader) (T, error) {
	var v T
	return v, AssignString[L](r, &v)
}

// AssignString reads an (octet or UTF-8) string from r and assigns it to *s.
//
// L is the type used to encode the string's length (in little-endian order).
func AssignString[
	L constraints.Unsigned,
	T ~string | ~[]byte,
](r io.Reader, s *T) error {
	var data []byte

	n, err := ReadInt[L](r)
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

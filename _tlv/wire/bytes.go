package wire

// import (
// 	"io"

// 	"golang.org/x/exp/constraints"
// )

// func MarshalBytes[L constraints.Integer](
// 	w io.Writer,
// 	v []byte,
// ) error {
// 	n := L(len(v))
// 	if err := MarshalLittleEndian(w, n); err != nil {
// 		return err
// 	}

// 	_, err := w.Write(v)
// 	return err
// }

package wire

import (
	"bytes"
)

func MarshalUint8(w *bytes.Buffer, v uint8) {
	w.WriteByte(v)
}

func MarshalUint16(w *bytes.Buffer, v uint16) {
	w.Grow(2)
	w.WriteByte(byte(v))
	w.WriteByte(byte(v >> 8))
}

func MarshalUint32(w *bytes.Buffer, v uint32) {
	w.Grow(4)
	w.WriteByte(byte(v))
	w.WriteByte(byte(v >> 8))
	w.WriteByte(byte(v >> 16))
	w.WriteByte(byte(v >> 24))
}

func MarshalUint64(w *bytes.Buffer, v uint64) {
	w.Grow(8)
	w.WriteByte(byte(v))
	w.WriteByte(byte(v >> 8))
	w.WriteByte(byte(v >> 16))
	w.WriteByte(byte(v >> 24))
	w.WriteByte(byte(v >> 32))
	w.WriteByte(byte(v >> 40))
	w.WriteByte(byte(v >> 48))
	w.WriteByte(byte(v >> 56))
}

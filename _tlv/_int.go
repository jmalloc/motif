package tlv

// import (
// 	"io"
// )

// func MarshalSigned(
// 	w io.Writer,
// 	v int64,
// ) error {
// 	// if 0xffffffff00000000&v != 0 {
// 	// 	return MarshalSigned8(w, v)
// 	// }
// 	// if 0xffff0000&v != 0 {
// 	// 	return MarshalSigned4(w, v)
// 	// }
// 	// if 0xff00&v != 0 {
// 	// 	return MarshalSigned2(w, v)
// 	// }
// 	// return MarshalSigned1(w, v)
// }

// func MarshalSigned1(v int8) []byte {
// 	return []byte{
// 		byte(Signed1ElementType),
// 		byte(v),
// 	}
// }

// func MarshalSigned2(v int16) []byte {
// 	return []byte{
// 		byte(Signed2ElementType),
// 		byte(v),
// 		byte(v >> 8),
// 	}
// }

// func MarshalSigned4(v int32) []byte {
// 	return []byte{
// 		byte(Signed4ElementType),
// 		byte(v),
// 		byte(v >> 8),
// 		byte(v >> 16),
// 		byte(v >> 24),
// 	}
// }

// func MarshalSigned8(v int64) []byte {
// 	return []byte{
// 		byte(Signed8ElementType),
// 		byte(v),
// 		byte(v >> 8),
// 		byte(v >> 16),
// 		byte(v >> 24),
// 		byte(v >> 32),
// 		byte(v >> 40),
// 		byte(v >> 48),
// 		byte(v >> 56),
// 	}
// }

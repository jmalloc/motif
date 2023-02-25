package sp80056c

import (
	"crypto/hmac"
	"encoding/binary"
	"hash"
	"math"
)

// DeriveKeyHMAC derives a key using the HMAC-based key derivation function
// defined in NIST 800-56C, section 4.1 option 2.
func DeriveKeyHMAC(
	algo func() hash.Hash,
	secret, salt, info []byte,
	size int,
) []byte {
	if len(salt) == 0 {
		panic("salt must be not be empty")
	}

	h := hmac.New(algo, salt)

	// 1. If L > 0, then set reps = [L / H_outputBits] otherwise, output an
	// error indicator and exit this process without performing the remaining
	// actions (i.e., omitting steps 2 through 8).
	if size == 0 {
		panic("length must be greater than zero")
	}

	reps := size / h.Size()
	if size%h.Size() != 0 {
		reps++
	}

	// 2. If reps > (2^32 −1), then output an error indicator and exit this
	// process without performing the remaining actions (i.e., omitting steps 3
	// through 8).
	if reps > math.MaxUint32 {
		panic("size is too large")
	}

	var (
		// 3. Initialize a big-endian 4-byte unsigned integer counter as
		// 0x00000000, corresponding to a 32-bit binary representation of the
		// number zero.
		counter uint32
		buffer  = make([]byte, 4)

		// 5. Initialize Result(0) as an empty bit string (i.e., the null
		// string).
		result []byte
	)

	buffer = append(buffer, secret...)
	buffer = append(buffer, info...)

	// 4. If counter || Z || FixedInfo is more tha max_H_inputBits bits long,
	// then output an error indicator and exit this process without performing
	// any of the remaining actions (i.e., omitting steps 5 through 8).
	if len(buffer) > h.BlockSize() {
		panic("input is too large")
	}

	// 6. For i = 1 to reps, do the following:
	for counter <= uint32(reps) {
		// 6.1 Increment counter by 1.
		counter++

		binary.BigEndian.PutUint32(buffer[:], counter)

		// 6.2 Compute K(i) = H(counter || Z || FixedInfo).
		h.Reset()
		h.Write(buffer)

		// 6.3 Set Result(i) = Result(i – 1) || K(i).
		result = h.Sum(result)
	}

	// 7. Set DerivedKeyingMaterial equal to the leftmost L bits of
	// Result(reps).
	//
	// 8. Output DerivedKeyingMaterial.
	return result[:size]
}

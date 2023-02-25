package cryptov1

import (
	"crypto/sha256"

	"github.com/jmalloc/motif/internal/cryptov1/internal/sp80056c"
)

// DeriveKey derives an encryption key from an input key.
//
// It corresponds to the Crypto_KDF() function as defined in the Matter Core
// specification.
func DeriveKey(inputKey, salt, info []byte, size int) []byte {
	return sp80056c.DeriveKeyHMAC(
		sha256.New,
		inputKey,
		salt,
		info,
		size,
	)
}

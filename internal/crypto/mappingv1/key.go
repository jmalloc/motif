package mappingv1

import (
	"crypto/sha256"

	kdf "github.com/jmalloc/motif/internal/crypto/internal/nist/sp80056c"
)

// DeriveKey derives an encryption key from an input key.
//
// It corresponds to the Crypto_KDF() function as defined in the Matter Core
// specification.
func DeriveKey(inputKey, salt, info []byte, size int) []byte {
	return kdf.DeriveKeyHMAC(
		sha256.New,
		inputKey,
		salt,
		info,
		size,
	)
}

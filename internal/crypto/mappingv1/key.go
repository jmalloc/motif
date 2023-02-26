package mappingv1

import (
	"crypto/sha256"

	kdf "github.com/jmalloc/motif/internal/crypto/internal/nist/sp80056c"
)

const (
	// SymmetricKeySize is the size to use when generating symmetric keys.
	//
	// It corresponds to the CRYPTO_SYMMETRIC_KEY_LENGTH_BYTES constant as
	// defined in the Matter Core specification.
	SymmetricKeySize = 16
)

// SymmetricKey is a key used for symmetric encryption.
type SymmetricKey [SymmetricKeySize]byte

// DeriveKey derives an encryption key from an input key.
//
// It corresponds to the Crypto_KDF() function as defined in the Matter Core
// specification.
func DeriveKey(key SymmetricKey, salt, info []byte) SymmetricKey {
	return SymmetricKey(
		kdf.DeriveKeyHMAC(
			sha256.New,
			key[:],
			salt,
			info,
			SymmetricKeySize,
		),
	)
}

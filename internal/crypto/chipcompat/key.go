package chipcompat

import (
	"crypto/sha256"
	"io"

	"github.com/jmalloc/motif/internal/crypto/mappingv1"
	"golang.org/x/crypto/hkdf"
)

// DeriveKey derives an encryption key from an input key.
//
// It is a replacement for Crypto_KDF() function as defined in the Matter Core
// specification. It uses HKDF instead the NIST SP 800-56C KDF.
//
// See https://github.com/project-chip/connectedhomeip/issues/23572.
func DeriveKey(key mappingv1.SymmetricKey, salt, info []byte) mappingv1.SymmetricKey {
	r := hkdf.New(
		sha256.New,
		key[:],
		salt,
		info,
	)

	if _, err := io.ReadFull(r, key[:]); err != nil {
		panic(err)
	}

	return key
}

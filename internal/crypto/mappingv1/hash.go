package mappingv1

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HashSize is the size of the hash output in bytes.
//
// It corresponds to CRYPTO_HASH_LEN_BYTES constant as defined in the Matter
// Core specification.
const HashSize = sha256.Size

// Hash is a cryptographic hash.
type Hash [HashSize]byte

// ComputeHash returns the hash of the given message.
//
// It corresponds to the Crypto_Hash() function as defined in the Matter Core
// specification.
func ComputeHash(m []byte) Hash {
	return sha256.Sum256(m)
}

// ComputeHMAC returns the cryptographic keyed-hash message authentication code of the
// given message.
//
// It corresponds to the Crypto_HMAC() function as defined in the Matter Core
// specification.
func ComputeHMAC(key, data []byte) Hash {
	var mac [HashSize]byte

	h := hmac.New(sha256.New, key)
	h.Write(data)
	h.Sum(mac[:0])

	return mac
}

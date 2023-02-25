package mappingv1

import (
	"crypto/aes"
	"crypto/cipher"

	ccm "github.com/jmalloc/motif/internal/crypto/internal/nist/sp80038c"
)

const (
	// aeadMICSize is the size of the "message identity code", aka "message
	// authentication code" (MAC) in bytes.
	//
	// It corresponds to the CRYPTO_AEAD_MIC_LENGTH_BYTES constant as defined in
	// the Matter Core specification.
	aeadMICSize = 16

	// aeadLengthSize is the size of the "length" field in bytes.
	//
	// It is referred to as the "q" parameter in NIST 800-38C, and defined as 2
	// in the Matter Core specification.
	aeadLengthSize = 2
)

// Encrypt encrypts the given payload using the given key and appends
// the authentication tag to the end of the ciphertext.
//
// It corresponds to the Crypto_AEAD_GenerateEncrypt() function as defined in
// the Matter Core specification.
func Encrypt(key, payload, additional, nonce []byte) ([]byte, error) {
	return newAEAD(key).Seal(nil, nonce, payload, additional), nil
}

// Decrypt decrypts the given ciphertext using the given key and
// verifies the authentication tag.
//
// It corresponds to the Crypto_AEAD_DecryptVerify() function as defined in
// the Matter Core specification.
func Decrypt(key, payload, additional, nonce []byte) ([]byte, error) {
	return newAEAD(key).Open(nil, nonce, payload, additional)
}

func newAEAD(key []byte) cipher.AEAD {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	return ccm.NewCCM(cipher, aeadLengthSize, aeadMICSize)
}

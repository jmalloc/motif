package mappingv1

import (
	"crypto/aes"
	"crypto/cipher"

	ccm "github.com/jmalloc/motif/internal/crypto/internal/nist/sp80038c"
)

const (
	// NonceSize is the required size of encryption nonces in bytes.
	//
	// It corresponds to the CRYPTO_AEAD_NONCE_LENGTH_BYTES constant as defined
	// in the Matter Core specification.
	NonceSize = 13

	// MICSize is the size of the "message identity code", aka "message
	// authentication code" (MAC) in bytes.
	//
	// It corresponds to the CRYPTO_AEAD_MIC_LENGTH_BYTES constant as defined in
	// the Matter Core specification.
	MICSize = 16

	// aeadLengthSize is the size of the "length" field in bytes.
	//
	// It is referred to as the "q" parameter in NIST 800-38C, and defined as 2
	// in the Matter Core specification.
	aeadLengthSize = 2
)

type (
	// Nonce is a nonce used for AEAD encryption.
	Nonce [NonceSize]byte

	// Extract is a message identity code (Extract) of an encrypted payload.
	Extract [MICSize]byte
)

// AEADEncrypt encrypts the given payload using the given key and appends
// the authentication tag to the end of the ciphertext.
//
// It corresponds to the Crypto_AEAD_GenerateEncrypt() function as defined in
// the Matter Core specification.
func AEADEncrypt(
	key SymmetricKey,
	payload, additional []byte,
	nonce Nonce,
) ([]byte, error) {
	return newCCM(key).Seal(nil, nonce[:], payload, additional), nil
}

// AEADDecrypt decrypts the given ciphertext using the given key and
// verifies the authentication tag.
//
// It corresponds to the Crypto_AEAD_DecryptVerify() function as defined in
// the Matter Core specification.
func AEADDecrypt(
	key SymmetricKey,
	payload, additional []byte,
	nonce Nonce,
) ([]byte, error) {
	return newCCM(key).Open(nil, nonce[:], payload, additional)
}

// ExtractMIC returns the message identity code (MIC) from the given payload.
//
// It panics if the payload is too short.
func ExtractMIC(payload []byte) Extract {
	if len(payload) < MICSize {
		panic("payload is too short")
	}

	var mic Extract
	copy(mic[:], payload[len(payload)-MICSize:])

	return mic
}

// PrivacyEncrypt encrypts the given payload using the given key.
//
// It corresponds to the Crypto_Privacy_Encrypt() function as defined in the
// Matter Core specification.
func PrivacyEncrypt(
	key SymmetricKey,
	payload []byte,
	nonce Nonce,
) []byte {
	ciphertext := make([]byte, len(payload))

	ccm.NewCTR(
		newCipher(key),
		aeadLengthSize,
		nonce[:],
	).XORKeyStream(ciphertext, payload)

	return ciphertext
}

// PrivacyDecrypt decrypts the given payload using the given key.
//
// It corresponds to the Crypto_Privacy_Decrypt() function as defined in the
// Matter Core specification.
func PrivacyDecrypt(
	key SymmetricKey,
	payload []byte,
	nonce Nonce,
) []byte {
	return PrivacyEncrypt(key, payload, nonce)
}

// newCipher returns a new AES cipher using the given key.
func newCipher(key SymmetricKey) cipher.Block {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	return c
}

// newCCM returns a new CCM "mode of operation" that uses an AES block cipher
// with the given key.
func newCCM(key SymmetricKey) cipher.AEAD {
	return ccm.NewCCM(
		newCipher(key),
		aeadLengthSize,
		MICSize,
	)
}

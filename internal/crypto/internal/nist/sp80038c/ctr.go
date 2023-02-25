package sp80038c

import (
	"crypto/cipher"
	"fmt"
)

// NewCTR returns the CTR-mode key stream used to encrypt and decrypt.
func NewCTR(c cipher.Block, l int, nonce []byte) cipher.Stream {
	if l < 2 || l > 8 {
		panic(fmt.Sprintf(
			"invalid length size (%d), must be in range [2...8]",
			l,
		))
	}

	return newCTR(c, ctrIV(l, nonce))
}

// newCTR returns the CTR-mode key stream used to encrypt and decrypt.
func newCTR(c cipher.Block, iv block) cipher.Stream {
	iv[blockSize-1]++
	return cipher.NewCTR(c, iv[:])
}

// ctrIV returns the initialization vector for CTR mode.
func ctrIV(l int, nonce []byte) block {
	var iv block
	iv[0] = byte(l) - 1
	copy(iv[1:], nonce)
	return iv
}

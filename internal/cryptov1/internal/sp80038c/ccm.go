package sp80038c

import (
	"crypto/cipher"
	"crypto/subtle"
	"encoding/binary"
	"errors"
	"fmt"
)

// NewCCM returns a new CBC-MAC (CCM) mode of operation.
func NewCCM(c cipher.Block, l, m int) cipher.AEAD {
	if c.BlockSize() != ccmBlockSize {
		panic(fmt.Sprintf(
			"incompatible cipher block size (%d bytes), must be %d bytes",
			c.BlockSize(),
			ccmBlockSize,
		))
	}

	if m < 4 || m > 16 || m%2 != 0 {
		panic(fmt.Sprintf(
			"invalid mac size (%d bytes), must be a multiple of 2 in range [4...16]",
			m,
		))
	}

	if l < 2 || l > 8 {
		panic(fmt.Sprintf(
			"invalid length size (%d), must be in range [2...8]",
			l,
		))
	}

	return &ccm{c, l, m}
}

// ccmBlockSize is the block size supported by CCM. The cipher must use the same
// block size.
const ccmBlockSize = 16

// block is a CCM block.
type block [ccmBlockSize]byte

// zeroBlock is a block of all zeros.
var zeroBlock block

// ccm is an implementation of cipher.AEAD that implements CCM (CBC-MAC)
type ccm struct {
	cipher  cipher.Block
	lenSize int // size of length field, in bytes [2, 8]
	macSize int // size of authentication field, in bytes {4, 6, 8, 10, 12, 14 or 16}
}

func (c *ccm) NonceSize() int {
	return 15 - c.lenSize
}

func (c *ccm) Overhead() int {
	return c.macSize
}

func (c *ccm) Seal(data, nonce, plaintext, additional []byte) []byte {
	ciphertext, err := c.seal(nonce, plaintext, additional)
	if err != nil {
		panic(err)
	}
	return append(data, ciphertext...)
}

func (c *ccm) seal(nonce, plaintext, additional []byte) ([]byte, error) {
	if err := c.validateNonce(nonce); err != nil {
		return nil, err
	}

	if len(plaintext) > c.maxLen() {
		return nil, fmt.Errorf(
			"plaintext is too long (%d bytes), maximum is %d bytes",
			len(plaintext),
			c.maxLen(),
		)
	}

	a0, stream := c.keyStream(nonce)
	mac := c.mac(a0, nonce, plaintext, additional)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return append(ciphertext, mac...), nil
}

func (c *ccm) Open(data, nonce, ciphertext, additional []byte) ([]byte, error) {
	if err := c.validateNonce(nonce); err != nil {
		return nil, err
	}

	if len(ciphertext) < c.macSize {
		return nil, fmt.Errorf(
			"ciphertext is too short (%d bytes), minimum is %d bytes",
			len(ciphertext),
			c.macSize,
		)
	}

	lenP := len(ciphertext) - c.macSize
	if lenP > c.maxLen() {
		return nil, fmt.Errorf(
			"ciphertext is too long (%d bytes), maximum is %d bytes",
			len(ciphertext),
			c.maxLen()+c.macSize,
		)
	}

	a0, stream := c.keyStream(nonce)

	plaintext := make([]byte, lenP)
	stream.XORKeyStream(plaintext, ciphertext[:lenP])

	mac := c.mac(a0, nonce, plaintext, additional)
	if subtle.ConstantTimeCompare(mac, ciphertext[lenP:]) == 0 {
		return nil, errors.New("message authentication failed")
	}

	return plaintext, nil
}

// validateNonce returns an error if the nonce is not valid.
func (c *ccm) validateNonce(nonce []byte) error {
	if len(nonce) == c.NonceSize() {
		return nil
	}

	return fmt.Errorf(
		"nonce is wrong length (%d bytes), must be exactly %d bytes",
		len(nonce),
		c.NonceSize(),
	)
}

// keyStream returns the CTR-mode key stream used to encrypt and decrypt.
func (c *ccm) keyStream(nonce []byte) (block, cipher.Stream) {
	var a0 block
	a0[0] = byte(c.lenSize) - 1
	copy(a0[1:], nonce)

	a1 := a0
	a1[ccmBlockSize-1]++ // increment counter

	return a0, cipher.NewCTR(c.cipher, a1[:])
}

// maxLen returns the maximum allowed length of the plaintext.
func (c *ccm) maxLen() int {
	return (1 << (8 * c.lenSize)) - 1
}

// mac returns the MAC for the given nonce, plaintext and additional data.
func (c *ccm) mac(
	a0 block,
	nonce, plaintext, additional []byte,
) []byte {
	enc := cipher.NewCBCEncrypter(c.cipher, zeroBlock[:])

	var macT block
	b0 := c.macB0(nonce, plaintext, additional)
	enc.CryptBlocks(macT[:], b0[:])

	var bN []byte
	bN = macAppendAdditional(bN, additional)
	bN = pad(bN)
	bN = append(bN, plaintext...)
	bN = pad(bN)

	for len(bN) > 0 {
		enc.CryptBlocks(macT[:], bN[:ccmBlockSize])
		bN = bN[ccmBlockSize:]
	}

	var s0 block
	c.cipher.Encrypt(s0[:], a0[:])

	macU := xor(macT, s0)
	return macU[:c.macSize]
}

// macB0 returns the first block used when computing the MAC, referred to as B0
// in RFC 3610.
func (c *ccm) macB0(nonce, plaintext, additional []byte) block {
	var b0 block

	// flags ...
	b0[0] |= byte((c.lenSize - 1))
	b0[0] |= byte((c.macSize - 2) << 2)
	if len(additional) != 0 {
		b0[0] |= 0b01_000_000
	}

	// nonce ...
	copy(b0[1:], nonce)

	// plaintext length ...
	n := uint64(len(plaintext))
	for i := 0; i < c.lenSize; i++ {
		b0[ccmBlockSize-i-1] = byte(n >> (8 * i))
	}

	return b0
}

// macAppendAdditional appends the length of the "additional data" and the
// "additional data" itself to the data used to derived the MAC.
func macAppendAdditional(data, additional []byte) []byte {
	n := len(additional)

	switch {
	case n == 0:
		return data

	case n < 0xff00: // 2^16 - 2^8
		data = binary.BigEndian.AppendUint16(data, uint16(n))

	case n < 0x100000000: // 2^32
		data = append(data, 0xff, 0xfe)
		data = binary.BigEndian.AppendUint32(data, uint32(n))

	default:
		data = append(data, 0xff, 0xff)
		data = binary.BigEndian.AppendUint64(data, uint64(n))
	}

	return append(data, additional...)
}

// pad appends padding to the given data such that its length is a multiple of
// the CCM block size.
func pad(data []byte) []byte {
	n := len(data)
	m := n % ccmBlockSize

	return append(
		data,
		zeroBlock[:ccmBlockSize-m]...,
	)
}

// xor returns the result of XORing the two given blocks.
func xor(a, b block) block {
	for i := 0; i < ccmBlockSize; i++ {
		a[i] ^= b[i]
	}
	return a
}

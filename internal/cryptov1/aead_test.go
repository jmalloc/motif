package cryptov1_test

import (
	. "github.com/jmalloc/motif/internal/cryptov1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Encrypt() and Decrypt()", func() {
	var key, payload, additional, nonce, ciphertext []byte

	BeforeEach(func() {
		key = []byte{
			0x5e, 0xde, 0xd2, 0x44, 0xe5, 0x53, 0x2b, 0x3c,
			0xdc, 0x23, 0x40, 0x9d, 0xba, 0xd0, 0x52, 0xd2,
		}

		payload = []byte{0x05, 0x64, 0xee, 0x0e, 0x20, 0x7d}

		additional = []byte{
			0x00, 0xb8, 0x0b, 0x00, 0x39, 0x30, 0x00, 0x00,
		}

		nonce = []byte{
			0x00, 0x39, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
		}

		ciphertext = []byte{
			0x5a, 0x98, 0x9a, 0xe4, 0x2e, 0x8d, 0x84, 0x7f,
			0x53, 0x5c, 0x30, 0x07, 0xe6, 0x15, 0x0c, 0xd6,
			0x58, 0x67, 0xf2, 0xb8, 0x17, 0xdb,
		}
	})

	It("encrypts and decrypts using CCM (CBC-MAC)", func() {

		result, err := Encrypt(key, payload, additional, nonce)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result).To(Equal(ciphertext))

		result, err = Decrypt(key, ciphertext, additional, nonce)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result).To(Equal(payload))
	})

	It("returns an error if the key is incorrect", func() {
		key[0]++

		_, err := Decrypt(key, ciphertext, additional, nonce)
		Expect(err).To(MatchError("message authentication failed"))
	})

	It("returns an error if the nonce is incorrect", func() {
		nonce[0]++

		_, err := Decrypt(key, ciphertext, additional, nonce)
		Expect(err).To(MatchError("message authentication failed"))
	})

	It("returns an error if the additional data is incorrect", func() {
		additional[0]++

		_, err := Decrypt(key, ciphertext, additional, nonce)
		Expect(err).To(MatchError("message authentication failed"))
	})

	It("returns an error if the MAC is incorrect", func() {
		ciphertext[len(ciphertext)-1]++

		_, err := Decrypt(key, ciphertext, additional, nonce)
		Expect(err).To(MatchError("message authentication failed"))
	})

})

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
)

type (
	Crypto interface {
		EncryptAes16(secret, text []byte) ([]byte, error)
		DecryptAes16(secret, text []byte) ([]byte, error)
	}

	impl struct{}
)

func (c *impl) EncryptAes16(secret, text []byte) ([]byte, error) {
	encKeyLen := len(secret)

	if encKeyLen < 16 {
		return nil, fmt.Errorf("The key must be 16 bytes long")
	}

	encKeySized := secret

	if encKeyLen > 16 {
		encKeySized = secret[:16]
	}

	cc, err := aes.NewCipher(encKeySized)

	if err != nil {
		return nil, err
	}

	//----------- Create the IV
	// remember that GCM normally takes a 12 byte (96 bit) nounce
	nonceSize := 12
	iv, err := nonce(nonceSize)

	if err != nil {
		return nil, err
	}

	//----------- Encrypt
	ivLen := len(iv)
	enc, err := cipher.NewGCMWithNonceSize(cc, nonceSize)

	if err != nil {
		return nil, err
	}

	cipherText := enc.Seal(nil, iv, text, nil)

	//----------- Pack the message
	// create output tag
	output := make([]byte, 1+ivLen+len(cipherText))

	i := 0
	output[i] = byte(ivLen)
	i++

	copycopy(iv, 0, output, i, ivLen)
	i += ivLen

	copycopy(cipherText, 0, output, i, len(cipherText))

	return output, nil
}

func (c *impl) DecryptAes16(secret, text []byte) ([]byte, error) {
	encKeyLen := len(secret)

	if encKeyLen < 16 {
		return nil, fmt.Errorf("The key must be 16 bytes long")
	}

	encKeySized := secret

	if encKeyLen > 16 {
		encKeySized = secret[:16]
	}

	//----------- Unpack the message

	//----------- read the IV
	i := 0
	ivLen := int(text[i])
	i++

	if ivLen != 12 {
		return nil, fmt.Errorf("IV length is not correct, expected 12 but got %d", ivLen)
	}

	iv := make([]byte, ivLen)
	copycopy(text, i, iv, 0, ivLen)
	i += ivLen

	//----------- read the cipher text

	cipherTextLen := len(text) - i
	cipherText := make([]byte, cipherTextLen)

	copycopy(text, i, cipherText, 0, cipherTextLen)

	//----------- Decrypt

	cc, err := aes.NewCipher(encKeySized)

	if err != nil {
		return nil, err
	}

	dec, err := cipher.NewGCMWithNonceSize(cc, ivLen)

	if err != nil {
		return nil, err
	}

	output, err := dec.Open(nil, iv, cipherText, nil)

	if err != nil {
		return nil, err
	}

	return output, nil
}

// Create a single random initialised byte array of size.
func nonce(size int) ([]byte, error) {

	b := make([]byte, size)

	// not checking len here because rand.Read doc reads:
	//             On return, n == len(b) if and only if err == nil.
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func copycopy(src []byte, srcI int, dest []byte, destI int, copyLen int) {
	srcI2 := srcI + copyLen
	copy(dest[destI:], src[srcI:srcI2])
}

func ToString(src []byte) string {
	return hex.EncodeToString(src)
}

func New() (Crypto, error) {
	return &impl{}, nil
}

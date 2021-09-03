package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type (
	Crypto interface {
		EncryptAes(text string) ([]byte, error)
		DecryptAes(text string) ([]byte, error)
	}

	impl struct {
		secret []byte
	}
)

func (i *impl) EncryptAes(text string) ([]byte, error) {
	block, _ := aes.NewCipher(i.secret)
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return ciphertext, nil
}

func (i *impl) DecryptAes(text string) ([]byte, error) {
	data, err := hex.DecodeString(text)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(i.secret)

	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return []byte(plaintext), nil
}

func ToString(src []byte) string {
	return hex.EncodeToString(src)
}

func New(secret string) (Crypto, error) {
	if len(secret) != 32 {
		return nil, errors.New("secret must be 32 bytes long")
	}

	return &impl{[]byte(secret)}, nil
}

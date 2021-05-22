package cipher

import (
	"crypto/sha256"

	"golang.org/x/crypto/chacha20poly1305"
)

var _nonce = make([]byte, chacha20poly1305.NonceSizeX)

var _key = []byte("8pUsXuZw4z6B9EbGdKgNjQnjqVsYv2x5")

func GenerateKey(key string) {
	sha := sha256.Sum256([]byte(key))
	buff := make([]byte, 32)
	copy(sha[:32], buff[:32])
	_key = buff
}

func Encrypt(data *[]byte) {
	aead, _ := chacha20poly1305.NewX(_key)
	ciphertext := aead.Seal(nil, _nonce, *data, nil)
	data = &ciphertext
}

func Decrypt(data *[]byte) {
	aead, _ := chacha20poly1305.NewX(_key)
	plaintext, _ := aead.Open(nil, _nonce, *data, nil)
	data = &plaintext
}

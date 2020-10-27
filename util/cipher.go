package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"log"
)

func CreateHash(key string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}

func Encrypt(data *[]byte, key []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	iv := key[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	dest := make([]byte, len(*data))
	stream.XORKeyStream(dest, *data)
	data = nil
	data = &dest
}

func Decrypt(data *[]byte, key []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	iv := key[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	dest := make([]byte, len(*data))
	stream.XORKeyStream(dest, *data)
	data = nil
	data = &dest
}

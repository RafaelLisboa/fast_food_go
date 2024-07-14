package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func encryptPassword(password string) string {
	block, err := aes.NewCipher([]byte("AMZI/y  u'u+!2^a"))

	if err != nil {
		log.Fatal("Error trying create secret ", err)
	}

	b := base64.StdEncoding.EncodeToString([]byte(user.Password))

	passwordTextByteBuffer := make([]byte, aes.BlockSize+len(b))

	iv := passwordTextByteBuffer[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(passwordTextByteBuffer[aes.BlockSize:], []byte(b))

	passwordText := base64.StdEncoding.EncodeToString(passwordTextByteBuffer)


	return passwordText
}

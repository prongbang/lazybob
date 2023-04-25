package lazybob

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

type RSACrypto interface {
	Decrypt(cipherText string) (string, error)
	Encrypt(plainText string) (string, error)
}

type rsaCrypto struct {
	KeyPair *KeyPair
}

func (r *rsaCrypto) Decrypt(cipherText string) (string, error) {
	cipherByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, r.KeyPair.PrivateKey, cipherByte)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func (r *rsaCrypto) Encrypt(plainText string) (string, error) {
	plainByte := []byte(plainText)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, r.KeyPair.PublicKey, plainByte)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func NewRSACrypto(keyPair *KeyPair) RSACrypto {
	return &rsaCrypto{
		KeyPair: keyPair,
	}
}

package lazybob

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type AESCrypto interface {
	Decrypt(cipherText string, secret string) (string, error)
	Encrypt(plainText string, secret string) (string, error)
}

type aesCrypto struct {
}

func (a *aesCrypto) Decrypt(cipherText string, secret string) (string, error) {
	key := []byte(secret)
	ciphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64-encoded cipher text: %v", err)
	}

	iv := ciphertext[:12]
	ciphertext = ciphertext[12:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create AES-GCM cipher: %v", err)
	}

	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt cipher text: %v", err)
	}

	return string(plaintext), nil
}

func (a *aesCrypto) Encrypt(plainText string, secret string) (string, error) {
	key := []byte(secret)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create AES-GCM cipher: %v", err)
	}

	iv := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %v", err)
	}

	ciphertext := aesgcm.Seal(nil, iv, []byte(plainText), nil)
	ivCipherTextMacB64 := base64.StdEncoding.EncodeToString(append(iv, ciphertext...))

	return ivCipherTextMacB64, nil
}

func NewAESCrypto() AESCrypto {
	return &aesCrypto{}
}

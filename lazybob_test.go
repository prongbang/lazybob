package lazybob_test

import (
	"testing"

	"github.com/prongbang/lazybob"
)

var secret string
var rsaCrypto lazybob.RSACrypto
var aesCrypto lazybob.AESCrypto

func init() {
	secret = "Mkv__TcKlAv-0crq2SV8hQ=="
	keyPair, _ := lazybob.GetKeyPair("keypair/private-key.pem", "keypair/public-key.pem")
	rsaCrypto = lazybob.NewRSACrypto(keyPair)
	aesCrypto = lazybob.NewAESCrypto()
}

func TestRSAEncrypt(t *testing.T) {
	// Given
	plainText := "Hi alice :)"

	// When
	cipherText, err := rsaCrypto.Encrypt(plainText)

	// Then
	if cipherText == "" || err != nil {
		t.Error("Cannot encrypt", err)
	}
}

func TestRSADecrypt(t *testing.T) {
	// Given
	cipherText := "gA8gzlaDOGjI8IzUC0PXr6020PBBvl5zXnUa7my7Uhx7dHvuo/WZfYxEPJfpL/+K8MRDcz6enh7nChhR783vhY/APB/E//EoN1EsSU/+iHIlixQb0lHe39JMpwt/BYN8fTvtIx13Z3UR7P3oEIwPg4DFeI9sa80XZIfaSHMVEH/bTTmSf3pnHSv7cvczFfr5zf5q/k3KSFwjH5h2mnR7tF9uUdE3Fxy9ffmr4z4BStDi+PRfMNcbtqVxYJy+tNqg2Ml3lfT+m0YBLRdhAyrU+Avh1EXa27n5Fab+HPTSkA1MWR+YkZX6izT7sv/1Yim1oRQG46HbIy23FEmyGx1w5g=="

	// When
	plainText, err := rsaCrypto.Decrypt(cipherText)

	// Then
	if plainText != "Hi alice :)" || err != nil {
		t.Error("Cannot encrypt", err)
	}
}

func TestSecureRandomString(t *testing.T) {
	// Given
	size := 16

	// When
	secret, err := lazybob.RandomString(uint(size))

	// Then
	if secret == "" || err != nil {
		t.Error("Cannot random string:", err)
	}
}

func TestAESEncrypt(t *testing.T) {
	// Given
	plainText := "Hi bob :)"

	// When
	cipherText, err := aesCrypto.Encrypt(plainText, secret)

	// Then
	if cipherText == "" || err != nil {
		t.Error("Cannot encrypt", err)
	}
}

func TestAESDecrypt(t *testing.T) {
	// Given
	cipherText := "1sY6nwucMF9pqJfhflbZUjs6gqXLLIesHbAErNphe+Pw74VR9A=="

	// When
	plainText, err := aesCrypto.Encrypt(cipherText, secret)

	// Then
	if plainText == "" || err != nil {
		t.Error("Cannot encrypt", err)
	}
}

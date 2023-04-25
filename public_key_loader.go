package lazybob

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type PublicKeyLoader interface {
	Load(filename string) (*rsa.PublicKey, error)
}

type publicKeyLoader struct {
	FileLoader Loader
}

// Load implements PublicKeyLoader
func (p *publicKeyLoader) Load(filename string) (*rsa.PublicKey, error) {
	// Read the PEM file
	pemData, err := p.FileLoader.Load(filename)
	if err != nil {
		fmt.Println("Failed to read PEM file:", err)
		return nil, err
	}

	// Decode the PEM data
	block, _ := pem.Decode(pemData)
	if block == nil {
		fmt.Println("Failed to decode PEM data")
		return nil, err
	}

	// Parse the public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed to parse public key:", err)
		return nil, err
	}

	return publicKey.(*rsa.PublicKey), nil
}

func NewPublicKeyLoader(fileLoader Loader) PublicKeyLoader {
	return &publicKeyLoader{
		FileLoader: fileLoader,
	}
}

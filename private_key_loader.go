package lazybob

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type PrivateKeyLoader interface {
	Load(filename string) (*rsa.PrivateKey, error)
}

type privateKeyLoader struct {
	FileLoader Loader
}

func (p *privateKeyLoader) Load(filename string) (*rsa.PrivateKey, error) {
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

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed to parse private key:", err)
		return nil, err
	}

	return privateKey, nil
}

func NewPrivateKeyLoader(fileLoader Loader) PrivateKeyLoader {
	return &privateKeyLoader{
		FileLoader: fileLoader,
	}
}

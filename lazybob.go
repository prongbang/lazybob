package lazybob

import (
	"crypto/rsa"
	"fmt"
)

type KeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var loaderPrivateKey PrivateKeyLoader
var loaderPublicKey PublicKeyLoader

func init() {
	fileLoader := NewFileLoader()
	loaderPrivateKey = NewPrivateKeyLoader(fileLoader)
	loaderPublicKey = NewPublicKeyLoader(fileLoader)
}

func GetKeyPair(privateKeyPath, publicKeyPath string) (*KeyPair, error) {
	privateKey, err := loaderPrivateKey.Load(privateKeyPath)
	if err != nil {
		fmt.Println("Load private key:", err)
		return nil, err
	}
	publicKey, err := loaderPublicKey.Load(publicKeyPath)
	if err != nil {
		fmt.Println("Load public key:", err)
		return nil, err
	}
	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

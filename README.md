# lazybob

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

```
go get github.com/prongbang/lazybob
```

## Generate RSA KeyPair

- Generate an RSA private key, of size 2048

```shell
openssl genrsa -out keypair/private-key.pem 2048
```

- Extract the public key from the key pair

```shell
openssl rsa -in keypair/private-key.pem -outform PEM -pubout -out keypair/public-key.pem
```

## Get KeyPair

```go
keyPair, _ := lazybob.GetKeyPair("keypair/private-key.pem", "keypair/public-key.pem")
rsaCrypto = lazybob.NewRSACrypto(keyPair)
```

## RSA Encrypt

```go
plainText := "Hi alice :)"

cipherText, err := rsaCrypto.Encrypt(plainText)
```

## RSA Decrypt

```go
cipherText := "gA8gzlaDOGjI8IzUC0PXr6020PBBvl5zXnUa7my7Uhx7dHvuo/WZfYxEPJfpL/+K8MRDcz6enh7nChhR783vhY/APB/E//EoN1EsSU/+iHIlixQb0lHe39JMpwt/BYN8fTvtIx13Z3UR7P3oEIwPg4DFeI9sa80XZIfaSHMVEH/bTTmSf3pnHSv7cvczFfr5zf5q/k3KSFwjH5h2mnR7tF9uUdE3Fxy9ffmr4z4BStDi+PRfMNcbtqVxYJy+tNqg2Ml3lfT+m0YBLRdhAyrU+Avh1EXa27n5Fab+HPTSkA1MWR+YkZX6izT7sv/1Yim1oRQG46HbIy23FEmyGx1w5g=="

plainText, err := rsaCrypto.Decrypt(cipherText)
```

## Generate secret

```go
secret, err := lazybob.RandomString(16)
```

## AES Encrypt

```go
plainText := "Hi bob :)"

cipherText, err := aesCrypto.Encrypt(plainText, secret)
```

## AES Decrypt

```go
cipherText := "1sY6nwucMF9pqJfhflbZUjs6gqXLLIesHbAErNphe+Pw74VR9A=="

plainText, err := aesCrypto.Encrypt(cipherText, secret)
```

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	publicKeyName := "testWebsite1/config/keys/public.pem"
	privateKeyName := "testWebsite1/config/keys/private.pem"

	publicKey, err := ioutil.ReadFile(publicKeyName)
	if err != nil {
		panic(err)
	}
	privateKey, err := ioutil.ReadFile(privateKeyName)
	if err != nil {
		panic(err)
	}

	password := []byte("123456")
	cipherdata, err := encryptOAEP(publicKey, password)
	if err != nil {
		panic(err)
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	fmt.Println(ciphertext)

	cipherdata, _ = base64.StdEncoding.DecodeString(ciphertext)
	password, err = decryptOAEP(privateKey, cipherdata)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(password))
}

func encryptOAEP(publicKey, password []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err := fmt.Errorf("failed to parse certificate PEM")
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey := pub.(*rsa.PublicKey)

	h := sha256.New() // sha1.New() or md5.New()
	return rsa.EncryptOAEP(h, rand.Reader, rsaPublicKey, password, nil)
}

func decryptOAEP(privateKey, cipherdata []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err := fmt.Errorf("failed to parse certificate PEM")
		return nil, err
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) // ASN.1 PKCS#1 DER encoded form.
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	return rsa.DecryptOAEP(h, rand.Reader, priv, cipherdata, nil)
}

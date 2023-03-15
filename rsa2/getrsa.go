package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
)

func generateRSAKey2File(bits int, publicKeyName, privateKeyName string) (err error) {
	publicKeyWriter := bytes.NewBuffer([]byte{})
	privateKeyWriter := bytes.NewBuffer([]byte{})

	err = generateRSAKey(bits, publicKeyWriter, privateKeyWriter)
	if err != nil {
		return
	}

	fmt.Println(string(publicKeyWriter.Bytes()))
	fmt.Println(string(privateKeyWriter.Bytes()))

	err = ioutil.WriteFile(publicKeyName, publicKeyWriter.Bytes(), 0744)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(privateKeyName, privateKeyWriter.Bytes(), 0700)
	if err != nil {
		return
	}

	return
}

func generateRSAKey(bits int, publicKeyWriter, privateKeyWriter io.Writer) error {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509PrivateKey,
	}
	err = pem.Encode(privateKeyWriter, privateBlock)
	if err != nil {
		return err
	}

	// generate public key
	publicKey := &privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509PublicKey,
	}

	err = pem.Encode(publicKeyWriter, publicBlock)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	publicKeyName := "./public.pem"
	privateKeyName := "./private.pem"
	bits := 2048
	err := generateRSAKey2File(bits, publicKeyName, privateKeyName)
	if err != nil {
		panic(err)
	}
}

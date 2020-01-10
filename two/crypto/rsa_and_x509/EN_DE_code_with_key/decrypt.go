package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func RsaS1Decrypt(baseEncryMsg string) ([]byte, error) {
	var rsaSertKey, _ = ioutil.ReadFile("./rsa_prv.key")
	block, _ := pem.Decode(rsaSertKey)
	if block == nil {
		return nil, fmt.Errorf("private key error!")
	}

	// Parse With rsa_privatekey with x509 ParsePKCS1PrivateKey
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("s1 private key err: %v\n", err)
	}

	// base64 decode
	decryptMsg, err := base64.StdEncoding.DecodeString(baseEncryMsg)
	if err != nil {
		return nil, fmt.Errorf("base64 decode err: %v\n", err)
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, decryptMsg)
}

func main() {
	ciphertext, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Printf("error happend while read file : %s", err)
	}
	deCodeCtx, err := RsaS1Decrypt(string(ciphertext))
	if err != nil {
		fmt.Printf("error happened while decode with err : %s", err)
	}
	fmt.Println(string(deCodeCtx))
}

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

type TokenContent struct {
	User    string `json:"user"`
	Role    string `json:"role"`
	OptCode string `json:"optCode"`
}

var (
	stringFlag = flag.String("str", "", "insert some string to be encrypt")
)

func main() {
	flag.Parse()
	var contentBytes []byte
	var err error

	if *stringFlag == "" {
		tokenRes := &TokenContent{}
		tokenRes.User = "jim"
		tokenRes.Role = "admin"
		tokenRes.OptCode = "123"
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		contentBytes, err = json.Marshal(tokenRes)
		// fmt.Println(string(contentBytes))
		if err != nil {
			fmt.Printf("error happend in marshal : %s", err)
		}
	} else {
		contentBytes = []byte(*stringFlag)
	}

	encryMsg, err := RsaEncrypt(contentBytes)
	if err != nil {
		fmt.Printf("error happend in encrypt : %s", err)
	}

	fmt.Println(encryMsg)

	if err := ioutil.WriteFile("myfile.data", []byte(encryMsg), 0777); err != nil {
		fmt.Println(err)
	}

}

func RsaEncrypt(origData []byte) (string, error) {
	var rsaPubKey, _ = ioutil.ReadFile("./rsa_pub.key")
	block, _ := pem.Decode(rsaPubKey)
	if block == nil {
		return "", fmt.Errorf("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("error happend in parse key : %s", err)
	}

	pub := pubInterface.(*rsa.PublicKey)

	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	if err != nil {
		tmpStr := fmt.Sprintf("rsa encrypt err: %v", err)
		return "", errors.New(tmpStr)
	}

	baseEncryMsg := base64.StdEncoding.EncodeToString(encryptBytes)

	return baseEncryMsg, nil
}

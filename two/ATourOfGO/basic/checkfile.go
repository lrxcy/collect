package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	email    = "rax845163@gmail.com"
	password = "qwer12345678"
	url      = "https://api.aiservice.io/devdcaccount/v1/login"
)

func main() {
	crit, _ := PathExists("govcc.sh")

	if crit != true {
		fmt.Println("create file")
		f, _ := os.Create("govcc.sh")

		f.WriteString(GetAuthority(email, password, url))

	} else {
		fmt.Println("file already exists")
		inputFile, _ := os.Open("govcc.sh")
		inputRead := bufio.NewReader(inputFile)

		fmt.Println(inputRead.ReadString('\n'))

	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAuthority(email string, password string, url string) string {
	type Payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	data := Payload{
		// fill struct
		Email:    email,
		Password: password,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		// handle err
	}

	var f interface{}

	jsonResponse := f
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &jsonResponse)
	authority := jsonResponse.(map[string]interface{})["Authentication"].(string)

	defer resp.Body.Close()

	return authority
}

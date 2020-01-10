package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// 訪問的網址
	queryUrl := "http://35.220.156.89:8080"
	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		panic(err)
	}

	client := sendAgent()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func dumpResp(respBody io.Reader) {
	html, err := ioutil.ReadAll(respBody)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(os.Stdout, string(html))
}

func sendAgent() *http.Client {

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment, // 在環境變數中夾帶ProxyFromEnvironment
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial, TLSHandshakeTimeout: 10 * time.Second}

	client := &http.Client{
		Transport: transport,
	}
	return client
}

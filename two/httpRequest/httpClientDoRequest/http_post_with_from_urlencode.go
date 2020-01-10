package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpDo() {
	url := "http://www.01happy.com/demo/accept.php"
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	httpDo()
}

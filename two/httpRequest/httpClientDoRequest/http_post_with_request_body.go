package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonStruct struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

func httpDo() {
	// url := "http://www.01happy.com/demo/accept.php"
	url := "http://127.0.0.1:8000/echo/api/echo_something"
	client := &http.Client{}

	message := &JsonStruct{
		Meta: map[string]interface{}{
			"meta": "jim",
			"code": 1,
		},
		Content: JsonStruct{
			Meta: map[string]interface{}{
				"test": "test",
			},
		},
	}

	postbody, _ := json.Marshal(message)

	req, err := http.NewRequest("POST", url, bytes.NewReader(postbody))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/json")
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

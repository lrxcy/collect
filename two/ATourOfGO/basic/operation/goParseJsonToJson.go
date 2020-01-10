// refer to website
// https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar string
}



func main() {
	foo1 := new(Foo) // or &Foo{}
	getJson("http://example.com", foo1)
	println(foo1.Bar)

	// alternately:

	foo2 := Foo{}
	getJson("http://example.com", &foo2)
	println(foo2.Bar)
}

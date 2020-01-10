package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type money struct {
	Base     string  `json:"base"`
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

type info struct {
	Data money
}

func main() {
	str := `{"data":{"base":"BTC","currency":"USD","amount":4225.87}}`

	var i info

	if err := json.Unmarshal([]byte(str), &i); err != nil {
		fmt.Println("ugh: ", err)
	}

	fmt.Println("info: ", i)
	fmt.Println("currency: ", i.Data.Currency)

	// var f map[string]interface{}
	aa := &OptDrawResp{}
	if err := json.Unmarshal(a, aa); err != nil {
		log.Println(err)
	}
	fmt.Printf("Value of aa is %v\n", aa)

	aa = &OptDrawResp{}
	// a = a[6:]
	// a = bytes.TrimPrefix(a, []byte{239, 187, 191})
	a = bytes.Trim(a, "\xef\xbb\xbf")
	if err := json.Unmarshal(a, aa); err != nil {
		log.Println(err)
	}
	fmt.Printf("Value of aa is %v\n", aa)

}

var a = []byte{239, 187, 191, 239, 187, 191, 123, 34, 115, 116, 97, 116, 117, 115, 34, 58, 48, 44, 34, 109, 115, 103, 34, 58, 34, 83, 117, 99, 99, 101, 115, 115, 34, 44, 34, 108, 105, 110, 101, 34, 58, 51, 54, 53, 44, 34, 100, 97, 116, 97, 34, 58, 123, 34, 99, 111, 100, 101, 34, 58, 34, 106, 115, 99, 116, 107, 108, 49, 48, 103, 100, 34, 44, 34, 105, 115, 115, 117, 101, 78, 111, 34, 58, 34, 50, 48, 49, 57, 48, 56, 49, 56, 53, 54, 52, 34, 44, 34, 112, 114, 105, 122, 101, 78, 117, 109, 34, 58, 34, 49, 44, 52, 44, 54, 44, 49, 56, 44, 51, 44, 57, 44, 49, 52, 44, 49, 57, 34, 125, 44, 34, 95, 114, 111, 117, 116, 101, 34, 58, 34, 117, 112, 100, 97, 116, 101, 103, 97, 109, 101, 104, 105, 115, 116, 111, 114, 121, 34, 125}

// operator draw resp
type OptDrawResp struct {
	Status int                    `json: "status"`
	Data   map[string]interface{} `json: "data"`
}

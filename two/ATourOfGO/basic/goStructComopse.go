package main

import (
	"fmt"
	"net/http"
	"time"
)

type tagType struct {
	C1 *http.Client
	C2 *http.Client
}

type cStr struct {
	name   string
	Client tagType
}

func main() {
	tCstr := cStr{
		Client: tagType{
			C1: &http.Client{Timeout: time.Second * 2},
			C2: &http.Client{Timeout: time.Second * 3},
		},
	}

	fmt.Println(tCstr.Client.C1.Timeout)
	fmt.Println(tCstr.Client.C2.Timeout)

}

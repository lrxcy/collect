package main

import (
	"fmt"
	"strings"
)

const (
	longstring = "1,1,3;1,1,4;3,5,6;2,2,4;4,6,6"
)

func main() {

	a := "1,1,3"
	b := "1,1,4"
	// "1,1,3;1,1,4;3,5,6;2,2,4;4,6,6"
	fmt.Println(strings.Contains(longstring, a))
	fmt.Println(strings.Contains(longstring, b))
}

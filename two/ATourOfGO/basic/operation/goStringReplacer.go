// refer to website
// https://golang.org/pkg/strings/#NewReplacer

package main

import (
	"fmt"
	"strings"
)

func main() {
	// replace "<" as "*" and replace "<" as "*"
	r := strings.NewReplacer("<", "*", ">", "*")
	fmt.Println(r.Replace("This is <Jim>"))
}

package main

import (
	"fmt"
	"strings"
)

func match(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s[i:], ")")
		if j >= 0 {
			return s[i+1 : j-i]
		}
	}
	return ""
}

func main() {
	s := `(tag)SomeText`
	r := match(s)
	fmt.Println(r)
}

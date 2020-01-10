package main

import (
	"fmt"
	"regexp"
)

var rgx = regexp.MustCompile(`\((.*?)\)`)

func main() {
	s := `(tag)SomeText`
	rs := rgx.FindStringSubmatch(s)
	fmt.Println(rs[1])
}

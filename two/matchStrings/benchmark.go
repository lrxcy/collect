package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var rgx = regexp.MustCompile(`\((.*?)\)`)

const n = 1000000

func main() {
	var rs []string
	var r string
	s := `(tag)SomeText`
	t := time.Now()
	for i := 0; i < n; i++ {
		rs = rgx.FindStringSubmatch(s)
	}
	fmt.Println(time.Since(t))
	fmt.Println(rs[1]) // [(tag) tag]

	t = time.Now()
	for i := 0; i < n; i++ {
		r = match(s)
	}
	fmt.Println(time.Since(t))
	fmt.Println(r)

}
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

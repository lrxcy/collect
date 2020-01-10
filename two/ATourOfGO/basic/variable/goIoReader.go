package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader("Read will return these bytes")

	fmt.Println(r)
	fmt.Println("----Reader to String---")

	buf := new(bytes.Buffer) // create a temp memory to store
	buf.ReadFrom(r)          // where r can be replace as any of Reader
	s := buf.String()        // claim s as string for Reader
	fmt.Println(s)
}

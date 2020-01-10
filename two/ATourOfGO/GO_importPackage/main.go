package main

import (
	"fmt"

	"github.com/jimweng/GO_importPackage/goexample"
	"github.com/jimweng/GO_importPackage/string"
)

func main() {
	goexample.Hi()
	goexample.Hello()
	var test = string.Reverse("Hello")
	fmt.Println(test)
}

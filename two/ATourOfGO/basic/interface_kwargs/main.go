package main

import "fmt"

func testf(a string, b ...interface{}) {
	fmt.Printf("%v\n", a)
	for i, j := range b {
		fmt.Printf("%v___%v\n", i, j)
	}
}

func main() {
	a := "123"
	var b []interface{}
	b = append(b, "234", 123, false)
	testf(a, b, a, b, 1, 2, 3, 4)
}

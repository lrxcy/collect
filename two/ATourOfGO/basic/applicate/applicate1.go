// https://gobyexample.com/pointers
package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// claim a self-define struct
type TestConfig struct {
	people string
	foot   int
}

// 宣告一個只有 TestConfig的專屬function
func (c *TestConfig) Pos(a int) (string, int) {
	fmt.Println("hello")
	return c.people, c.foot + a
}

func main() {
	a := 1
	fmt.Println("initial:", a)

	zeroval(a)
	fmt.Println("zeroval:", a)

	zeroptr(&a)
	fmt.Println("zeroval:", a)

	fmt.Println("zeroptr:", &a)

	var test TestConfig
	test = TestConfig{"jim", 123}
	fmt.Println(test.Pos(1))
}

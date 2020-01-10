// https://gobyexample.com/slices
package main

import (
	"fmt"
)

func normal_slice() {
	s := []int{2, 3, 4, 5, 6}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]==%d\n", i, s[i])
	}
}

func main() {
	// nil-slices
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
	fmt.Println("----------for normal slice would be like this-----------")
	normal_slice()
	fmt.Println("----------or you can use make to build a dynamic slice-----------")
	// 'make' a slice of string of length 3
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

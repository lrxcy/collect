package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("s == ", s)
	fmt.Println("s[1:4] == ", s[1:4])

	// 省略下標代表從0開始
	fmt.Println("s[:3] == ", s[:3])

	// 省略上標代表到len(s)結束
	fmt.Println("s[4:] == ", s[4:])
}

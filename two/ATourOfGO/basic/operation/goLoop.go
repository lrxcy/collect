package main

import (
	"fmt"
)

// for loop
func normalFor(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += 1
	}
	return sum
}

// 'for' is equal to 'while(in c++)'
func rangeFor(n int) int {
	sum := 1
	for sum < n {
		sum += sum
	}
	return sum
}

// while loop

func main() {
	fmt.Println(normalFor(12))
	fmt.Println(rangeFor(12))
}

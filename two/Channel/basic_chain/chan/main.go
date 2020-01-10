package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -15, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // -11
	go sum(s[len(s)/2:], c) //17
	go sum(s, c)            // 6

	// x, y := <-c, <-c         // receive from c
	x, y, z := <-c, <-c, <-c // receive from c

	// fmt.Println(x, y, x+y)
	fmt.Println(x, y, z, x+y+z) // 6 17 -11 12 ;FILO

}

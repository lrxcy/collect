package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	c := make(chan bool)
// 	m := make(map[string]string)
// 	go func() {
// 		m["1"] = "a" // First conflicting access.
// 		c <- true
// 	}()
// 	m["2"] = "b" // Second conflicting access.
// 	<-c
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }

// - Fixed Version
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(5)
// 	for i := 0; i < 5; i++ {
// 		go func(i int) {
// 			fmt.Println(i) // Not the 'i' you are looking for.
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var wg sync.WaitGroup // without wg, i can't be iterator
	// var i *int            // cause panic
	// i := new(int) // safe way to declare an `i` address and assign value to `i`
	i := 0
	wg.Add(1)
	go func() {
		// *i = 5
		i = 5
		wg.Done()
	}()
	wg.Wait()

	// return *i
	return i
}

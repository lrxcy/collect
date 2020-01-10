// defer would delay the process function until the last function is finished

package main

import "fmt"

func main() {
	defer fmt.Println("hello")
	fmt.Println("After this sentence is finished, console would print hello")
}

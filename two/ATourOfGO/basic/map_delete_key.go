package main

import "fmt"

func main() {
	verties := make(map[string]int)
	verties["triangle"] = 2
	verties["square"] = 3
	verties["dodecagon"] = 12

	delete(verties, "square")
	fmt.Println(verties) // map[dodecagon:12 triangle:2]
}

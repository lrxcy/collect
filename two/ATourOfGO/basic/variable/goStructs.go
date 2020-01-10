package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// which is equal to
	data := Vertex{
		X: 1,
		Y: 2,
	}
	fmt.Println(data)
	// or

}

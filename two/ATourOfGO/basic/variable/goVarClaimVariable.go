package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}  // claim type as Vertex2
	v2 = Vertex2{X: 1}  // Ignore Y:0 is fine
	v3 = Vertex2{}      // Ignore both X:0 and Y:0
	p1 = &Vertex2{1, 2} //claim type *Vertex
)

func main() {
	v := Vertex{1, 2}
	// insert X value via v.X
	v.X = 4
	fmt.Println(v.X, v.Y)

	// combine with pointer
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// check var content
	fmt.Println(v1, p1, v2, v3)

}

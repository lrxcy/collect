package main

import (
	"fmt"
)

// parse node info
type NodeInfo struct {
	tag      string
	label    string
	domainId string
	name     string
	vcsa     string
}

type test1 struct {
	V string
}

type test2 struct {
	V int
}

func genNode(a interface{}) *NodeInfo {

	switch a.(type) {
	case *test1:
		a := a.(*test1)
		fmt.Println(a.V)
	default:
		a := a.(*test2)
		fmt.Println(a.V)
	}

	return nil
}

func main() {
	a := &test1{
		V: "string",
	}
	b := &test2{
		V: 123,
	}
	genNode(a)
	genNode(b)

	fmt.Print("This is example for a.(type)")
}

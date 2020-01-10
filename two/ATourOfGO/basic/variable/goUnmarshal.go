// https://golang.org/pkg/encoding/json/#example_Unmarshal
// http://lolikitty.pixnet.net/blog/post/129844768

package main

import (
	"encoding/json"
	"fmt"
)

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

// Remark with nil
// https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang

// In Go, nil is the zero value for pointers, interfaces, maps, slices, channels and function types,
// representing an uninitialized value.

// nil doesn't mean some "undefined" state,
// it's a proper value in itself.
// An object in Go is nil simply if and only if it's value is nil,
// which it can only be if it's of one of the aforementioned types.

// An error is an interface, so nil is a valid value for one, unlike for a string.

// *** For obvious reasons a nil error represents no error.

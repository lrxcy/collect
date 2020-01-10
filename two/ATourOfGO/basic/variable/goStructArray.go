// Refer website
// https://stackoverflow.com/questions/26159416/init-array-of-structs-in-go
package main

import (
	"fmt"
)

// claim a opt struct
type opt struct {
	shortnm      byte
	longnm, help string
	needArg      bool
}

// declare basenameOpts store []opt
var basenameOpts []opt

// declare function initdd to set up initial value of basenameOpts
func initdd() {
	basenameOpts = []opt{
		opt{
			shortnm: 'a',
			longnm:  "multiple",
			needArg: false,
			help:    "Usage for a",
		},
		opt{
			shortnm: 'b',
			longnm:  "b-option",
			needArg: false,
			help:    "Usage for b",
		},
	}
}

func main() {
	initdd()
	fmt.Println(basenameOpts)
}

package main

import (
	"fmt"
	"time"
)

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of decimal numbers,
// each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m".

// *** Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

// claim a struct with datatype time.Duration called Duration

func main() {
	test, err := time.ParseDuration("3000ms")

	fmt.Println(test)
	fmt.Println(err)
}

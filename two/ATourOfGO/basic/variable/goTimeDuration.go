// refer to https://golang.org/pkg/time/

package main

import (
	"fmt"
	"time"
)

// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 years.
type Duration int64

func main() {

	second := time.Second
	fmt.Println(second)
	fmt.Println(int64(second / time.Millisecond)) // prints 1000

	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second) // prints 10s

}

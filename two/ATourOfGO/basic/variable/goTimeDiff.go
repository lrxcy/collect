package main

import "fmt"
import "time"

func main() {
	a, err := time.Parse("2006-01-02 15:04 MST", "2014-05-03 20:57 UTC")
	if err != nil {
		// ...
		return
	}

	delta := time.Now().Sub(a)
	fmt.Println(delta.Hours())
}

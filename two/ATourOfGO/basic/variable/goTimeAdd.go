package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	now2 := now.Add(1)
	fmt.Println(now2)
}

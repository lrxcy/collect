package main

import (
	"fmt"
	"time"
)

var (
	test_Second = time.Second * 60
	test_Minute = time.Minute * 60
)

func main() {
	fmt.Printf("The value of time.Second * 60 would be %s\n", test_Second)
	fmt.Printf("The value of time.Minute * 60 would be %s\n", test_Minute)

	if test_Second == time.Minute {
		fmt.Println("test_Second woudl equl to 1 minute")
	} else {
		fmt.Errorf("test_Second would not equal to 1 minute %s", test_Second)
	}
}

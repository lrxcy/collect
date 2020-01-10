package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var now time.Time = time.Now().UTC()
	fmt.Println("now is a type of: ", reflect.TypeOf(now))
	var name string = "Carl Johannes"
	fmt.Println("name is a type of: ", reflect.TypeOf(name))
	var age int = 5
	fmt.Println("age is a type of: ", reflect.TypeOf(age))
}

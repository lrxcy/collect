// refer to webstie
// https://tour.golang.org/moretypes/15

package main

import "fmt"

var testArray []string
var test2Array = []string{"test", "test2"}

func main() {

	var test3Array []string
	test3Array = []string{"test3"}

	testArray = append(testArray, test2Array...)
	testArray = append(testArray, test3Array...)

	fmt.Println(test2Array[0], test2Array[1])
	fmt.Println(testArray)
}

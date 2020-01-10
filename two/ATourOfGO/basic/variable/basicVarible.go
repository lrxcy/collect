package main

// can import package from ()
import (
	"fmt"
	"math"
	"math/rand"
)

// self-defined function
func add(x int, y int) int {
	return x + y
}

// use (int ,int) to note the return type and what value to return
func swap(x, y int) (int, int) {
	return y, x
}

// specifiic x and y to return in order (ps : force transfer type with type(x))
func split(sum int) (x, y int) {
	x = int(sum / 10)
	y = sum - x*10
	return x, y
}

// claim variable
var noInit_int int
var noInit_bool bool
var noInit_string string
var withInit_int = 49
var withInit_bool = true
var withInit_string = "string"

// define const ps: syntax error with `const birthday:=49`
const birthday = 49

// claim a math.Sqrt function for
func test(x1, y1 float64) float64 {
	a := x1 - y1
	return math.Sqrt(a * a)
}

// global variable with string ``
var sampleConfig = `
# global can be defined like this
servers = ["localhost"]`

// all the function would be printed here
func main() {
	fmt.Println(sampleConfig)
	fmt.Println(test(1, 1))
	fmt.Println("Self-define add fucntion ", add(42, 43))
	fmt.Println("Use math package obj Pi with value ", math.Pi)
	fmt.Println("Use variables with and without initial values:", noInit_int, noInit_bool, noInit_string, withInit_int, withInit_bool, withInit_string)
	fmt.Print("the birthday is ")
	fmt.Println(split(birthday))

	// fmt.Println(math/rand.Intn(112)) can't work
	fmt.Println("Import math/rand directory with value ", rand.Intn(112))

	// note that with function can't add more infor
	// Also can claim a,b :=swap(12,34) ,then fmt.Println(a,b) => cames the same result
	fmt.Println(swap(12, 34))
	fmt.Println(split(17))

	// also can claim variable in main func
	var i, j int = 6666, 6666
	fmt.Println(i, j)
}

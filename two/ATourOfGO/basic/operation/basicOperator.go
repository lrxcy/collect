package main

import (
	"fmt"
)

// create a fun loop function with iterator
func forLoop(iterator int) int{
	sum :=1
	for ;sum < iterator; {
		sum += sum
	}
	return sum
}

//	create a while loop fucntion with if-else and loop_criterion
func whileLoop(loop_criterion bool) int{
	sum := 1
	for loop_criterion{
		sum += 1
		if (sum > 100) {
			loop_criterion = false
		}
	}
	return sum
}


func main(){
	// basic sum function
	sum:= 0
	for i:=0; i< 10; i++{
		sum+=1
	}
	fmt.Println(sum)

	// call func
	fmt.Println(forLoop(3))
	fmt.Println(whileLoop(true))
}
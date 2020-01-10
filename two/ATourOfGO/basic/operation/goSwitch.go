package main

import (
	"fmt"
	"time"
)

var numbers = [6]int{1, 2, 3, 4, 5}

// no return value funciton - switch with particular name
func normalSwi() {
	for name, _ := range numbers {
		fmt.Println("numbers:", name)

		switch name {
		case 1:
			fmt.Println("name:", "John")
		case 2:
			fmt.Println("name:", "Jim")
		default:
			fmt.Println("name:", "Joker")
		}
	}
}

// add ordering in switch cases
func abnormalSwi() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow~")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away...")
	}
}

// switch without any criterion
func nocriterSwi() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func main() {
	normalSwi()
	abnormalSwi()
	nocriterSwi()
}

package main

import (
	"fmt"
	"math"
)

// function to check whether the sqrt is correct
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// another way to use if criterion
func pow(x, n, lim float64) float64 {
	// set v=x^n and if v is smaller than lim than return v,else return lim
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if and else
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 7),
		pow(3, 3, 28),
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

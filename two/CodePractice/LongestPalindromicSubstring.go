package main

import (
	"log"
)

func longestPalindrome(s string) string {
	var mapS = make(map[string]int)
	// var chooseMap  []string{}
	var maxCount = 0
	for _, j := range s {
		if string(j) != "" {
			mapS[string(j)]++
		}
	}

	// calculate max count
	for _, j := range mapS {
		if maxCount < j {
			maxCount = j
		}
	}

	for i, j := range mapS {
		if j == maxCount {
			return i
		}
	}
	return "conflict"
}

func main() {
	dummyString := "aaabbb"
	log.Println(longestPalindrome(dummyString))
}

package main

import (
	"fmt"
	"log"
	"strings"
)

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if ok := strings.Contains(haystack, needle); ok {
		haystackl := len(haystack)
		window := len(needle)
		// log.Println(haystackl)
		// log.Println(window)

		for i, _ := range haystack {
			if i+window <= haystackl {
				log.Println("value of i:", i)
				tmp := fmt.Sprintf("%v", haystack[i:i+window])
				// log.Println(tmp)
				// log.Println(needle)
				if needle == tmp {
					log.Println("enterHere")
					return i
				}
			}
		}
	}

	return -1
}

func main() {
	haystack := "mississippi"

	needle := "pi"
	log.Println(strStr(haystack, needle))
}

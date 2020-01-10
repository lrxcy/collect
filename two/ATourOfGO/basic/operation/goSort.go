// Refer to https://golang.org/pkg/sort/

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 3, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Reverse Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}

package main

import (
	"fmt"
	"sort"
)

// 為了要客製一個可以排序的字串陣列，我們預先設定一個陣列
type byLength []string

/*
	We implement sort.Interface - Len, Less, and Swap
	on our type so we can use the sort package’s generic Sort function.
	Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic.
	In our case we want to sort in order of increasing string length, so we use len(s[i]) and len(s[j]) here.
*/
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

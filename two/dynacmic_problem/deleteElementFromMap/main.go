package main

import "fmt"

var (
	originMap map[string]int
)

func main() {
	originMap = make(map[string]int, 5)
	originMap["test"] = 1

	originMap = *deleteMapKey(originMap, "test")
	fmt.Println(originMap)
}

func copyMap(m map[string]int) *map[string]int {
	m2 := make(map[string]int, 0)
	for i, j := range m {
		m2[i] = j
	}
	m = nil
	return &m2
}

func deleteMapKey(m map[string]int, deleteKey string) *map[string]int {
	delete(m, deleteKey)
	m2 := copyMap(m)

	return m2
}

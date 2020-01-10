package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Map 映射，將data的index以及element抽出
func Map(data []int, mapper func(int) int) []int {
	results := make([]int, len(data))
	for i, ele := range data {
		// 利用golang的一級函式特性，將results[i] 儲存為 mapper(ele)
		results[i] = mapper(ele)
	}
	// 回傳的results可以預見為 []int{mapper(ele_1), mapper(ele_2), ... }
	return results
}

func TestMap(t *testing.T) {
	// Map 傳遞 1. int陣列 2. 函數，lambda(x) : x+1
	results := Map([]int{1, 2, 3}, func(x int) int { return x + 1 })
	assert.Equal(t, []int{2, 3, 4}, results)
}

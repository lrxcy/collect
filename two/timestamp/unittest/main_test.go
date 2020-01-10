package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// WhatTimeIsIt is un testable: 因為控制時間的變數放在裡面，每次進入程序都會重新時間函數的生成，故沒辦法進行測試
func WhatTimeIsIt() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}

// WhatTimeIsIt2 is testable: 只做打印時間的函數，因為time.Time是從外部傳入的，所以可以掌握送進去的時間參數; 換言之，WhatTimeIsIt2只做打印時間，而非生成+打印時間
func WhatTimeIsIt2(clock time.Time) string {
	return fmt.Sprintf("%v", clock.Unix())
}

func TestWhatTimeIsIt(t *testing.T) {
	testT := time.Now()
	testTString := fmt.Sprintf("%v", testT.Unix())
	time.Sleep(1 * time.Second)

	// function string will variant with time
	assert.NotEqual(t, testTString, WhatTimeIsIt())

	// function would base on input time.Time
	assert.Equal(t, testTString, WhatTimeIsIt2(testT))
}

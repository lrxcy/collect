// gonna to wirte some test code
// https://www.youtube.com/watch?v=mb5En4BizFM
// Run$: go test -v -cover=true gotest_test.go gotest.go
package main

import "testing"

func TestSuccessStringInSlice(t *testing.T) {
	if ok := StringInSlice("a", []string{"a", "b"}); ok {
		t.Log("pass")
	} else {
		t.Error("faild")
	}
}

func TestFaildStringInSlice(t *testing.T) {
	if ok := StringInSlice("c", []string{"a", "b"}); ok {
		t.Error("faild")
	} else {
		t.Log("pass")
	}
}

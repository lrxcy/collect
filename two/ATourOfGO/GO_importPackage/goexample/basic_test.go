package goexample

import "testing"

func TestSomething(t *testing.T) {
	t.Skip()
}

func TestAdd(t *testing.T) {
	// t.Fail()
	// t.Error("Test for error log")
	if Add(1, 2) == 3 {
		t.Log("Add is correct!")
	} else {
		t.Error("Add is wrong")
	}

}

// Run$: go test -v -cover=true goTable_test.go gotest.go
// Any test would be added after TODO: add test cases
package main

import "testing"

func TestStringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: add test cases.
	}
	for _, tt := range tests {
		if got := StringInSlice(tt.args.a, tt.args.list); got != tt.want {
			t.Errorf("%q. StringInSlice() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

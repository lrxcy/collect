// go test -c go_test.go
// ./main.test -test.bench=. -test.cpuprofile=cpu-profile.prof
package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random()
	}
}

func random() int {
	return rand.Intn(100)
}

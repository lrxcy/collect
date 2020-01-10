package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 3; i++ {
		fmt.Println(randomString(10))
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// a better way rather than use rand.Seed
func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// refer:
// - https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
// - https://yourbasic.org/golang/crypto-rand-int/
// - https://goruncode.com/generate-a-random-int-in-go/

package main

import (
	cryRand "crypto/rand"

	"fmt"
	"math/big"
	"strings"
)

func main() {
	s := ""
	l := "132"
	testStr := "2,3,4;1,2,3"

	for i := 0; i < 3; i++ {
		s = s + ";" + l

	}

	fmt.Println(s)

	splitTestStr := strings.Split(testStr, ";")
	for i, j := range splitTestStr {
		fmt.Printf("---%v___%v---\n", i, j)
	}

	fmt.Println(len(splitTestStr))
	fmt.Println(RandNum(0, int64(len(splitTestStr))))
}

func RandNum(startNum, maxNum int64) int64 {
	maxBigInt := big.NewInt(maxNum)
	tmp, _ := cryRand.Int(cryRand.Reader, maxBigInt)
	tmpInt := tmp.Int64()

	if tmpInt >= startNum {
		return tmpInt
	}

	return RandNum(startNum, maxNum)
}

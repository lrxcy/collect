package main

import (
	cryRand "crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func retriveRandPrizeFromMsg(s string) string {
	prizeArr := strings.Split(s, ";")
	return prizeArr[RandNum(0, int64(len(prizeArr)))]

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

func main() {
	s := "2,3,4;1,2,3;3,5,6"
	fmt.Println(retriveRandPrizeFromMsg(s))
}

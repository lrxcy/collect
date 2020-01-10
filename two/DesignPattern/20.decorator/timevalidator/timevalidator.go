package timevalidator

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	IType1 = 1
	IType2 = 2

	LType1 = 1
	LType2 = 2

	STypeInc = 1
	STypeDec = 2
)

type Component interface {
	checkformat() bool
	getIType() int
	getLType() int
	getSType() int
	getIssue() int
}

type Validator struct {
	Component
}

func WrapValidator(c Component) Component {
	return &Validator{
		Component: c,
	}
}

func (w *Validator) checkformat() bool {

	// 判斷Insert型別是否符合
	switch w.Component.getIType() {
	case IType1:
		log.Println("The insert type is 1")
	case IType2:
		log.Println("The insert type is 2")
	default:
		log.Println("Invlid insert type")
	}

	// 判斷load型別是否符合
	switch w.Component.getLType() {
	case LType1:
		log.Println("The load type is 1")
	case LType2:
		log.Println("The load type is 2")
	default:
		log.Println("Invlid load type")
	}

	// 判斷Sort型別是否符合
	switch w.Component.getSType() {
	case STypeInc:
		log.Println("The sort type is increasing")
	case STypeDec:
		log.Println("The sort type is decreasing")
	default:
		log.Println("Invalid sort type")
	}

	log.Println(w.Component.getIssue())

	return w.Component.checkformat()
}

func frencyUnix(timeUStr string) (int64, int, error) {
	timelen := len(timeUStr)
	unit := timeUStr[timelen-1 : timelen]

	numStr := timeUStr[:timelen-1]

	var numUnix int64

	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("Time num %v to int64 is err : %v", numStr, err))
	}

	switch unit {
	case "s":
		numUnix = num
	case "m":
		numUnix = int64(num * 60)
	case "h":
		numUnix = int64(num * 60 * 60)
	case "d":
		numUnix = int64(num * 60 * 60 * 24)
	default:
		return 0, 0, errors.New(fmt.Sprintf("Time str uni %v is err", unit))
	}

	return numUnix, int(num), nil
}

/*
	如何去動代計算一個基於某值的迭代數字
	1. 需要知道基於值 bValue
	2. 需要知道基於日期 bDate
	3. 需要知道迭代頻率 bFrenc
	4. 需要知道一天迭代最大次數上限 bMaxDay
	公式:
		bValue + [(NowTime - bDate) % ( aDay / bFrenc )] * bMaxDay
*/

func baseValueIssueNo(frenc string, baseT string, inum int, baseValue int) (string, error) {

	fUnix, _, err := frencyUnix(frenc)
	if err != nil {
		return "", errors.New("error happened while validating issueNo")
	}

	baseTime, err := time.Parse("2006-01-02 15:04:05", baseT)
	if err != nil {
		return "", errors.New("error happened while parse baseTime")
	}

	nowTimeUnix := time.Now().Unix()

	baseUnix := baseTime.Unix() - 8*3600

	interval := nowTimeUnix - baseUnix
	IDay := interval / (3600 * 24)
	residueUnix := interval % (3600 * 24)

	issueNo := IDay*int64(inum) + int64(baseValue)
	issueNo += residueUnix / fUnix

	issueNoStr := strconv.FormatInt(issueNo, 10)

	return issueNoStr, nil
}

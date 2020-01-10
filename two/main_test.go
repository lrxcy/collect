package main

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/fatih/structs"
)

func TestRetriveString(t *testing.T) {
	str := "1234512345678"
	assert.Equal(t, retriveString(str), "")
}

type StructA struct {
	Str string
	Int int
}

func TestConvertStructToMap(t *testing.T) {
	sa := &StructA{
		Str: "str",
		Int: 111,
	}

	msa := structs.Map(sa)
	assert.Equal(t, map[string]interface{}{"Str": "str", "Int": 111}, msa)
}

func TestDate(t *testing.T) {
	nowTime := time.Now()
	// nowTimeDateStr := nowTime.Format(Date_Whippletree) // 如: 2018-01-02
	// nowTimeUnix := nowTime.Unix()
	// year := nowTime.Year()
	day := nowTime.YearDay()

	assert.Equal(t, 0, day)
}

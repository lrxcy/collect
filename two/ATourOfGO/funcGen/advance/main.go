package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 定義一個泛型函數，讓後面生成'函數"時可以承接
type formaterFunc func(formater string) (string, error)

// 定義對應的interface
type formaterGenator interface {
	getfunc() formaterFunc
}

// 定義依據生成所需的資料結構
type Validator struct {
	choice int
}

// 儲存函數的基礎資料結構
type formatStruct struct {
	ff formaterFunc
}

// 對於某些可能有動態函數的情境，額外定義一個執行函數的方法。讓原本執行的函數做更新
func (v *Validator) execFunc(formater string) (string, error) {
	a := v.getfunc()
	return a(formater)
}

// 定義出能返回"函數"的函數
func (v *Validator) getfunc() formaterFunc {
	switch v.choice {
	case 1:
		return Formater1
	case 2:
		return Formater2
	default:
		return nil

	}
	return nil
}

func Formater1(name string) (string, error) {
	return "foramter1 : " + name, nil
}

func Formater2(name string) (string, error) {
	return "foramter2 : " + name, nil
}

func main() {
	// 讀取檔案
	demos, err := ReadJson()
	if err != nil {
		panic(err)
	}

	// 宣告一個儲存初始化函數用的陣列
	formaterArray := []*formatStruct{}

	// 宣告一個用於"建造者"模式的儲存陣列
	validatorArray := []*Validator{}

	// 開始針對讀取的設定檔進行初始設定
	for _, j := range demos.ConfArray {
		/*
			如果追求方便性，可以直接將Validator與讀入的檔案宣告為同一資料結構
			此處考慮到讀入的檔案可能不同，所以額外宣告一個資料結構來承接
		*/
		tmp := &Validator{}
		tmp.choice = j.CrawlerType
		validatorArray = append(validatorArray, tmp)

		formater := formatStruct{}
		formater.ff = tmp.getfunc()
		formaterArray = append(formaterArray, &formater)
	}

	// 在需要時在提領設定檔案先前預儲存的函數
	for i, j := range formaterArray {
		fmt.Printf("%v__%v\n", i, j)
		fmt.Println(j.ff("123"))
	}

	// 動態函數，適用於建造者模式
	for i, j := range validatorArray {
		respString, err := j.execFunc("123")
		fmt.Printf("The %d result of validator func result is %v___%v\n", i, respString, err)
		if j.choice == 2 {
			fmt.Println("Change the origin choice!")
			j.choice = 1
			respString, err = j.execFunc("123")
			fmt.Printf("The %d result of validator func result is %v___%v\n", i, respString, err)
		}
	}

}

type demoStruct struct {
	SourceID  int  `json:"sourceID"`
	IsProxyIP bool `json:"isProxyIP"`
	ConfArray []struct {
		CrawlerCode string `json:"crawlerCode"`
		TimePerid   string `json:"timePerid"`
		CrawlerType int    `json:"crawlerType"`
		Type        int    `json:"Type"`
		URL         string `json:"url"`
	} `json:"ConfArray"`
}

func ReadJson() (*demoStruct, error) {
	confBytes, err := ioutil.ReadFile("./test.json")
	if err != nil {
		return nil, err
	}

	jsonStruct := &demoStruct{}
	if err = json.Unmarshal(confBytes, jsonStruct); err != nil {
		return nil, err
	}

	return jsonStruct, nil
}

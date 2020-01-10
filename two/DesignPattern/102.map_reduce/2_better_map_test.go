package main

import (
	"io"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// StringProudcer為假定要帶入的資料結構，具有指標以及字符陣列
type StringProducer struct {
	index int
	data  []string
}

// OutputConsumer ...
type OutputConsumer struct{}

func TestBetterMap(t *testing.T) {
	// BetterMap: 輸入參數 1.資料結構`StringProducer{}` 2. 資料結構`OutputConsumer{}` 3. func(str string) (int64, error)`
	int64arr, err := BetterMap(&StringProducer{data: []string{"1", "10", "11"}}, OutputConsumer{}, func(str string) (int64, error) {
		// 此處的lambda將字符串以二進制形式轉為整數，strconv.ParseInt 返回 (int64, error)
		return strconv.ParseInt(str, 2, 64)
	})
	assert.Nil(t, err)
	assert.Equal(t, []int64{1, 2, 3}, int64arr)
}

// // BetterMap 為預計的映射map reduce 函數
// func BetterMap(sp *StringProducer, oc OutputConsumer, f func(string) (int64, error)) ([]int64, error) {
// 	return nil, nil
// }

/*
	對 `StringProducer` 以及 `OutputConsumer` 做一層抽象
*/

func BetterMap(p producer, c consumer, m mapper) ([]int64, error) {
	results := make([]int64, 0) // 初始化一個回傳的陣列
	for {
		next, err := p.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err // 生產者本身遇到錯誤，返回err，終止Map
			}
		}

		datum, err := m(next)
		if err != nil {
			return nil, err // m(mapper) 如果出現錯誤，返回err，終止Map
		}
		c.Send(datum)
		results = append(results, datum) // 將回傳值append到results裡面
	}
	return results, nil
}

/*
	producer / StringProducer 代表一個數據生產者，會透過Next方以有序的方式迭代到下一個元素
	當返回io.EOF時表示窮盡了序列。同時也考慮到其他proudcer在迭代過程中可能發生的錯誤
*/
type producer interface {
	Next() (string, error)
}

// Next : 迭代 prouducer 的 index 來傳送下一個數字
func (ip *StringProducer) Next() (string, error) {
	/*
		拿取 index 這個計數器，來反饋現在已經取到哪一個數字了
		將取到的數字返回。當取過該次數字以後，index+1 在進行下次的迭代
	*/
	if ip.index < len(ip.data) {
		defer func() { ip.index++ }()
		return ip.data[ip.index], nil
	}
	// return "", nil ---> 當沒有數據可以傳輸時，回傳 ""(空字串), io.EOF(讀取已窮盡)
	return "", io.EOF
}

/*
	consumer / OutputConsumer 代表消費者，會透過Send(int64)方法，讀取新的數據
*/
type consumer interface {
	Send(int64)
}

// Send : 單純做打印功能...把返回的數字做打印
func (c OutputConsumer) Send(ele int64) {
	log.Println(ele)
}

// mapper為一個單純的映射函數，規定限制輸入參數為 string，產出參數為 (int64, error)
type mapper func(string) (int64, error)

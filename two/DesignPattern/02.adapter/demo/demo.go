package main

import "log"

// Target 代表適配的目標接口 : 手機的轉接器接口
type Target interface {
	Request() string
}

// Adaptee 適配者 : 插座
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 是被適配的工廠函數 : 生成一個轉接器到插座的介面
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

// AdapteeImpl 是被適配的目標類別 : 定義一個轉接器的介面
type AdapteeImpl struct{}

// SpecificRequest 是目標類的一個方法 : 定義轉接器使用插座的方法
func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdapter 是Adapter的工廠函數 : 生成一個轉接器
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter是轉換Adaptee為Target接口的配適器 : 定義轉接器
type adapter struct {
	Adaptee
}

// Request實現Target接口 : 提供手機一個充電的方法
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	log.Println(res)
}

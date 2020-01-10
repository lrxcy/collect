package main

import (
	"fmt"
	"os"
)

// 宣告一個變數 outputWriter 讓後面的 Fprintf 在打印的時候可以依照指定輸出
// var outputWriter io.Writer = os.Stdout // modified during testing
var outputWriter = os.Stdout

// WildStallion 定義一個 SetMediator 的接口
// WildStallion describes an interface for a Wild Stallion band member.
type WildStallion interface {
	SetMediator(mediator Mediator)
}

// Bill 為一個要被 Mediator 實踐的資料結構
// Bill describes Bill S. Preston, Esquire.
type Bill struct {
	mediator Mediator
}

// 定義 Bill 的 SetMediator 方便後續讓其被 WildStallion 做承接
// SetMediator sets the mediator.
func (b *Bill) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

// 額外對 Bill 定義一個 Respond 的方法，來做測試
// Respond responds.
func (b *Bill) Respond() {
	fmt.Fprintf(outputWriter, "Bill: What?\n")
	b.mediator.Communicate("Bill")
}

// Ted 為一個要被 Mediator 實踐的資料結構
// Ted describes Ted "Theodore" Logan.
type Ted struct {
	mediator Mediator
}

// 定義 Ted 的 SetMediator 方便後續讓其被 WildStallion 做承接
// SetMediator sets the mediator.
func (t *Ted) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

// 額外對 Ted 定義一個 Talk 的方法，來做測試
// Talk talks through mediator.
func (t *Ted) Talk() {
	fmt.Fprintf(outputWriter, "Ted: Bill?\n")
	t.mediator.Communicate("Ted")
}

// 額外對 Ted 定義一個 Respond 的方法，來做測試
// Respond responds.
func (t *Ted) Respond() {
	fmt.Fprintf(outputWriter, "Ted: Strange things are afoot at the Circle K.\n")
}

// 定義一個 Mediator 接口，提供函數(string)
// Mediator describes the interface for communicating between Wild Stallion band members.
type Mediator interface {
	Communicate(who string)
}

// ConcreateMediator 為一個具有 Bill 以及 Ted 的實例資料結構
// ConcreateMediator describes a mediator between Bill and Ted.
type ConcreateMediator struct {
	Bill
	Ted
}

/*
	NewMediator 創建一個新的 Mediator;
	備註：因為這邊 Bill 以及 Ted 都需要等 mediator 先實例後才能進行方法調度，所以無法直接做 reutnr &Container{obj: ob1}的方式宣告返回
	此外，NewMediator() 可以利用上面的接口做承接，只是在使用時需要額外做型別轉換
*/
// NewMediator creates a new ConcreateMediator.
// func NewMediator() Mediator {
func NewMediator() *ConcreateMediator {
	mediator := &ConcreateMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

// Communicate 實作 ConcreateMediator 的接口方法
// Communicate communicates between Bill and Ted.
func (m *ConcreateMediator) Communicate(who string) {
	if who == "Ted" {
		m.Bill.Respond()
	} else if who == "Bill" {
		m.Ted.Respond()
	}
}

func main() {

	bufferOutputWriter := outputWriter
	// outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	mediator := NewMediator()
	// mediator.(*ConcreateMediator).Ted.Talk()
	mediator.Ted.Talk()

	/*
		Ted.Talk() 先行觸發，之後 mediator在觸發 Bill.Respond()，Bill.Respond() 接著又在觸發Ted.Respond()

		Ted: Bill?
		Bill: What?
		Ted: Strange things are afoot at the Circle K.
	*/
}

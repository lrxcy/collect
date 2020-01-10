package main

import "fmt"

// Command 定義這個interface，只要是在這個interface下所有的物件都具有Execute()這個方法
type Command interface {
	Execute()
}

// 用MotherBoard 物件來當作後續用Command操作的物件
type MotherBoard struct{}

// 針對MotherBoard定義兩個主要的function Start(), Reboot()
func (*MotherBoard) Start() {
	fmt.Println("system starting")
}

func (*MotherBoard) Reboot() {
	fmt.Println("system rebooting")
}

// 定義開啟的指令-Struct
type StartCommand struct {
	mb *MotherBoard
}

// 定義重新開啟的指令-Struct
type RebootCommand struct {
	mb *MotherBoard
}

// 個別定義兩個Command的生成方式-這邊都會回傳對應的Struct
func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

/*
將前面定義好的Function放在Execute()這個方法下面
透過宣告StartCommand可以生成一個StartCommand的物件
該物件可以執行Execute()做對應的動作
*/
func (c *StartCommand) Execute() {
	c.mb.Start()
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

/*
開始時做Command模式
宣告一個Box資料結構
*/
type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}

func main() {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box2 := NewBox(rebootCommand, startCommand)

	box1.PressButton1()
	box1.PressButton2()

	fmt.Println()
	box2.PressButton1()
	box2.PressButton2()
}

package main

import "fmt"

/*
	定義好一週的天數迭代方法
*/
type Week interface {
	Today()
	Next(*DayContext)
}

type DayContext struct {
	today Week
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

func main() {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
}

/*
	預先定義好 Day Context
*/
type Sunday struct{}

func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct{}

func (*Monday) Today() {
	fmt.Printf("Monday\n")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct{}

func (*Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}

/*
	定義"Sunday"為一週的第一天
*/
func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

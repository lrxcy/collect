package main

import "fmt"

// 基本的component接口需要具備returnvalue()的方法
type component interface {
	returnvalue() int
}

// wraper包住component，讓component可以承接
type wraper struct {
	component
}

func (w *wraper) returnvalue() int {
	// 做return前的處理
	return w.component.returnvalue() + 999999
}

func wraperComponent(c component) component {
	return &wraper{
		component: c,
	}
}

func main() {
	i := &cstruct{
		rint: 1,
	}

	w := wraperComponent(i)
	fmt.Println(w.returnvalue())

}

type cstruct struct {
	rint int
}

func (c *cstruct) returnvalue() int {
	return c.rint
}

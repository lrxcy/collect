// refer website
// http://xiaorui.cc/2016/03/11/%E5%85%B3%E4%BA%8Egolang-struct-interface%E7%9A%84%E7%90%86%E8%A7%A3%E4%BD%BF%E7%94%A8/

package main

import (
	"fmt"
)

//定义了一个接口
type I interface {
	Get() int
	Put(int)
}

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

func f1(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

//interface{}空接口，能接受任何类型。.(I)是类型断言.
func f2(p interface{}) {
	if t, ok := p.(S); ok {
		fmt.Println("S:", t)
	} else if t, ok := p.(I); ok {
		fmt.Println("I:", t.Get())
	}
}

func f3(p interface{}) {
	switch t := p.(type) {
	case S:
		fmt.Println("S:", t.Get())
	case R:
		fmt.Println("R:", t.Get())
	case I:
		fmt.Println("I:", t.Get())
	default:
		fmt.Println("unknow type")
	}
}

func main() {
	s := S{101}

	f1(&s)
	f2(&s)

	r := R{1111}
	f3(&r)
}

package decorator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {
	// 第一次必須要先宣告一個component供後面的參數做迭代
	var c1 Component = &ConcreteComponent{
		a: 39,
		b: "string",
	}

	// 做第一次包裝，確認c2的回傳值
	c2 := WrapValidator(c1)
	res := c2.ReturnInt()
	res2 := c2.ReturnString()

	assert.Equal(t, 29, res)
	assert.Equal(t, "string test", res2)

	// 做第二次的包裝，會基於第一次的包裝，在做一次回傳
	c3 := WrapValidator(c2)
	res3 := c3.ReturnInt()
	res4 := c3.ReturnString()

	assert.Equal(t, 30, res3)
	assert.Equal(t, "string test test", res4)
}

// 宣告一個需要被使用到裝飾器的struct
type ConcreteComponent struct {
	a interface{}
	b interface{}
}

func (c *ConcreteComponent) ReturnInt() int {
	v, ok := c.a.(int)
	if !ok {
		panic(fmt.Sprintf("error happend while type error: %v", c.a))
	}
	return v
}
func (c *ConcreteComponent) ReturnString() string {
	v, ok := c.b.(string)
	if !ok {
		panic(fmt.Sprintf("error happend while type error: %v", c.b))
	}
	return v
}

func TestDecorator2(t *testing.T) {
	var cc1 Component = &AnotherConcreteComponent{
		a: 21,
		b: "test",
	}

	cc2 := WrapValidator(cc1)
	res := cc2.ReturnInt()
	res2 := cc2.ReturnString()
	assert.NotEqual(t, 21, res)
	assert.Equal(t, "test test", res2)

}

type AnotherConcreteComponent struct {
	a int
	b string
}

func (c AnotherConcreteComponent) ReturnInt() int {
	return c.a
}

func (c AnotherConcreteComponent) ReturnString() string {
	return c.b
}

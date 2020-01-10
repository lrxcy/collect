package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node 為一個基本運算元的乘載節點
type Node interface {
	Interpret() int
}

// ValNode為一個基本運算元
type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

/*
	以下分別定義出 "加法" 以及 "減法" 的運算節點
*/
// AddNode 代表加法的運算元
type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

// MinNode 代表減法的運算元
type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

/*
	Parser 代表一個解析器，會針對輸入的字串做解析，再依據內容調度合適的加減法運算元
*/
// Parser 為一個解析器
type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

/*
	Parser的 加法節點 及 減法節點
*/
func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newMinNode() Node {
	p.index++
	return &MinNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

// newValNode 回傳一個 有值的 Node
func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val: v,
	}
}

func (p *Parser) Result() Node {
	return p.prev
}

func main() {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()

	fmt.Println(res)
}

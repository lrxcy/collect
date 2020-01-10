package main

import "fmt"

/*
Component宣告一個interface，並且該方法具有
1. Parent() Component可以回傳 父類
2. SetParent(Component) 可以設定 父類
3. AddChild(Component) 可以設定 子類

-- 如何設定該Component的一些方法
4. Name() string
5. SetName() string

-- 以及一些挾帶的額外功能
6. Print(string)
*/
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

// 定義一些const來給後面宣告的leaf使用
const (
	LeafNode = iota
	CompositeNode
)

// 定義一個基礎的component 夾帶parent參數也為Component 以及 name為string
type component struct {
	parent Component
	name   string
}

type Leaf struct {
	component
}

type Composite struct {
	component
	childs []Component
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) AddChild(Component) {}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) Print(string) {}

/*
NewLeaf 回傳一個 Leaf實例的指標
NewComposite 回傳一個 Composite實例的指標
*/
func NewLeaf() *Leaf {
	return &Leaf{}
}

func (c *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

func main() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}

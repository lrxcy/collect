# composite 組合 模式

組合模式統一對象和對象集，使得相同接口使用對象和對象集

組合模式常用於樹狀結構，用於統一葉子節點和樹節點的訪問，並且可以用於應用某一操作到所有子節點。

> 使用情境: 

# refer:
- http://www.dotspace.idv.tw/Patterns/Jdon_Composite.htm

# keynotes

```go
1. 定義出Component的幾種方法

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName() string
	AddChild(Component)
	Print(string)
}

2. 定義出一些基本單位元，方便後面迭代使用，

const {
    LeafNode = iota
    CompositeNode
}

type component struct {
    parent Component
    name string
}

type Leaf struct {
	component
}

type Composite struct {
    component
    childs []Component
}

3. 定義出基本單位元component需要的方法Parent()/ SetParent(parent Component)/ AddChild(Component)/ Name()/ SetName(name string)/ Print(string)

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

4. 需要定義出一個生成Leaf的函數，以及顯示該Leaf值的方法

func NewLeaf() *Leaf {
    return &Leaf{}
}

func (c *Leaf) Print(pre string) {
    fmt.Printf("%s-%s\n", pre, c.Name())
}

5. 定義出Composite來實現後面要新增的Leaf

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

```
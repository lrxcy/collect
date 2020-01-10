# 解釋器模式
解釋器模式定義一套語言文法，並設計該語言解釋器，使用戶能使用特定文法控制解釋器行為。

解釋器模式的意義在於，它分離多種複雜功能的實現，每個功能只需關注自身的解釋。

對於調用者不用關心內部的解釋器的工作，只需用簡單的方式組合命令就可以。


### 介紹
意圖：給定一個語言，他定義他的文法表示，並定義一個解釋器。這個解釋器使用該標示來解釋語言中的句子
主要解決：對於一些固定文法構建一個解釋句子的解釋器
使用時機：如果一種特定類型的問題發生的頻率足夠高，那麼可能就值得將該問題的各個實例表述為一個簡單語言中的句子。

### 實作
如何解決：構建語法樹，定義終結符與非終結符
關鍵代碼：構建環境類，包含解釋器之外的一些全局信息，一般是HashMap
應用實例：編譯器、運算表達式計算

### 優點
1. 可擴展性比較好，靈活
2. 增加了新的解釋表達式的方式
3. 易於實現簡單文法

### 缺點
1. 可利用的場景比較少
2. 對於複雜的文法比較難維護
3. 解釋器模式容易引起類膨脹
4. 解釋器模式採用遞歸調用方法

### 場景
1. 可以將一個需要解釋執行的語言中的句子表示為一個抽象語法樹
2. 一些重複出現的問題可以用一種簡單的語言來進行表達
3. 一個簡單語法需要解釋的場景

# Unified Modeling Language
創建一個接口Expression和實現Expression接口的實體類。定義上下文解釋器的TerminalExpression類。其他的類OrExpression、AndExpression用於創建組合式表達式
```
InterpreterPatternDemo使用了Expression類 創建規則和表達式的解析

                            InterpreterPatternDemo
                            +getMaleExpression(): void
                            +getmarriedWomenexpression(): void
                                        |
                                        |
                                        uses
                                        |
                                        |
                            Expresion << interface >>
                            +interpret(): void
    ---------------------------------------------------------------------------
implements                          implement                              implement
    |                                   |                                      |
    |                                   |                                      |
TerminalExpression                  AndExpression                       OrExpression
-data: String                       -expr1: Expression                  -expr1: Expression
+TerminalExpression()               -expr2: Expression                  -expr1: Expression
+interpret(): booleam               +AndExpression()                    +OrExpression()
                                    +interpret(): bolleam               +interpret(): booleam

```

# refer:
- https://www.runoob.com/design-pattern/interpreter-pattern.html
- https://xyz.cinc.biz/2013/08/interpreter-pattern.html

# keynotes:
此處範例使用`ValNode`當作運算節點，並且定義對應的接口`Node`來乘載節點上的數據回傳實現方法
```go
type Node interface {
    Interpret() int
}

type ValNode struct {
    val int
}

func (n *ValNode) Interpret() int {
    return n.val
}
```

額外定義兩個運算解釋節點，方便後續定義解釋器(Parser)
```go
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
```

定義重點核心解釋器Parser
- exp: 存儲後面拆開的字元
- index: 緩存，解析後的字元陣列長度
- prev: 緩存，解析後的運算字元
```go
type Parser struct {
    exp []string
    index int
    prev Node
}
```

具體實現方法
```go
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
```

實踐解析器
```go
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
```
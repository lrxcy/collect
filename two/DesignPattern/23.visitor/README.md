# 訪問者模式

訪問者模式可以給一系列對象透明的添加功能，並且把相關代碼封裝到一個類中。

對象只要預留訪問者接口`Accept`則後期為對象添加功能的時候就不需要改動對象。

# 使用情境
假設一個系統，其中有一些相似類別。類別中都有某些方法內容相似，但還是需要判斷目前要做事的是哪個類別才能呼叫對應的適當類別。

- 當一個"物件結構"中的`元素`幾乎不會異動，但這些"元素的行為"常會增減，則適合用訪問者模式
- 醠問者模式是將"元素的行為"，提取出來，每一種行為做成一個`Visitor(訪問者)物件`
- 每一個`Visitor(訪問者)物件`，都能依據不同的`元素`，對應到不同的行為結果


# 範例
在一個遊戲結構中，有兩種人，誠實人、說謊人

誠實人都說真話，說謊人都說假話。

誠實人、說謊人即是不會異動的元素，

但對誠實人、說謊人`提出的問題`，是可能會改變的，

所以`提出的問題`即可當作`Visitor(訪問者)物件`，用訪問者模式(Visitor Pattern)，對以後要增減`提出的問題`是較方便的

(但若要新增第三種人則較麻煩)


# Unified Modeling Language
1. (優點)元素的個數是固定時(穩定的資料結構); (缺點)反之，資料結構(元素)的增加變得更為困難
2. 將處理和資料結構兩者分開來
3. 增加操作就等於是增加新的Visitor

```
                 Object Structure <---------------------------------------------------------->  Client
                        |                                                                         |
                        |                                                                         |
                     Element                                                                   Visitor
                     +Accept(invisitor: Visitor)                                               +VisitConcreteElementA(in: ConcreteElementA)
                        |                                                                      +VisitConcreteElementB(in: ConcreteElementB)
                        |                                                                         |
                        |                                                                         |
        ------------------------------------                                    -----------------------------------------
        |                                  |                                    |                                       |
        |                                  |                                    |                                       |
 ConcreteElementA                   ConcreteElementB                     ConcreteVisitorA                       ConcreteVisitorA
 +Accept(in visitor:Visitor)        +Accept(in visitor:Visitor)          +VisitorConcreteElementA               +VisitorConcreteElementA
 +OperatorA()                       +OperatorB()                         (in: ConcreteElementA)                 (in: ConcreteElementA)
                                                                         +VisitorConcreteElementB               +VisitorConcreteElementB
                                                                         (in: ConcreteElementB)                 (in: ConcreteElementB)
```

- ObjectStruct: 能枚舉他的元素，提供給Visitor訪問的介面
- Element: 以訪問者為參數的Accept操作
- Visitor: 為每一個具體的元素(Element)類別宣告一個Visit操作
- ConcreteVisitor: 需要對每一個元素實作具體的Visit行為
- ConcreteElement: 需要實作Accept方法，通常是指接受存取的方法的實作

> 如果需要頻繁的修改Visitor介面的話，代表不適合使用Visitor模式(資料結構不穩定)

# refer
- http://corrupt003-design-pattern.blogspot.com/2017/02/visitor-pattern.html 
- http://twmht.github.io/blog/posts/design-pattern/visitor.html
- https://ithelp.ithome.com.tw/articles/10208766
- https://xyz.cinc.biz/2013/08/visitor-pattern.html



# keynotes
定義出`Customer`及`Visitor`兩個接口。
- Customer (實現接口 Accept(Visitor)): 表示`不會變動`的物件，用來做`後端的邏輯處理`
- Visitor (實現接口 Visit(Customer)): 實作訪問邏輯，表示前面近來的訪問者是誰`用來處理後端的前端`

透過`CustomerCol`這個資料結構把`Customer`收集起來。這邊定義兩個方法`Add(customer Customer)`及`Accept(visitor Visitor)`

前者供應物件做串接以及新增，後者方便使用物件時做適當的調度

```go
type Customer interface {
    Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type CustomerCol struct {
    customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}
```

定義一些`Customer`，來提供後端行為
```go
type EnterpriseCustomer struct {
	name string
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

type IndividualCustomer struct {
	name string
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}
```


定義後端的前端邏輯層，使用以下幾個`Visitor`來實作
```go
type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// only for enterprise
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("%s is not an enterprise customer, no analysis would be applied\n", c.name)
	}
}
```

調度的時候，先宣告要使用的`CustomerCol{}`來儲存Customer物件，存儲的物件在`Accept`會一起被調度
```go
	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("A company"))
	c1.Add(NewEnterpriseCustomer("B company"))
	c1.Add(NewIndividualCustomer("bob"))
	c1.Accept(&ServiceRequestVisitor{})
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// serving individual customer bob

	c2 := &CustomerCol{}
	c2.Add(NewEnterpriseCustomer("A company"))
	c2.Add(NewIndividualCustomer("bob"))
	c2.Add(NewEnterpriseCustomer("B company"))
	c2.Accept(&AnalysisVisitor{})
	// Output:
	// analysis enterprise customer A company
	// bob is not an enterprise customer, no analysis would be applied
	// analysis enterprise customer B company
```
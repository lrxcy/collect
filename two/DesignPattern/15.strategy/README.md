# 策略模式
定義一系列算法，讓這些算法在運行時可以互相交換。使得算法分離，符合開閉原則。
> 用interface封裝公共介面
```
策略模式用 策略的介面 來替換在某個 實例 中的 方法，
藉由替換不同的 策略 使得 物件 擁有不同的 行為。
經過策略的組合，在運行中可以獲得行為不同的物件
```

# 優點
1. 靈活的替換不同的行為(or演算法)
2. 策略拓展容易
3. 避免使用很多if else

# 缺點
1. 必須自行決定要使用哪種策略
2. 可能因為產生很多的策略類，導致管理策略困難

# 架構
```
Context -> Strategy(+algorithm())
              |
        ______________________________
        |               |            |
        Concrete     Concrete     Concrete
        StrategyA    StrategyB    StrategyC
```


# refer
- https://ithelp.ithome.com.tw/articles/10202506



# keynote:
定義出的`(文本)資料結構`同時具有方法來操作文本內容，該方法為`策略接口`
```go
type PaymentContext struct {
    Name, CardID    string
    Money           int
    payment         PaymentStrategy
}

type PaymentStrategy interface {
    Pay(*PaymentContext)
}
```

透過定義出來的`PaymentContext`可以額外宣告不同的資料結構具有`Pay`這個方法，來操作文本
```go
type Cash struct {}

func (*Cash) Pay(ctx *PaymentContext) {
    fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct {}

func (*Bank) Pay(ctx *PaymentContext) {
    fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}
```

透過`NewPaymentContext(...)`這個方法來調用物件，並利用Cash/Bank物件下的Pay來實現策略模式
```go
func NewPaymentContext(name, cardid string, money int, payment PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		Name:    name,
		CardID:  cardid,
		Money:   money,
		payment: payment,
	}
}

// 使用方法，用New生成物件後使用Pay來調用
func main() {
	ctxc := NewPaymentContext("Ada", "", 123, &Cash{})
	ctxc.Pay()

	ctxb := NewPaymentContext("Bob", "002", 888, &Bank{})
	ctxb.Pay()
}
```
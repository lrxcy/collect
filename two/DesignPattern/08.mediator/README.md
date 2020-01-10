# 中介者模式
> Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it lets you vary their interaction independently.
- 中介者模式定義一個可以封裝一組物件互動的物件，可以使物件不用直接互相引用而降低耦合性，且可以獨立改變物件之間的互動關係。

中介者模式，又稱栛調者模式，直接照字面上解釋，就是有一個中介者負責處理事情。
因為有一個中介者可以負責事情，如處理、傳遞、通知，就可以簡化物件之間的溝通和控制制方式，進而降低物件之間的耦合性(相依性，依賴性)。

# Unified Modeling Language
```
    Mediator ------ mediator ------ Colleague
        |                               |
        |                               |
        |                               |
    ConcreteMediator           -----------------------
        |                      |                     |
        |                      |                     |
        |-------------- ConcreteColleague1   ConcreteColleague2
        |                                            |
        |                                            |
        |---------------------------------------------
```
類別圖中的Colleague 相關類別，指的是實際上要做的事，彼此可能因為不同的需求而會有相依性的類別。

ConcreteColleague 之間不會直接溝通，而是透過 Mediator。

### 與Facade模式比較
Facade 與 Mediator 同樣都具有管理一群類別，但不同之處在於Client的角度。
Facade 中 Client 是不會直接跟物件們溝通，都是透過表象來做事。而 Mediator 中 Client 可以`透過中介者`來讓達成讓物件彼此溝通

# refer:
- https://github.com/bvwells/go-patterns/blob/master/behavioral/mediator.go
- http://corrupt003-design-pattern.blogspot.com/2017/01/mediator-pattern.html

# extend-refer ... 莫名其妙的 ... 寫個design pattern還需要解決os.Stdout以及io.Writer的問題...
```go
/*
    因為原先定義的io.Writer為一個 接口，且該接口沒有辦法直接打印出結果到主控台
    所以讓outputWriter直接宣告等於os.Stdout
    原先作者在撰寫的時候，是拿來做測試用。但過度設計導致有礙學習...
    這邊把接口拿掉，改成 var outputWriter = os.Stdout
*/
type Writer interface{
    Write(p []byte) (n int, err error)
}

// 由於拿掉變數 outputWriter 中宣告為 io.Writer 的限制，所以也不用在new一個buffer給 outputWriter，所以可以在main函式中將其註解
```
- https://wiki.jikexueyuan.com/project/the-way-to-go/12.8.html
- https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
- https://golang.org/pkg/fmt/#Fprint
- https://my.oschina.net/solate/blog/1596263
- https://dev.to/dayvonjersen/how-to-use-io-reader-and-io-writer-ao0


# keynotes:
先定義出接口`Mediator`以及`WildStallion`
```go
Mediator: 實現 Communicate(who string) 來讓 資料結構 間可以互相溝通
WildStallion: 實現 SetMediator(mediator Mediator) 定義出 要呼叫以及使用的資料結構接口

type Mediator interface {
    Communicate(who string)
}

type WildStallion interface {
    SetMediator(mediator Mediator)
}

```

另外，為了確保所有的`中介層`可以實現兩個資料結構互相溝通，

定義出`ConcreatedMediator`來包括子物件，並且實現方法`Communicate(who string)`

確保可以被接口`WildStallion`所使用。
```go
type Bill struct {
    mediator Mediator
}

type Ted struct {
    mediator Mediator
}

type ConcreatedMediator struct {
    Bill
    Ted
}

func (m *ConcreateMediator) Communicate(who string) {
    if who == "Ted" {
        m.Bill.Respond()
    }else if who == "Bill"{
        m.Ted.Respond()
    }
}
```

因為前面定義了`Bill`以及`Ted`這兩個資料結構，

這兩個資料結構的母結構`ConcreatedMediatro`因為要被`WildStation`所涵括

所以裡面的子結構需要實現方法`SetMediator(mediator Mediator)`

以及一些額外的`溝通方法`，這邊範例使用的是`Respond()`

```go
// 定義子結構的SetMediator
func (b *Bill) SetMediator(mediator Mediator) {
    b.mediator = mediator
}

// 定義子結構的溝通方法Respond()
func (b *Bill) Respond() {
    fmt.Fprintf(outputWriter, "Bill: What?\n")
    b.mediator.Communicate("Bill")
}


func (t *Ted) SetMediator(mediator Mediator) {
    t.mediator = mediator
}

func (t *Ted) Respond() {
    fmt.Fprintf(outputWriter, "Ted: Bill?\n")
    b.mediator.Communicate("Ted")
}
```

最後，使用`NewMediator() *ConcreateMediator`來宣告以及開始建構一個中介者模式
```go
func NewMediator() *ConcreateMediator {
    mediator := &ConcreateMediator{}
    mediator.Bill.SetMediator(mediator)
    mediator.Ted.SetMediator(mediator)
    return mediator
}

```

後記...

這邊的NewMediator() *ConcreateMediator 並不是一個比較好的寫法，因為他綁死了資料結構。

所以在使用以前要先確保好，要使用的中介者有誰這件事。
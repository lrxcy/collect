# 責任鏈模式

責任鍊模式用於分離不同的職責，並且動態組合相關職責。

Golang實現責任鍊模式的時候，因為沒有繼承的支持，使用的鏈對象包含職責的方式，即：

- 練對象包含當前職責對象以及下一個責任鍊。
- 職責對象提供接口表示是否能處理對應請求。
- 職責對象提供處理函數處理相關職責。

同時可在責任鍊的類中實現責任鏈接口的相關函數，使責任鍊對象可以當作一般責任鍊對象使用。

# 情境

當收信的時候，收件者會依據信件的內容來對信件做對應的處置。

例如，垃圾信件會放置到信箱、會議信件會收到重要信箱，並且設定會議鬧鐘、來自客服信件會轉給工程部...

針對不同的信件做出不同的處置，即是一種責任鏈模式。

直接宣告一個類來處理是最簡單的方法，但勢必會用到大量的`if ... else ...` 進而導致更多信件種類進來時，判斷及維護處理信件系統的困難。

此時，需要把每個處理環節當作一個接口物件來做實現(`type newMail interface{}`)

> 讓多個物件都有機會可以處理請求，以避免請求的發送者和接受者之間產生耦合關係。將這些物件連成一條鏈，並沿著這條鏈傳遞請求，直到有一個物件處理這個請求為止。

>原文: Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handle it.

- 類別對於狀況判斷太多，承擔太多責任。造成未來新增或修改時必須違反`擴充開放`及`修改封閉`原則
- 發出請求的客戶端，不知道最終者是誰


### 實作思維
由小範圍到大範圍，特殊情形至一般情形來組織責任鍊

### Compare to Decorator: 責任鍊與裝飾器模式的差異
裝飾器模式的主要用途在`增加元件的行為`，而責任鏈模式是在`組織處理請求元件`

責任鍊可以將`請求的發出者和接受者之間予以鬆綁`。

請求者不需要知道實際接受者是誰，也不用知道請求是如何被處理，各個接受者之間也是彼此獨立且鬆綁的。

此外，責任鍊模式也可以`動態修改責任鍊`，如新增或刪除請求處理元件。

> 因此，很常用來使用在視窗程式中，處理像是滑鼠或是鍵盤事件。但，此模式沒有保證鏈下所有的物件一定都會被處理到。


# Unified Modeling Language
```
Client ------------>                Handler                   <--- successor----------------
                      + setSuccessor(int sueccessor: Handler)                               |
                      + HandleRequest(int request: int)                                     |
                                        |                                                   |
                                        |                                                   |
                -------------------------------------------------                           |
                |                                               |                           |
                |                                               |                           |
                |                                               |                           |
        ConcreateHandlerA                               ConcreateHandlerB         <---------|
  +HandleRequest(int request: int)                +HandleRequest(int request: int)    

```
- Handler: 是請求的接收者，Handler接收到請求(Requset)後，假如可以處理則處理，不能處理則將請求發送給後繼者
### 備註:
1. 有幾個物件都能處理某種請求時，個別物件能處理的範圍(權限)不會相同
2. 當這個物件沒有處理權限時，能夠將這個請求傳遞給下一個物件繼續處理

- 優點:
1. 降低耦合度(將請求的發送者和接收者解耦)
2. 簡化了對象，使得對象不需要知道鍊的結構
3. 增強給對象指派責任的靈活性，通過改變鏈內的成員或者調動他們的次序，允許動態的新增或刪除責任
4. 增加新的請求處理很方便
- 缺點:
1. 不能保證請求一定被接收
2. 系統性能將受到一定影響，而且在進行代碼調試時不太方便，可能會造成循環調用
3. 可能不容易觀察運行時的特徵，有礙於除錯


# refer:
- https://www.itread01.com/content/1543828626.html
- https://www.runoob.com/design-pattern/chain-of-responsibility-pattern.html
- https://xyz.cinc.biz/2013/07/chain-of-responsibility-pattern.html
- https://ithelp.ithome.com.tw/articles/10208172
- http://corrupt003-design-pattern.blogspot.com/2017/01/chain-of-responsibility-pattern.html



# keynotes:
定義一個data link資料結構`RequestChain`來做責任鍊基礎物件(具有迭代的功能)，並且附有`Manager`的實現方法。該接口
- HaveRight: 確認是否有錢
- HandleFeeRequest: 確認稅收是否允許
> 用抽象定義抽象: 定義 RequestChain 在處理 HandleFeeRequest 以及 HaveRight 的方法
額外對於`RequestChain`定義一個新增責任鍊物件的方法，`SetSuccessor(m *RequestChain)`
```go
type RequestChain struct {
    Manager
    successor *RequestChain
}

type Manager interface {
    HaveRight(money int) bool
    HandleFeeRequest(name string, money int) bool
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
    if r.Manager.HaveRight(money) {
        return r.Manager.HandleFeeRequest(name, money)
    }
    if r.successor != nil {
        return r.successor.HandleFeeRequest(name, money)
    }
    return false
}

func (r *RequestChain) HaveRight(money int) bool {
    return true
}

// 添加新的責任鍊的方法
func (r *RequestChain) SetSuccessor(m *RequestChain) {
    r.successor = m
}
```


以下定義三位不同的管理者，來做責任鍊相對物件
```go
type ProjectManager struct {}

type DepManager struct {}

type GeneralManager struct {}
```


- 三名管理者的生產方法方法
方法幾乎都是不帶輸入參數，回傳接口統一用`*RequestChain`
```go
func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}
```

- 實現資料結構下`RequestChain`的方法
1. Manager: 
    1. HaveRight(money int) boll
    2. HandleFeeRequest(name string, money int) bool

```go
// ProjectManager
func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}


// DepManager
func (*DepManager) HaveRight(money int) bool {
    return money < 500
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}


// GeneralManager
func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}
```

- 使用調度方式
透過先生成新的`RequestChain`，承接`SetSuccessor(*RequestChain)`
```go
    // 生成各個資料結構下的 RequestChain
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

    // 承接對應資料結構 RequestChain 的 下一個責任練物件
	c1.SetSuccessor(c2)
    c2.SetSuccessor(c3)
    
    // 將c1, ProjectManagerChain 當作頭
    var c Manager = c1

    // 呼應資料鏈的請求處理
	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
```

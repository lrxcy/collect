# flow
```go
main -> 
1. 實例 fc := fakeCheckout{}
2. 將先前實例的`fc`放進函數getRates()裡面

getRates(checkoutI InterfaceCheckout): 吃參數具有滿足InterfaceCheckout方法的資料結構
--> for _, item := range checkoutI.GetItems()  : 對應到checkoutI下面的方的方法`GetItems()`勢必會獲得一個陣列或是map的元素
    --> fmt.Println("%v\n", item.GetProduct())  : 對應到item底下底下的方法`GetProduct()`


-- 分析放入getRates的資料結構`fakeCheckout{}`
該資料結構會涵括`InterfaceCheckout`這個接口
type fakeCheckout struct {
	InterfaceCheckout
}

---- 分析fakeCheckout資料結構下的接口`InterfaceCheckout`
type InterfaceCheckout interface {
	GetID() int
	GetItems() []InterfaceCartItem
}

-------- 分析InterfaceCheckout接口底下實踐的兩個方法`GetID()`以及`GetItems()`
# notes: 其中`...`表示`fakeCheckout`，也代表一開始負入的參數所需具備的`方法`
func ( ... ) GetID() int 

func ( ... ) GetItems() []InterfaceCartItem {
	return []InterfaceCartItem{fakeItem{}, anotherFakeItem{}}
}

    -------- 分析`GetItems()`回傳的資料結構`[]InterfaceCartItem`內部的結構
    # notess: 函式中，具體回傳的資料結構是[]InterfaceCartItem{`fakeItem{}`, `anotherFakeItem{}`}
    type fakeItem interface {
        InterfaceCartItem
    }

    type anotherFakeItem interface {
        InterfaceCartItem
    }

-------- 分別定義出fakeItem以及anotherFakeItem內部InterfaceCartItem要實踐的方法
func (fakeItem) GetProduct() string {
    return "This is the end"
}

func (anotherFakeItem) GetProudct() string {
    return "This is another end"
}



```


# refer
- https://stackoverflow.com/questions/34956268/nested-interfaces-in-go?rq=1
- https://stackoverflow.com/questions/37484935/fulfilling-nested-interfaces-in-golang
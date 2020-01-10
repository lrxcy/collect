# state pattern
State Pattern 主要是解決控制一個物件狀態條件表達過於複雜的情況(簡化複雜的判斷邏輯)。

很多流程判斷幾乎都使用一連串 `if else`或 `switch` 來判斷某個狀態後得到什麼結果，

雖然一開始需求很少判斷條件寫起來也很少，當然 if else 自然也就沒啥問題，
--
但需求變動是不可避免的，後續只要需求有異動，

大部分的開發人員會很直覺的添加if else或 switch 條件，

以至於最後整個程式碼變的相當攏長也相當難維護，
--
而 State Pattern 將每個條件分支抽取成獨立類別，

每一個狀態都視成一個獨立物件(減少物件彼此依賴)，

這樣往後需求有變更時，大多也不用有太大的變更，

但由於把每個條件分支抽取成獨立類別，所以無法去避免類別變多的情況發生。

# Unified Modeling Language
```
                    Context(+request()) <---> State(+handle())
                    |                               |
state.handle() ------                               |
                                            ---------------------
                                            |                   |
                                        ConcreteStateA      ConcreteStateB
                                        (+handle())         (+handle())
```

# 使用場景
允許一個對象在內部改變它的狀態，並根據不同的狀態有不同的操作行為。
例如，水在固體、液體、氣體是三種狀態，但是展現在我們面前的確實不同的感覺。通過改變水的狀態，就可以更改它的展現方式。
- 應用場景
1. 當一個對象的行為，取決於它的狀態時;
2. 當類結構中存在大量的分支，並且每個分支內部的動作抽象相同，可以當做一種狀態來執行時。


# refer:
- https://dotblogs.com.tw/ricochen/2012/02/09/68609
- https://kknews.cc/code/6qbqggq.html

# keynotes
定義出的方法均具備能覆寫既有的資料結構的功能`Next(*DayContext)`
```go
type Week interface {
    Today()
    Next(*DayContext)
}

type DayContext struct {
    today week
}
```
後續定義出一連串用來代表改變狀態後的資料結構`Sunday/ Monday/ Tuesday/ Wedensday/ Thursday/ Friday/ Staurday`，且每個資料結構下的`Next(*DayContext)`均會覆寫既有的`today`物件
```go
type Sunday struct{}

func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}
```
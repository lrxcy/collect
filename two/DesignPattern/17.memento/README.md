# 備忘錄模式
忘錄模式的用途很單純，就是提供物件回到之前狀態的功能。
簡單說就是備份 (存檔) 的機制。備忘錄模式是一個在現實世界中很常使用到的模式，如遊戲的儲存記錄，文書編輯器的「上一步」功能等。

1. 備忘錄模式用於保存程序內部狀態到外部，又不希望暴露內部狀態的情形。
2. 程序內部狀態使用窄街口傳體給外部進行存儲，從而不暴露程序實現細節。
3. 備忘錄模式同時可以離線保存內部狀態，如同保存到數據庫，文件等。(永久性資料會存在資料庫(硬碟中)，而備忘錄模式是把某個物件得住太存在記憶體中，以便為來可以反悔)

# Unified Modeling Language
> Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later
在不違反封裝的情形下，取得物件的內部狀態。如此可以回復物件之前的狀態。
```
                        ------- Originator------------- Memento --------------Caretaker
                        |       state                   state
                        |       setMemento()            getState()
                        |       createMemento()]        setState()]
                        |           |
                        |           |
                        |           |
                        |        return new Memento(state)
                        |
state = m.getState() ---|
```
- Originator: 就是定義中提到要保留內部狀態的物件。現實例子就像是遊戲角色狀態，或是文書編輯器中的文字等。
- Memento: 保留Originator內部狀態(資料)的物件，例如遊戲中要存擋的資料。
- Caretaker: 主要功用是管理Memento物件。




# refer:
- http://corrupt003-design-pattern.blogspot.com/2017/02/memento-pattern.html
- https://ithelp.ithome.com.tw/articles/10206939


# keynotes
1. 額外宣告一個物件(progress ... 類別為 `gameMemento`並且實踐interface `Memento` )來存儲目前的狀態`progress:=game.Save()`
2. 並在需要回復狀態的時候進行讀取`Load(m Memento)`

```go
type Memento interface{}

type Game struct {
    hp, mp int
}

type gameMemento struct {
    hp, mp int
}
```

假設有一個遊戲的資料結構`Game`並且以`gameMemento`來做該遊戲的備忘錄，定義對應的方法

```go
func (g *Game) Play(mpDelta, hpDelta int) {
    g.mp += mpDelta
    g.hp += hpDelta
}

func (g *Game) Save() Memento {
    return &gameMemento {
        hp: g.hp,
        mp: g.mp,
    }
}

func (g *Game) Load(m Memento) {
    gm := m.(*gameMemento)
    g.mp = gm.mp
    g.hp = gm.hp
}

func (g *Game) Status() {
    fmt.Printf("Current HP: %d, MP: %d\n", g.hp, g.mp)
}
```

# observer 觀察者 模式
定義：觀察者模式是一種一對多的依賴關係，當物件的狀態改變，所有依賴於它的物件都會得到通知並被自動更新
> 使用情境: 當有一系列的物件可能皆具有某種方法，程式設計師會希望透過一個list或array來管理這些物件。並且期望任何的更動能夠對該list上的物件做統一的操作

# refer:
- https://en.wikipedia.org/wiki/Observer_pattern
- https://ithelp.ithome.com.tw/articles/10204117


# keynote
一個觀察者模式需要提供一個Register(收容Observer)，以及一個Observer的interface好方便在Register收入Observer以後對底下的所有觀察者作統一的行為。
```go
// 1. 首先定一Observer的interface以及Register

type Observer interface {
    Update(*Register)
}

type Register struct {
    observers []Observer
    context string
}

// 2. 定義生成出來的Register，需要涵括一個可以容納Obserrver的array

func NewRegister() *Register {
    return &Register{
        observers: make([]Observer, 0),
    }
}

// 3. 提供在Register上新增Observer的方法以及對於所有Observer的操作

func (s *Register) Attach (o Observer) {
    s.observers = append(s.observers, o)
}

func (s *Register) notify() {
    for _, o := range s.observers {
        o.Update(s)
    }
}

func (s *Register) UpdateContext(context string) {
    s.context = context
    s.notify()
}

// 4. [使用方法]在這邊的範例提供一個Reader當作Observer，並且讓Reader具有Update這個方法，以利後面被Register做append via Attach method

type Reader struct {
    name string
}

func NewReader(name string) *Reader {
    return &Reader{
        name: name,
    }
}

func (r *Reader) Update(s *Register) {
    fmt.Printf("%s receive %s\n", r.name, s.context)
}
```
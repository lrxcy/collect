# command 命令模式
命令模式(Command Pattern)有三個主要角色Invoker、ICommand和Receiver，是將對行為進行封裝的典型模式，將命令的`命令接收(請求操作者)`跟`執行命令(實際操作者)`之間切分開來

幾乎所有的類別都可以套用命令模式，但是只有在需要某些特殊功能，如`記錄操作步驟`、`取消上次命令`的時候

> 使用情境: 把某些特殊指令封裝起來，提供上層使用者直接呼叫使用
```
Architecture:

Invoker +invoker() ==> <interface>ICommand +execute()
                         A
                         |
                         |
                         V
Receiver +action() <== Concrete Command +execute() 
```

# refer:
- https://ithelp.ithome.com.tw/articles/10204425


# keynote:
Command 模式需要定義一個interface來執行對應的method
> 定義主要的物件，及該物件會需要的一些基本操作。接著封裝這些操作。
```go
// 1. 定義Command interface讓他下面具有方法Execute()，後面會需要透過呼叫Execute()這個方法來執行確實的指令

type Command interface {
    Execute()
}

// 2. 定義主要操作的物件，以及該物件預計的一些方法

type MotherBoard struct {}

func (*MotherBoard) Start() {
    fmt.Println("system starting")
}

func (*MotherBoard) Reboot() {
    fmt.Println("system rebooting")
}

// 3. 透過宣告StartCommand以及RebootCommand兩個struct，並且在struct裡面埋入MotherBoard調度該物件的method，後面可以直接使用Execute()方法調度對應的功能

type StartCommand struct {
    mb *MotherBoard
}

type RebootCommand struct {
    mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
    return &StartCommand {
        mb: mb,
    }
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
    return &RebootCommand {
        mb: mb,
    }
}

func (c *StartCommand) Execute() {
    c.mb.Start()
}

func (c *RebootCommand) Execute() {
    c.mb.Reboot()
}


/* ... 至此, 可以允許使用Command模式的前置條件已經完成... */


// 4. 宣告要使用Command模式的Box並且制定，對應的執行模式

type Box struct {
    button1 Command
    button2 Command
}

func NewBox(button1, button2 Command) *Box {
    return &Box{
        button1: button1,
        button2: button2,
    }
}

func (b *Box) PressButton1() {
    b.button1.Execute()
}

func (b *Box) PressButton2() {
    b.button2.Execute()
}


--
實作調度
1. NewMotherBoard物件
2. NewStartCommand & NewRebootCommand物件
3. NewBox物件並且賦予對應的Command
4. 呼叫Box下的Method

```
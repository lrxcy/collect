# 橋接模式

橋接模式分離抽象部分和實現部分。使得兩部分獨立擴展。

橋接模式類似於策略模式，區別在於策略模式封裝一系列算法使得算法可以互相替換。

策略模式使抽象部分和實現部分分離，可以獨立變化。


# 使用情境
當需求分成"前端輸入"和"列印訂單"的時候，通常後端事負責後者，至於訂單的來源都已經在資料庫了。

如果使用策略模式雖然可以達成需求，但是訂單除了各家不同還需要區分產品性質"一般"和"急件"

- Strategy屬於行為模式(Behavioral design patterns)
- Bridge屬於結構型模式(Structural design patterns)，他將抽象何時做解耦合，使兩者可以獨立的變化

橋接模式把抽象何時做分開，如此他們兩個可以互相獨立
> Decouple an abstraction from its implementation so that the two can vary independently

橋接模式允許你改變實作及抽象，採取的作法是將實作與抽象放在兩個不同的類別階層中


# Unified Modeling Language
1. 分離Abstraction及Implementor，使兩者可獨立變化。
2. 在run time設定Abstraction裡的Implementor
```
	  Client
		|
		|
	Abstraction  ------------------------------- Has-A ------------------------->  Implementor
	+Operation()						                                           +OperationImp()
	    |																			|		|
		|---------------------------|												|		|
		|			           		|												|		|
		|			           		|									     |-------       --------|
        |               Implementor OperationImp()							 |					    |
		|																	 |					    |
		|														   ConcreteImplimentorA		ConcreateImplimentorB
		|														   +OperationImp()			+OperationImp()
		|
	RefinedAbstraction
```
- Abstraction: 定義抽象的接口; 該接口包含實現具體行為、具體特徵的Implementor接口
- Refined Abstraction: 抽象接口Abstraction的子類，依舊是一個抽象的事物名
- Implementor: 定義具體行為、具體特徵的應用接口
- ConcreteImplementor: 實現Implementor接口

橋接模式的優點
1. 將實作跟介面鬆綁
2. 抽象跟實作可以各自擴充，不會影響到對方
3. 對於"具象的抽象類別"所做的改變，不會影響到客戶


# refer:
- https://ithelp.ithome.com.tw/articles/10193914
- http://corrupt003-design-pattern.blogspot.com/2017/01/bridge-pattern.html
- https://openhome.cc/Gossip/DesignPattern/BridgePattern.htm
- https://zh.wikipedia.org/wiki/%E6%A9%8B%E6%8E%A5%E6%A8%A1%E5%BC%8F
- https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/74929/


# keynotes:
假設今天要發送訊息給指定人員：

可以透過`簡訊`或`郵件`來寄送

而寄送的等級又可以分成`一般`或`急件`

- 定義出 簡訊/郵件 的本質: 1. `ViaSMS` 2. `ViaEmail` ===> MessageImplementer
- 定義出 簡訊/郵件 的等級: 1. `NewCommonMessage` 2. `NewUrgencyMessage` ===> AbstractMessage

```go
type MessageImplement interface {
    Send(text, to string)
}

type AbstractMessage interface {
    SendMessage(text, to string)
}
```

製作對應的 簡訊/郵件 物件 的資料結構及生成方法
```go
type MessageSMS struct {}

type MessageEmail struct {}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}
```

賦予該對應物件的相對生成方式
```go
func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS\n", text, to)
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}
```

製作對應的 簡訊/郵件 的等級 的資料結構及生成方法
```go
type CommonMessage struct {}

type UrgencyMessage struct {}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}
```

賦予該對應物件的相對生成方式
```go
func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
```

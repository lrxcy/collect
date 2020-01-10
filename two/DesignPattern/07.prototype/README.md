# prototype 原型 模式
原型模式使對象能複製自身，並且提供interface，讓使用者可以透過interface創建新的對象。原型模式通常會配合`管理器`使用，使得客戶端在不知道具體類的情況下，通過interface得到新的實例，並且包含部分預設配置。
> 使用情境: 在程序執行中，有基層物件。需要被儲存，方便後續拿取以及使用

# refer:
- https://www.itread01.com/content/1545065404.html
- https://www.cnblogs.com/gaochundong/p/design_pattern_prototype.html

# keynotes
```go
// 1. 定義出一個Cloneable interface，方便讓後面的 PrototypeManager(原型管理者)可以用mapping的方式儲存

type Cloneable interface {
    Clone() Cloneable
}

// 2. 定義出原型管理者`PrototypeManager`，以及其生成的方式

type PrototypeManager struct {
    prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
    return &PrototypeManager{
        prototypes: make(map[string]Cloneable),
    }
}

// 3. 針對原型管理者`PrototypeManager`定義出`增加`以及`拿取`原型的方法

//// -1. 增加 Set(name string)
func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

//// -2. 拿取 Get(name string)
func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

```

```go
/*
- 關於如何創造一個原型以及拿取一個原型態
1. 創造一個原型態
- 1. 宣告一個原型的資料結構
- 2. 給予該資料結構一個Clone() Cloneable 的方法(...通常是返回自己)
*/
type Type1 struct {
    name string
}

func (t *Type1) Clone() Cloneable {
    tc := *t
    return &tc
}

```
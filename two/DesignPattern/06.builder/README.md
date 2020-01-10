# builder 建造者模式
使用情境: 當需要彈性的控制每個生成物件下的狀態，甚至動態調整時。會給予該生成物件一定的參數配置，並且定義一些方法來動態調整該參數
```
Director -> Builder
                |
            ConcreteBuilder -> Product

1. 指揮者(Director)直接和客戶(Client)進行需求溝通
2. 溝通後指揮者將客戶創建產品的需求劃分為各個部件的建造請求(Builder)
3. 將各個部件的請求委派到具體的建造者(ConcreteBuilder)
4. 各個具體間早者負責進行產品部件的構建
5. 最終構建成具體產品(Product)
```

# refer:
- https://litotom.com/2016/07/07/builder-design-pattern/
- https://blog.csdn.net/carson_ho/article/details/54910597

# keypoint
```go
// 1. 先生成一個生成器街口
type Builder interface {
    Part1()
    Part2()
    Part3()
}

// 2. 在定義出Director的資料結構，讓Direcotr具有Builder的method
type Director struct {
    builder Builder
}

func NewDirecotr(build Builder) *Director {
    return &Director {
        builder: builder,
    }
}

// 同時定義出Direcotr的ConcreteBuilder方法
func (d *Director) Construct() {
    d.builder.Part1()
    d.builder.Part2()
    d.builder.Part3()
}

// 3. 製作各式各類的builder
type Builder1 struct {
    result string
}

func (b *Builder1) Part1() {
    ...
}

func (b *Builder1) Part2() {
    ...
}

func (b *Builder1) Part3() {
    ...
}

// 額外多定義一個GetResult() string來返回b.result
func (b *Builder1) GetResult() string{
    return ...
}

==
備註:
對於不同的Builder可以宣告不同的GetResult() interface{}

```
# factory_method 工廠方法
簡單工廠管理物件的創造，如果client要取得物件，只要給簡單供醠正確的參數就可以
```
PizzaBase: spicy, vegetarian
SimplePizzaFactory: [PizzaFactory():] 
1. SpicyPizza
2. VegetarianPizza
--> Pizza: [AddSpicy(), RemoveMeat(), Result()]
```

# refer:
- https://blog.techbridge.cc/2017/05/22/factory-method-and-abstract-factory/


# keypont:
在工廠模式，會使用到interface創建interface的情境
```go
// 1. 先定義廣義的Pizza屬性
type Pizza interface { 
    AddSpicy()
    RemoveMeat()
    Result()
}

// 2. 利用上面定義過的Pizza屬性，製作一個Pizza工廠，給後面各種口味的Pizza調用
type PizzaFactory interface {
    CreatePizza() Pizza // 使用 CreatePizza() 會返回一個 Pizza 的interface
}

// 3. 宣告一個PizzaBase供後面製作的Pizza做父類，主要給後面不同口味的Pizza做繼承用
//// -1. 該PizzaBase可以涵括一些通用的方法，供後面不同口味的Pizza做使用
//// -2. 務必確認該 `PizzaBase` 與後面生成的 `不同口味的Pizza`，需要能夠被Pizza interface這個接口所承接

type PizzaBase struct {
    spicy bool
    vegetarian bool
}

func (p PizzaBase) AddSpicy() {
    p.spicy = true
}

func (p PizzaBase) RemoveMeat() {
    p.vegetarian = true
}


// 4. 針對各種口味的Pizza製作專屬的struct，但因為前面PizzaBase只有定義了`AddSpicy()`以及`RemoveMeat()`所以尚需要定義`Result() string`
//// -1. 先定義出 SpicyPizza
//// -2. 定義出 SpicyPizza的 `Result() string` 函數

type SpicyPizza struct {
    *PizzaBase
}

func (o *SpicyPizza) Result() string {
	o.AddSpicy()                            // 在裡面調用AddSpicy()是可以的，因為繼承了PizzaBase的方法
	log.Println("the pizza is add spicy")
	return fmt.Sprintf("spciy: %v; vegetarian: %v\n", o.spicy, o.vegetarian)
}

// 5. 開始製作工廠接口，方便使用時直接呼叫接口即可生成該口味Pizza
type OrderSpciyPizza struct{}

func (OrderSpciyPizza) CreatePizza() Pizza {
	return &SpicyPizza{
		PizzaBase: &PizzaBase{},
	}
}


-- 備註：
這邊在生成該SpicyPizza時，使用到了原本第二步所使用的Interface: PizzaFactory 這一類
利用這個方法，可以在創建時直接透過宣告該物件，並調度方法`CreatePizza()`就可以生成一個對應口味的Pizza
其次，因為該口味Pizza具有事先定義的 type Pizza interface 的 `Result` 方法，所以可以直接調用 `Result()`來得到結果

```
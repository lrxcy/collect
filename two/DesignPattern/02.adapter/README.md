# adapter 配適器模式
這個模式可以讓客戶和被轉接者之間是鬆綁的，他們並不認識彼此，而客戶接收到呼叫的結果的時候，並未察覺這一切是透過一個轉接器中介傳導。
> 假設ad

```
Target : 配適的目標接口

architecture:


Client{客戶只看到目標介面} --> Target: request() void
                                    |
                                    |
                            Adapter: request() void --> Adaptee {轉接器與被轉接者合成}: SpecificRequest() void
```


# refer:
- https://dotblogs.com.tw/pin0513/2010/05/30/


# keypoint
使用adapter要先搞清楚需求，從最末端物件開始定義會有助於架構程式設計。
1. 先定義出plug表示插頭
2. 在定義出plugAdaptee表示插頭轉接器
3. 最後依據`生成插頭`，`生成插頭轉接器`來使用adapter
4. 在定義`生成插頭轉接器`的時候，要記得把`插頭`當作輸入參數之一
```go
// 1. 定義出插頭，給出該插頭的充電方法/訊息以及生成該插頭的方法
type PlugImple interface {
        Charging() string
}

type Plug struct{}

func (*Plug) Charging() string {
        return "device is charging"
}

func NewPlug() PlugImple {
        return &Plug{}
}

// 2. 定義出插頭轉接器
//// -1. 轉接器使用的資料結構需要囊括插頭(這邊是直接拿interface放入，方便後面定義多種插頭)
type PowerAdaptee struct{
        PlugImple
}

//// -2. 轉接器會提供轉接器上面的方法供裝置使用
type PowerAdapteeImpl interface{
        Charge() string
}

func (a *PowerAdaptee) Charge() string {
        return a. Charging()   // 把PowerAdaptee裡面的方法Charging()透過a物件調用出來
}

//// -3. 實作生成插頭轉接器函數
func NewPowerAdaptee(pluginImple PlugImple) PowerAdapteeImpl {
        return &PowerAdaptee{
                PlugImple: pluginImple,
        }
}

```
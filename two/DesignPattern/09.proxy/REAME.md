# proxy 代理 模式
代理模式用於延遲處理操作，或者再進行實際操作前(或後)額外進行其他處理(e.g. header驗證，資料正規化)
> 使用情境: 當某隻api server不想被第三方直接請求時，可以透過代理模式，在api server前做反向代理在對api做請求。又或者某些物件，不希望直接被調用。而是透過一些合法的代理來使用調度時。(某種程度上來說Data Access Object也是代理模式的一種)
```
client -> interface <-> Proxy(request)
                            |
                            |
RealSubject(request) <-------
```

# refer:
- http://twmht.github.io/blog/posts/design-pattern/proxy.html
- http://corrupt003-design-pattern.blogspot.com/2016/10/proxy-pattern.html


# keynote:
```go
// 1. 先定義一個真實物件RealSubject，並且給該物件調度方法Do() string

type RealSubject struct{}

type Subject interface {
    Do() string
}

// 2. 透過定義一個代理物件資料結構，繼承真實物件的資料結構並且重新定義方法Do() string

type Proxy struct{
    real RealSubject
}

func (p Proxy) Do() string{
    var res string
    ...
    res += p.real.Do() //調用真實對象
    ...
    retrun res
}

```


## 代理模式的常见用法有

* 虚代理
* COW代理
* 遠程代理
* 保護代理
* Cache 代理
* 防火牆代理
* 同步代理
* 智能指引

等。。。

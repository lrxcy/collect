# decorator 裝飾器模式
動態的將責任加諸於物件上
> 如果今天伺服器每次都需要驗證某些固定的身份，或者在計算某些迭代數列(費波曼數列)。可以透過裝飾器模式，預先確認來的請求的身份驗證，或是在做數列迭加時先確認該數列是否已經計算過並且緩存在記憶體之中。

```
1. 抽象物件(Component): 定義了對象的接口，可以給這些對象動態增加職責(方法)
2. 具體物件(ConcreteComponent): 定義了具體的構件對象，實現了在抽象構件中聲明的方法，裝飾器可以給它增加額外的職責(方法)
3. 抽象裝飾物件(Decorator): 抽象構件的子類，用於給具體構件增加職責，但是具體職責在其子類中實現。
4. 具體裝飾物件(ConcreteDecorator): 抽象裝飾類的子類，負責向構件添加新的職責。
```


# refer
- https://rongli.gitbooks.io/design-pattern/content/chapter8.html
- https://openhome.cc/Gossip/DesignPattern/DecoratorPattern.htm
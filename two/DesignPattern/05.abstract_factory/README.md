# abstract_factory 抽象工廠模式
抽象工廠模式用於，當生成的工廠之間是有關聯的時候。
> 用一個抽象工廠來定義一個創建 `產品族` 的介面，產品族裡面每個產品的具體類別由繼承抽象工廠的實體工廠決定。

# refer
- https://blog.techbridge.cc/2017/05/22/factory-method-and-abstract-factory/

# keynotes:
抽象工廠旨在從結果反推回生成過程，並在呼叫生成端與結果中間取平衡。而定義出一個合適的抽象介面
```
承接04.factory_method，在PizzaFactory前在定義一層interface叫做Restaurant，
並且給不同的Restaurant定義出可以生產的對應的Pizza口味

在這邊Restaurant就是抽象工廠的代表
```
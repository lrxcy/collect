# intro
有時候，我們需要去分類一些非預設排序的狀況，假使我們想要依據`字句的長度`來排序。而非`字句的字首`來排序，可以透過`sort`函數來完成。

# refer:
- https://gobyexample.com/sorting-by-functions

# refer_generic:
- https://michaelchen.tech/golang-programming/generics/

# 泛型(generics)
是一種多型的模式，透過泛型程式，可將同一套程式碼用到不同型別的資料上。
> 基本上GO不支援泛型

# 少數的泛型支援
直接說Go不支援泛型是過於簡化的說法；其實GO在內建的`陣列`、`切片`、`map`是支援的


# 替代策略
1. 介面
2. 空介面
3. 程式碼生成
4. 發展新語言

# 替代---介面
```golang
type strgeneric []string

func (s *strgeneric) Len() int{
    ...
    
    return ...
}
```

# 空介面
```golang
type List struct{
    head *node
    tail *node
}

type node struct {
    data interface{}
    next *node
    prev *node
}
```
在範例中`data`為空介面，可以放入任意的資料型別。

對於和資料本身無關的操作，不需要處理議題。例如，`Push`方法在串列尾端加入一個元素：
```golang
func (l *List)Push(data interface{}){
    node := node{data: data, next:nil, prev:nil}
    if list.head == nil{
        list.head = &node
        list.tail = &node
    }else {
        list.tail.next = &node
        node.prev = list.tail
        list.tail = &node
    }
}
```

當某些操作和資料相關時。用函數程式設計的方法，將資料相關的運算`委任(delegation)`給套件使用者；
以`Select`過濾掉不符合條件的元素
```golang
// Excerpt
// Select items from the list when item fulfill the predicate.
// Implement grep function object to use this method.
func (list *List) Select(grep func(interface{}) (bool, error)) (*List, error) {
    newList := New()
 
    current := list.head
    for current != nil {
        b, err := grep(current.data)
        if err != nil {
            return newList, err
        }
 
        if b == true {
            newList.Push(current.data)
        }
 
        current = current.next
    }
 
    return newList, nil
}
```
假設資料是可拷貝的(cloneable)，因Go不支援重載函式，若明確呼叫`Clone`方法會使得大部分的內建型別無法使用，這是設計上的考量。
對於沒有內建拷貝機制的型別，要由使用者將資料傳入串鏈前即拷貝一份。

```golang main.go
package main

import (
    "errors"
    "log"
    "github.com/cwchentw/algo-golang/list"
)

func main() {
    l := list.New()

    for i := 1; i<= 10; i++ {
        l.Push(i)
    }

    evens, _ := l.Select(isEven)
    for e := range evens.Iter() {
        n, _ := e.(int)
        if n % 2 != 0 {
            log.Fatalln("Error Select result")
        }
    }
}

func isEven(a interface{}) (bool, error) {
    n, ok := a.(int)
    if ok != true {
        return false, errors.New("Failed type assertion on a")
    }

    return  n % 2 ==0, nil
}
```

總結：
使用介面和空介面都需要透過撰寫一些樣板程式碼，只是
1. 前者是物件導向的風格
2. 後者是使用函數式程式碼來撰寫
但是大多開法者不會使用後者，可能是因為函數表達方式筆物件導向方式難以離解。
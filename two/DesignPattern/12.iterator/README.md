# iterator 迭代器模式

迭代器模式用於使用相同方式迭代不同類型集合或者隱藏集合類型的具體實現

可以讓使用者透過特定的介面巡訪容器中的每一個元素而不用了解底層的實作

# refer:
- https://en.wikipedia.org/wiki/Iterator_pattern
- http://www.runoob.com/design-pattern/iterator-pattern.html

# keynotes
```go
// 1. 先定義一個迭代器的interface，使用該interface所預先宣告的三個方法來嘗試做迭代

type Iterator interfacce {
    First()
    IsDone() bool
    Next() interface{}
}


// 2. 利用迭代器(Iterator)創建一個累加器(Aggregate): 繼承迭代器的interface

type Aggregate interface {
    Iterator() Iterator
}

// 3. 宣告一個Numbers結構，方便後面要實作Numbers迭代器

type Numbers struct {
    start, end int
}

func NewNumbers(start, end int) *Numbers {
    return &Numbers{
        start: start,
        end:    end,
    }
}

// 4. 宣告NumbersIterator

type NumbersIterator struct {
    numbers *Numbers
    next int
}

func (i *NumbersIterator) First() {
    i.next = i.numbers.start
}

func (i *NumbersIterator) IsDone() bool {
    return i.next > i.numbers.ned
}

func (i *NumbersIterator) Next() interface{} {
    if !i.IsDone() {
        next := i.next
        i.next++
        return next
    }
    return nil
}

5. 宣告Numbers的Iterator方法

func (n *Numbers) Iterator() Iterator {
    return &NumbersIterator{
        numbers: n,
        next:   n.start,
    }
}

// 製作一個func打印iterator裡面的數字
func IteratorPrint(i Iterator) {
    for i.First(); !i.IsDone(); {
        c := i.Next()
        fmt.Printf("%v\n",c)
    }
}

```
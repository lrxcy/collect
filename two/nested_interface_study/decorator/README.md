# flow
```go
main ->
1. 實例 i := &cstruct{ rint: 1}，具有方法
> func (c *struct) returnvalue() int{return c.rint}
2. 將先前的實例包裝起來 `w := wrapperComponent(i)`
3. 呼叫之前包裝起來的實例方法`returnvalue()`

wrapperComponent(i)，將i包裝起來

func wrapperComponent(c component) component {
    return &wraper{
        component: c,
    }
}

此時物件一樣包在`component`這個接口底下，但是呼叫的對象從原本的`i`變成`w`
所以在下次接下來的`returnvalue()`是採用func (w *wraper) returnvalue() int {return w.component.returnvalue() + 999999}

```
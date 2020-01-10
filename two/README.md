# Golang
hi 大家好，我是Jim。這是一篇記錄著我工作上用go語言開發所遇到的一些故事。希望能藉由一點一滴的紀錄，把自己的成長過程記錄下來，或是幫助有緣的您。

## 基本類型
- Go的基本類型有basic types
1. bool
2. string
3. int  int8  int16  int32  int64
4. uint uint8 uint16 uint32 uint64 uintptr
5. byte // uint8 的別名
6. rune // int32 的別名
>     // 代表一個Unicode碼
7. float32 float64
8. complex64 complex128

## go語言的一些特性
1. 一級函式(First-class function): https://openhome.cc/Gossip/Go/FirstClassFunction.html
2. 匿名函式與閉包(Closure): https://openhome.cc/Gossip/Go/Closure.html

## Reference website
- https://openhome.cc/Gossip/Go/Package.html
- https://golang.org/doc/code.html
- https://gobyexample.com/

### GO project overview
1. Go programmers typically keep all their Go code in a single workspace.
> Go 程式編譯者只有一個工作環境
2. A workspace contains many version control repositories (managed by Git, for example).
> 該工作環境可以透過Git管控達到擁有多個專案的版本管控
3. Each repository contains one or more packages.
> 每一個專案都可以涵括一個以上的套件包
4. Each package consists of one or more Go source files in a single directory.
> 每一個套件包內可能都是許多的 .go 檔
5. The path to a package's directory determines its import path.
> 要 import 的套件包必須放在他指定的路徑下


### Advance GO
- Discussion: Concurrency is not Parallelism
https://talks.golang.org/2012/waza.slide#1
> example : https://medium.com/@thejasbabu/concurrency-patterns-golang-5c5e1bcd0833

- lightweight loadbalance
https://github.com/yyyar/gobetween

- check goroutine : stacktrace
https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/

- detect goroutine leaks
https://blog.minio.io/debugging-go-routine-leaks-a1220142d32c

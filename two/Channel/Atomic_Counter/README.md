# intro
Go中首要機制就是channels之間的溝通，可以透過worker pools
```go
func worker(id int, jobs <-chan int, result chan<-int){
    // 透過range chane，可以把jobs裡面的數字拉出來
    for j := range jobs{
        fmt.Printf("worker %v start jobs _%v\n", id, j)
        ...
        fmt.Printf("worker %v finished jobs _%v\n", id, j)
        // 可以對jobs裡面的數字`j`做一些計算，在儲存回chan results
        results <- j*2
    }
}
```
但是，隨著平行化帶來的議題是同步與否。便需要使用`sync/atomic`包來實踐。

此處使用WaitGroup以及WaitGroup.Add來確保等待所有goroutine會做完。透過`atomic.AddUint64(&ops, 1)`來讓所有goroutine對ops做加1的動作。


# refer:
- https://gobyexample.com/atomic-counters
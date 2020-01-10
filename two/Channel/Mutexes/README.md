# 代碼介紹
開啟一個`state = map[int]int`，並且反覆寫入，以及讀取值。
透過mutex.Lock()以及mutex.UnLock()來實踐，確保每次都只有一個人去操作state的值這件事。


1. 啟用10個goroutine來模擬執行write這件事
   1. 每個(write)goroutine會以間隔time.Millisecond的時差來對writeOps執行`+1`的動作
   2. 每次在執行`state[key] = val`的動作時，確保彼此是獨立的。會啟用`mutex.Lock()`來確保`state`是只有被一個人寫入，或是拿到
      1. 在這邊`state[key] = val`裡面的值，會在(read)goroutine中被拿到，也可能在下次寫入時被另一個goroutine覆寫
   3. 每次拿取完後，會執行`mutext.Unlock()`來確保其他人可以在對`state`做操作
2. 啟用100個goroutine來模擬執行read這件事
   1. 每個(read)goroutine會以間隔time.Millisecond的時差來對readOps執行`+1`的動作
   2. 每次在執行`total += state[key]`的動作時，確保彼此是獨立的。會啟用`mutex.Lock()`來確保`state`是只有被一個人拿到，或是寫入
      1. total不重要，重要的是。每次都會有人去拿state[key]這件事。但是會因為另一個goroutine在模擬寫入這件事。導致拿取特定值的時候報錯。
   3. 每次拿取完後，會執行`mutex.Unlock()`來確保其他人可以在對`state`做操作


# refer:
- https://gobyexample.com/mutexes
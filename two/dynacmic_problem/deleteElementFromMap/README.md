# 在golang中刪除map裡的元素
因為golang對於delete map的行為是把key值歸到nil，並不會重新allocate map。
所以需要額外對map做resize。但是目前golang語言對於該問題的處理還沒有一個很好的解決辦法。

目前的解決辦法是，在額外宣告一個Nmap，把所有舊Omap的所有資料倒入Nmap後，把舊Omap的值砍掉，讓舊Omap整個被gc回收...


# refer:
- https://forum.golangbridge.org/t/will-golang-free-the-memory-when-i-use-delete-map-somekey/9081
- https://github.com/patrickmn/go-cache/issues/110
- https://github.com/golang/go/issues/20135
```
The only available workaround is to make a new map and copy in elements from the old.

That is, you have to let the entire map be garbage-collected. Then all its memory will eventually be made available again, and you can start using a new and smaller map. If this doesn't work, please provide a small Go program to reproduce the problem.

As for progress - if there was any, you'd see it in this thread.
```
- https://forum.golangbridge.org/t/will-golang-free-the-memory-when-i-use-delete-map-somekey/9081
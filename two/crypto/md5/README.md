# MD5
Golang的加密庫都放在crypto目錄下，其中MD5庫在crypto/md5包中，該包主要提供了New()和Sum()函數

- md5.New(): 初始化一個MD5對象，返回一個`hash.Hash`函數
- hash.Hash函數原型為 `func New() hash.Hash`
- 透過 hash.Hash 的 Sum 接口計算出 MD5 校驗和。
- Sum(): 函數原型 `func Sum(data []byte) [Size]byte

Sum()並不是對data進行校驗計算，而是對hash.Hash對象內部存儲的內容進行校驗和計算

然後將其追加到data後面形成一個新的byte切片。

因此，通常的方法是將data設置為nil

該方法返回一個 Size 大小為 16的 byte 數組，對於MD5來書就是一個128bit的16字節byte數組


# refer:
- https://www.jianshu.com/p/2639cfa973e8
- https://gist.github.com/sergiotapia/8263278
- https://golangtc.com/t/53c484dc320b525d64000065

### 重要!!shell中使用md5加密需要兜上`-n` => echo -n "hello"|md5
- https://superuser.com/questions/71554/why-is-my-command-line-hash-different-from-online-md5-hash-results

# decrypt-refer:
- https://www.cmd5.com/
(備註：免費的解密網站，只能解開簡單的md5 ... 一般驗證都是透過重新加密一遍後對比結果的...)
# 摘要
請求失敗時，透過定義RetryPolicy來判定哪種情形下要做請求重試。Backoff來定義請求失敗到重試之間的間隔要做什麼
```go
// 宣告一個具有重試請求的http client
type Client struct {
	HTTPClient   *http.Client
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
	RetryMax     int

	CheckForRetry CheckForRetry

	Backoff Backoff
}
```

# 發起請求
golang的http請求具有以下幾個流程
1. 製作請求用封包http.NewRequest("請求方法","請求url","如果是post或put...要帶body[io.Reader]")
2. 透過HTTPClient.Do("先前做好的請求") 來做請求
3. 針對請求回來的Response做解析，通常是拿Response.Body。會透過  ioutil.ReadAll(resp.Body)

```
1. 製作請求http.NewRequest(body)
2. 使用Client進行請求發送
3. 解析Response的字串
```
# 實做Client所需要的請求方法
0. drainBody(body io.ReadCloser)：透過讀取已經送出去的body來重複利用計有的連線
```
defer body.Close()
_, err := io.Copy(ioutil.Discard, io.LimitReader(body,10))
```

1. Do：實做請求的流程
```
1. 每次重試請求前rewind 請求的body
2. 執行請求
3. 針對請求回來的resp以及err套用CheckForRetry來檢驗是否重做
```

2. Get：延用之前請求的Do實做Request在用Do來做請求
```
NewRequest("請求方法", "請求位址", nil)
return c.Do(req)
```

3. Post：延用之前請求的Do實做Request在用Do來做請求
```
NewRequest("請求方法", "請求位址", "請求所帶的內容")
return c.Do(req)
```

# refer:
- https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f
- https://stackoverflow.com/questions/23494950/specifically-check-for-timeout-error
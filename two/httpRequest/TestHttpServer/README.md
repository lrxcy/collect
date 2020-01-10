# intro
製作一個api測試的教學文件

# 定義api
```go
type API struct {
	Client  *http.Client
	baseURL string
}

func (api *API) DoStuff() ([]byte, error) {
	resp, err := api.Client.Get(api.baseURL + "/some/path")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
```

# 測試
### 測試httpServer
```go
func TestDoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), server.URL}
	body, err := api.DoStuff()

	ok(t, err)
	equals(t, []byte("OK"), body)

}
```
### 測試傳輸協定
```go
package api_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDoStuffWithRoundTripper(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		equals(t, req.URL.String(), "http://example.com/some/path")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
 			// Must be set to non-nil value or it panics
			Header:     make(http.Header),
		}
	})

	api := API{client, "http://example.com"}
	body, err := api.DoStuff()
	ok(t, err)
	equals(t, []byte("OK"), body)

}
```

# refer:
- http://hassansin.github.io/Unit-Testing-http-client-in-Go
- https://stackoverflow.com/questions/50081104/how-to-mock-second-try-of-http-call
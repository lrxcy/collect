
# 筆記TestServer

```go
// 筆記TestServer
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpTestServer(t *testing.T) {
	// 宣告一個新的測試Server（參數：handlerFunc() : http.HandlerFunc( func( ... ))
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 可以自行定義header ,這邊不進行定義。
		// w.Header().Set("Content-Type", "application/json")
		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}
		if r.Body != nil {
			// 這邊因為要接收傳進來的body並判斷，所以做json
			jsonResponse := map[string]string{}
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &jsonResponse)

			// 利用傳送進來的body判斷是否符合需求
			if jsonResponse["email"] == "testUser" && jsonResponse["password"] == "testPasswd" {
				// 寫入handler
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("pass"))
			} else {
				fmt.Printf("the email is %v\n ", jsonResponse["email"])
				fmt.Printf("the password is %v\n ", jsonResponse["password"])
			}
		}
	}))
	defer ts.Close()
	data := map[string]string{}
	data["email"] = "testUser"
	data["password"] = "testPasswd"

	payloadBytes, _ := json.Marshal(data)
	body := bytes.NewReader(payloadBytes)

	res, err := http.Post(ts.URL, "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	respbody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "pass", string(respbody))

}
```
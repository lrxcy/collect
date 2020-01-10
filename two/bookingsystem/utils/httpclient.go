package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	. "github.com/jimweng/bookingsystem/logger"
)

// 使用 Singleton 避免client做copy ...
func RetriveHttpClient() *Client {
	// once.Do(func() {
	// 	httpClient = &Client{}
	// })

	return httpClient
}

var (
	httpClient *Client
	once       sync.Once
)

type Client struct {
	HTTPClient        *http.Client
	RetryAfterSeconds retryAfterSeconds
	RetryMax          int

	CheckForRetry checkForRetry
}

/*
   TODO_1: Initialize httpclient from config?
   TODO_2: Add method for specific retry methods
*/
func init() {
	httpClient = newClient()
}

const (
	MaxIdleConns         int = 60000
	MaxIdleConnsPerHost  int = 10000
	IdleConnTimeout      int = 150
	RequestTimeout       int = 30
	RetryMaxTimes        int = 0
	InitRetryWaitSeconds int = 1
)

/*
   有需要帶上特殊header或是特殊處理的httpClient可以定義在這邊，供後面直接做調度
   TODO: 目前先把key hard code 在代碼裡面，之後再看有沒有需要放到設定檔或動態帶入
*/
// Customerize Get method
func (c *Client) Get(url string) (*http.Response, error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// os.Getenv("secretKey") ... 可以儲存在電腦或是環境變數裡面
	req.Header.Set("key", os.Getenv("secretKey"))
	return c.Do(req)
}

// Customerize Post method : only support Content-Type: application/json
func (c *Client) Post(url string, body io.ReadSeeker) (*http.Response, error) {
	req, err := NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// os.Getenv("secretKey") ... 可以儲存在電腦或是環境變數裡面
	req.Header.Set("key", os.Getenv("secretKey"))
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}

type retryAfterSeconds func(int) int

func defaultDetryAfterSeconds(i int) int {
	return 2*i + InitRetryWaitSeconds
}

type checkForRetry func(resp *http.Response, err error) (bool, error)

func defaultRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}

	if resp.StatusCode == 0 || resp.StatusCode >= 500 {
		return true, nil
	}

	return false, nil
}

func newClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(RequestTimeout) * time.Second,
					KeepAlive: time.Duration(RequestTimeout) * time.Second,
				}).DialContext,

				MaxIdleConns:        MaxIdleConns,
				MaxIdleConnsPerHost: MaxIdleConnsPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Duration(RequestTimeout) * time.Second,
		},
		RetryMax:          RetryMaxTimes,
		RetryAfterSeconds: defaultDetryAfterSeconds,
		CheckForRetry:     defaultRetryPolicy,
	}
}

type request struct {
	body io.ReadSeeker
	*http.Request
}

func NewRequest(method, url string, body io.ReadSeeker) (*request, error) {
	var rcBody io.ReadCloser
	if body != nil {
		rcBody = ioutil.NopCloser(body)
	}

	httpReq, err := http.NewRequest(method, url, rcBody)
	if err != nil {
		return nil, err
	}

	return &request{body, httpReq}, nil
}

func (c *Client) Do(req *request) (*http.Response, error) {

	count := 0
	retryAfterSeconds := 0
	for {
		if req.Body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v\n", err)
			}
		}

		resp, err := c.HTTPClient.Do(req.Request)
		if err != nil {
			Log.Errorf(fmt.Sprintf("Err occur while Do request : %v", err))
		}

		checkOk, checkErr := c.CheckForRetry(resp, err)
		if err != nil {
			Log.Error(fmt.Sprintf("Err %s %s request failed: %v\n", req.Method, req.URL, err))
		}

		if !checkOk {
			if checkErr != nil {
				err = checkErr
			}
			return resp, err
		}

		// going to retry, consume any response to reuse the connection
		if err == nil {
			c.drainBody(resp.Body)
		}

		if count != c.RetryMax {
			retryAfterSeconds = defaultDetryAfterSeconds(count)
			count++
		} else {
			break
		}
		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		Log.Info(fmt.Sprintf("DEBUG %s: retrying in %d mins(s) %d left)", desc, retryAfterSeconds, RetryMaxTimes-count))
		time.Sleep(time.Duration(retryAfterSeconds) * time.Second)

	}
	return nil, fmt.Errorf("%s %s giving up after %d attemps", req.Method, req.URL, count)
}

// Try to read the response body so we can reuse this connection
func (c *Client) drainBody(body io.ReadCloser) {
	defer body.Close()

	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, 10))
	if err != nil {
		Log.Error(fmt.Sprintf("Error reading response body: %v\n", err))
	}
}

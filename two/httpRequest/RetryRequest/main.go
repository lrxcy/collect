package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

var (
	// default retry configuration
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
	defaultRetryMax     = 4
)

type CheckForRetry func(resp *http.Response, err error) (bool, error)

// 創建一個RetryPolicy函數，當回應狀態為500時進行重新請求的機制
func DefaultRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}

	if resp.StatusCode == 0 || resp.StatusCode >= 500 {
		return true, nil
	}
	return false, nil
}

// 創建一個Backoff函數用於做請前之間的延遲
type Backoff func(min, max time.Duration, attempNum int, resp *http.Response) time.Duration

func DefaultBackoff(min, max time.Duration, attempNum int, resp *http.Response) time.Duration {
	mult := math.Pow(2, float64(attempNum)*float64(min))
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	return sleep
}

// 宣告一個具有重試請求的http client
type Client struct {
	HTTPClient   *http.Client
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
	RetryMax     int

	CheckForRetry CheckForRetry

	Backoff Backoff
}

func NewClient() *Client {
	return &Client{
		HTTPClient:    http.DefaultClient,
		RetryWaitMin:  defaultRetryWaitMin,
		RetryWaitMax:  defaultRetryWaitMax,
		RetryMax:      defaultRetryMax,
		CheckForRetry: DefaultRetryPolicy,
		Backoff:       DefaultBackoff,
	}
}

type Request struct {
	body io.ReadSeeker

	*http.Request
}

func NewRequest(method, url string, body io.ReadSeeker) (*Request, error) {
	var rcBody io.ReadCloser
	if body != nil {
		rcBody = ioutil.NopCloser(body)
	}

	httpReq, err := http.NewRequest(method, url, rcBody)
	if err != nil {
		return nil, err
	}
	return &Request{body, httpReq}, nil
}

// claim a Do method for Client
func (c *Client) Do(req *Request) (*http.Response, error) {
	fmt.Printf("DEBUG %s %s\n", req.Method, req.URL)
	for {
		var code int // HTTP response code

		// Always rewind the request body when non-nil
		if req.Body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v\n", err)
			}
		}

		// Attempt the request
		resp, err := c.HTTPClient.Do(req.Request)

		// Check if we should continue with retries.
		checkOk, checkErr := c.CheckForRetry(resp, err)

		if err != nil {
			fmt.Printf("ERR %s %s request failed: %v\n", req.Method, req.URL, err)
		} else {
			// Call this here to maintain the behavior of logging all requests, etc
			// even if CheckForRetry signals to stop
		}

		// decide whether to continue
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

		remain := c.RetryMax - 1 // <--- it's a bug
		if remain == 0 {
			break
		}

		wait := c.Backoff(c.RetryWaitMin, c.RetryWaitMax, 1, resp)
		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		if code > 0 {
			desc = fmt.Sprintf("%s (status: %d)", desc, code)
		}
		fmt.Printf("DEBUG %s: retrying in %s (%d left)", desc, wait, remain)
		time.Sleep(wait)
	}
	return nil, fmt.Errorf("%s %s giving up after %d attempts", req.Method, req.URL, c.RetryWaitMax+1)
}

// Try to read the response body so we can reuse this connection
func (c *Client) drainBody(body io.ReadCloser) {
	defer body.Close()

	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, 10))
	if err != nil {
		fmt.Println("Error reading response body: %v", err)
	}
}

// Get method
func (c *Client) Get(url string) (*http.Response, error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// Post method
func (c *Client) Post(url, bodyType string, body io.ReadSeeker) (*http.Response, error) {
	req, err := NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return c.Do(req)
}

func main() {
	retryclient := NewClient()

	resp, err := retryclient.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

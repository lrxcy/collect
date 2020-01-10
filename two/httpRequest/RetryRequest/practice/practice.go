package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func main() {
	client := NewClient()

	// resp, err := client.Get("http://jimqaweb.mlytics.ai/cache.txt")
	resp, err := client.Get("http://107.167.176.135")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}

type RetryAfterMins func(int) int

func DefaultDetryAfterMins(i int) int {
	return 2*i + 1
}

type CheckForRetry func(resp *http.Response, err error) (bool, error)

func DefaultRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}

	if resp.StatusCode == 0 || resp.StatusCode >= 500 {
		return true, nil
	}

	return false, nil
}

type Client struct {
	HTTPClient     *http.Client
	RetryAfterMins RetryAfterMins
	RetryMax       int

	CheckForRetry CheckForRetry
}

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,

				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     150 * time.Second,
			},

			Timeout: 30 * time.Second,
		},
		RetryMax:       3,
		RetryAfterMins: DefaultDetryAfterMins,
		CheckForRetry:  DefaultRetryPolicy,
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

func (c *Client) Do(req *Request) (*http.Response, error) {
	count := 0
	retryAfterMins := 0
	for {
		if req.Body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v\n", err)
			}
		}

		resp, err := c.HTTPClient.Do(req.Request)

		checkOk, checkErr := c.CheckForRetry(resp, err)
		if err != nil {
			log.Printf("Err %s %s request failed: %v\n", req.Method, req.URL, err)
		} else {

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
			retryAfterMins = DefaultDetryAfterMins(count)
			count++
		} else {
			break
		}
		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		log.Printf("DEBUG %s: retrying in %d mins(s) %d left)", desc, retryAfterMins, count)
		time.Sleep(time.Duration(retryAfterMins) * time.Second)

	}
	return nil, fmt.Errorf("%s %s giving up after %d attemps", req.Method, req.URL, count)
}

// Try to read the response body so we can reuse this connection
func (c *Client) drainBody(body io.ReadCloser) {
	defer body.Close()

	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, 10))
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
	}
}

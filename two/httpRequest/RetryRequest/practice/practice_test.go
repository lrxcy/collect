package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoStuffWithTestServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/", req.URL.String())
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	httpClient := &Client{
		HTTPClient:     server.Client(),
		RetryMax:       3,
		RetryAfterMins: DefaultDetryAfterMins,
		CheckForRetry:  DefaultRetryPolicy,
	}

	resp, err := httpClient.Get(server.URL)
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, []byte("OK"), body)

}

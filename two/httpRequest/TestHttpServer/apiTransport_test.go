package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDoStuffWithRoundTripper(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://example.com/some/path")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	api := API{client, "http://example.com"}
	body, err := api.DoStuff()
	assert.Nil(t, err)
	assert.Equal(t, []byte("OK"), body)

}

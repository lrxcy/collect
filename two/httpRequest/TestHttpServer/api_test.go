package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoStuffWithTestServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/some/path", req.URL.String())
		rw.Write([]byte(`OK`))
	}))

	defer server.Close()

	api := API{server.Client(), server.URL}
	body, err := api.DoStuff()
	assert.Nil(t, err)
	assert.Equal(t, []byte("OK"), body)
}

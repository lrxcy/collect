package api

import "net/http"

import "io/ioutil"

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

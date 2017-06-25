package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type (
	// Gitlab contain Auth and BaseURL
	Gitlab struct {
		Host string
	}
)

// NewGitlab is initial Gitlab object
func NewGitlab(host string) *Gitlab {
	url := strings.TrimRight(host, "/")
	return &Gitlab{
		Host: url,
	}
}

func (g *Gitlab) sendRequest(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

func (g *Gitlab) parseResponse(resp *http.Response, body interface{}) (err error) {
	defer resp.Body.Close()

	if body == nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return json.Unmarshal(data, body)
}

func (g *Gitlab) trigger(id string, params url.Values, body interface{}) (err error) {
	requestURL := g.buildURL(id, params)
	req, err := http.NewRequest("POST", requestURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := g.sendRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	return g.parseResponse(resp, body)
}

func (g *Gitlab) buildURL(id string, params url.Values) string {
	url := g.Host + "/api/v4/projects/" + id + "/trigger/pipeline"

	if params != nil {
		queryString := params.Encode()
		if queryString != "" {
			url = url + "?" + queryString
		}
	}

	return url
}

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type (
	// Gitlab contain Auth and BaseURL
	Gitlab struct {
		host   string
		debug  bool
		client *http.Client
	}
)

// NewGitlab is initial Gitlab object
func NewGitlab(host string, insecure, debug bool) *Gitlab {
	url := strings.TrimRight(host, "/")
	g := &Gitlab{
		host:   url,
		debug:  debug,
		client: http.DefaultClient,
	}

	if insecure {
		g.client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	return g
}

func (g *Gitlab) sendRequest(req *http.Request) (*http.Response, error) {
	return g.client.Do(req)
}

func (g *Gitlab) parseResponse(resp *http.Response, body interface{}) (err error) {
	defer resp.Body.Close()

	if body == nil {
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if g.debug {
		fmt.Println()
		fmt.Println("========= Response Body =========")
		fmt.Println(string(data))
		fmt.Println("=================================")
	}

	err = json.Unmarshal(data, body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if g.debug {
		fmt.Println()
		fmt.Println("========= JSON Body ==========")
		fmt.Printf("%+v\n", body)
		fmt.Println("==============================")
	}

	return nil
}

func (g *Gitlab) trigger(id string, params url.Values, body interface{}) (err error) {
	requestURL := g.buildURL(id, nil)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	if err := w.WriteField("token", params.Get("token")); err != nil {
		return err
	}
	if err := w.WriteField("ref", params.Get("ref")); err != nil {
		return err
	}
	// Remove token and ref from params
	params.Del("token")
	params.Del("ref")

	for key := range params {
		if err := w.WriteField(fmt.Sprintf("variables[%s]", key), params.Get(key)); err != nil {
			return err
		}
	}
	w.Close()

	req, err := http.NewRequest("POST", requestURL, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := g.sendRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	return g.parseResponse(resp, body)
}

func (g *Gitlab) buildURL(id string, params url.Values) string {
	url := g.host + "/api/v4/projects/" + id + "/trigger/pipeline"

	if params != nil {
		queryString := params.Encode()
		if queryString != "" {
			url = url + "?" + queryString
		}
	}

	return url
}

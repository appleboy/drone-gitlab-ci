package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	ci := &Gitlab{
		Host: "https://gitlab.com",
	}

	assert.Equal(t, "https://gitlab.com/api/v4/projects/1234/trigger/pipeline", ci.buildURL("1234", nil))
}

func TestHostNotFound(t *testing.T) {
	ci := &Gitlab{
		Host: "https://foo.bar",
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	err := ci.trigger("1234", params, nil)
	assert.NotNil(t, err)
}

func TestNilBody(t *testing.T) {
	ci := &Gitlab{
		Host: "https://gitlab.com",
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	err := ci.trigger("1234", params, nil)
	assert.Nil(t, err)
}

func TestResponse404Body(t *testing.T) {
	type body struct {
		Message string `json:"message"`
	}

	ci := &Gitlab{
		Host:  "https://gitlab.com",
		Debug: true,
	}

	params := url.Values{
		"token": []string{"foo"},
		"ref":   []string{"bar"},
	}

	resp := &body{}
	err := ci.trigger("1234", params, resp)
	assert.Equal(t, "404 Not Found", resp.Message)
	assert.Nil(t, err)
}

func TestTriggerMaster(t *testing.T) {
	type body struct {
		Message string `json:"message"`
	}

	ci := &Gitlab{
		Host:  "https://gitlab.com",
		Debug: true,
	}

	params := url.Values{
		"token": []string{"9184302d980918efad05bce8b97774"},
		"ref":   []string{"master"},
	}

	resp := &body{}
	// https://gitlab.com/appleboy/go-hello
	err := ci.trigger("3573921", params, resp)
	assert.Nil(t, err)
}

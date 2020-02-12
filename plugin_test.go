package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingConfig(t *testing.T) {
	var plugin Plugin

	err := plugin.Exec()

	assert.NotNil(t, err)
}

func TestJSONBodyParseError(t *testing.T) {
	plugin := Plugin{
		Host:  "http://example.com",
		Token: "foo",
		ID:    "bar",
		Ref:   "master",
	}

	err := plugin.Exec()
	assert.NotNil(t, err)
}

func TestGitlabHost(t *testing.T) {
	plugin := Plugin{
		Host:  "https://gitlab.com",
		Token: "testing this gitlab client",
		ID:    "foo",
		Ref:   "bar",
	}

	err := plugin.Exec()
	assert.True(t, strings.Contains(err.Error(), "Unauthorized"))
}

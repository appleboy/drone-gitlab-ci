package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingConfig(t *testing.T) {
	var plugin Plugin

	err := plugin.Exec()

	assert.NotNil(t, err)
}

func TestPluginTriggerBuild(t *testing.T) {
	plugin := Plugin{
		Host:  "http://example.com",
		Token: "foo",
		ID:    "bar",
		Ref:   "master",
	}

	err := plugin.Exec()

	assert.Nil(t, err)
}

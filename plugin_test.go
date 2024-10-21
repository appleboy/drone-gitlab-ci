package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingConfig(t *testing.T) {
	var plugin Plugin

	err := plugin.Exec()

	assert.NotNil(t, err)
}

func TestSetOutput(t *testing.T) {
	t.Run("GITHUB_OUTPUT not set", func(t *testing.T) {
		var plugin Plugin
		err := plugin.SetOutput(map[string]string{"key": "value"})
		assert.NotNil(t, err)
		assert.Equal(t, "GITHUB_OUTPUT is not set", err.Error())
	})

	t.Run("GITHUB_OUTPUT set and file write successful", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "github_output")
		assert.Nil(t, err)
		defer os.Remove(tempFile.Name())

		os.Setenv("GITHUB_OUTPUT", tempFile.Name())
		defer os.Unsetenv("GITHUB_OUTPUT")

		var plugin Plugin
		err = plugin.SetOutput(map[string]string{"key": "value"})
		assert.Nil(t, err)

		content, err := os.ReadFile(tempFile.Name())
		assert.Nil(t, err)
		assert.Contains(t, string(content), "key=value\n")
	})

	t.Run("GITHUB_OUTPUT set but file write fails", func(t *testing.T) {
		os.Setenv("GITHUB_OUTPUT", "/invalid/path")
		defer os.Unsetenv("GITHUB_OUTPUT")

		var plugin Plugin
		err := plugin.SetOutput(map[string]string{"key": "value"})
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "failed to open file")
	})
}

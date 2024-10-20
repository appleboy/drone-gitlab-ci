package main

import (
	"errors"
)

type (
	// Plugin values.
	Plugin struct {
		Host      string
		Token     string
		Ref       string
		ID        string
		Debug     bool
		Variables map[string]string
		Insecure  bool
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	if len(p.Host) == 0 {
		return errors.New("missing host")
	}
	if len(p.Token) == 0 {
		return errors.New("missing token")
	}
	if len(p.ID) == 0 {
		return errors.New("missing project id")
	}

	// Create Gitlab object
	g, err := NewGitlab(p.Host, p.Token, p.Insecure, p.Debug)
	if err != nil {
		return err
	}

	// Create pipeline
	err = g.CreatePipeline(p.ID, p.Ref, p.Variables)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"errors"
	"log"
	"net/url"
)

type (
	// Plugin values.
	Plugin struct {
		Host  string
		Token string
		Ref   string
		ID    string
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {

	if len(p.Host) == 0 || len(p.Token) == 0 || len(p.Ref) == 0 || len(p.ID) == 0 {
		log.Println("missing gitlab-ci config")

		return errors.New("missing gitlab-ci config")
	}

	ci := NewGitlab(p.Host)

	params := url.Values{
		"token": []string{p.Token},
		"ref":   []string{p.Ref},
	}

	return ci.trigger(p.ID, params, nil)
}

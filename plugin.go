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

	// Commit struct
	Commit struct {
		ID     int64  `json:"id"`
		Sha    string `json:"sha"`
		Ref    string `json:"ref"`
		Status string `json:"status"`
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

	body := &Commit{}

	err := ci.trigger(p.ID, params, body)

	if err != nil {
		log.Println("gitlab-ci error:", err.Error())
		return err
	}

	log.Println("build id:", body.ID)
	log.Println("build sha:", body.Sha)
	log.Println("build ref:", body.Ref)
	log.Println("build status:", body.Status)

	return nil
}

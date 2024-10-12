package main

import (
	"errors"
	"log"
	"net/url"
	"time"
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

	// Commit struct
	Commit struct {
		ID         int         `json:"id"`
		Sha        string      `json:"sha"`
		Ref        string      `json:"ref"`
		Status     string      `json:"status"`
		BeforeSha  string      `json:"before_sha"`
		Tag        bool        `json:"tag"`
		YamlErrors interface{} `json:"yaml_errors"`
		User       struct {
			Name      string `json:"name"`
			Username  string `json:"username"`
			ID        int    `json:"id"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		} `json:"user"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		StartedAt   time.Time   `json:"started_at"`
		FinishedAt  time.Time   `json:"finished_at"`
		CommittedAt time.Time   `json:"committed_at"`
		Duration    interface{} `json:"duration"`
		Coverage    interface{} `json:"coverage"`
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	if len(p.Host) == 0 || len(p.Token) == 0 || len(p.ID) == 0 {
		return errors.New("missing gitlab-ci config")
	}

	ci := NewGitlab(p.Host, p.Insecure, p.Debug)

	params := url.Values{
		"token": []string{p.Token},
		"ref":   []string{p.Ref},
	}

	for key, value := range p.Variables {
		params[key] = []string{value}
	}

	body := &Commit{}

	err := ci.trigger(p.ID, params, body)
	if err != nil {
		return err
	}

	log.Println("build id:", body.ID)
	log.Println("build sha:", body.Sha)
	log.Println("build ref:", body.Ref)
	log.Println("build status:", body.Status)

	return nil
}

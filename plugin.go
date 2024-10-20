package main

import (
	"errors"
	"log"
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
		Timeout   time.Duration
		Interval  time.Duration
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
	pipeline, err := g.CreatePipeline(p.ID, p.Ref, p.Variables)
	if err != nil {
		return err
	}

	log.Println("gitlab-ci: Pipeline ID: ", pipeline.ID)
	log.Println("gitlab-ci: Build SHA: ", pipeline.SHA)
	log.Println("gitlab-ci: Build Ref: ", pipeline.Ref)
	log.Println("gitlab-ci: Build Status: ", pipeline.Status)
	log.Println("gitlab-ci: Build WebURL: ", pipeline.WebURL)
	log.Println("gitlab-ci: Build CreatedAt: ", pipeline.CreatedAt)
	log.Println("gitlab-ci: Build UpdatedAt: ", pipeline.UpdatedAt)

	// Wait for pipeline to complete
	ticker := time.NewTicker(p.Interval)
	defer ticker.Stop()

	log.Println("gitlab-ci: Waiting for pipeline to complete...")
	for {
		select {
		case <-time.After(p.Timeout):
			return errors.New("timeout waiting for pipeline to complete after " + p.Timeout.String())
		case <-ticker.C:
			// Check pipeline status
			status, err := g.GetPipelineStatus(p.ID, pipeline.ID)
			if err != nil {
				return err
			}

			log.Println("gitlab-ci: Current pipeline status:", status)
			log.Println("gitlab-ci: Trigger by user:", pipeline.User.Name)

			// https://docs.gitlab.com/ee/api/pipelines.html
			// created, waiting_for_resource, preparing, pending,
			// running, success, failed, canceled, skipped, manual, scheduled
			if status == "success" ||
				status == "failed" ||
				status == "canceled" ||
				status == "skipped" {
				log.Println("gitlab-ci: Pipeline completed with status:", status)
				return nil
			}
		}
	}
}

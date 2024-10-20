package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type (
	// Plugin values.
	Plugin struct {
		Host      string
		Token     string
		Ref       string
		ProjectID string
		Debug     bool
		Variables map[string]string
		Insecure  bool
		Timeout   time.Duration
		Interval  time.Duration
		Wait      bool
		IsGitHub  bool
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	l := slog.New(slog.NewTextHandler(os.Stderr, nil)).
		With("project_id", p.ProjectID)

	if err := p.Validate(); err != nil {
		return err
	}

	// Create Gitlab object
	g, err := NewGitlab(p.Host, p.Token, p.Insecure, p.Debug)
	if err != nil {
		return err
	}

	// Create pipeline
	pipeline, err := g.CreatePipeline(p.ProjectID, p.Ref, p.Variables)
	if err != nil {
		return err
	}

	l.Info(
		"pipeline created",
		"pipeline_id", pipeline.ID,
		"pipeline_sha", pipeline.SHA,
		"pipeline_ref", pipeline.Ref,
		"pipeline_status", pipeline.Status,
		"pipeline_web_url", pipeline.WebURL,
		"pipeline_created_at", pipeline.CreatedAt,
	)

	// Set output
	if p.IsGitHub {
		if err := p.SetOutput(map[string]string{
			"id":      fmt.Sprint(pipeline.ID),
			"sha":     pipeline.SHA,
			"ref":     pipeline.Ref,
			"web_url": pipeline.WebURL,
		}); err != nil {
			return err
		}
	}

	// Wait for pipeline to complete
	if !p.Wait {
		return nil
	}

	// Wait for pipeline to complete
	ticker := time.NewTicker(p.Interval)
	ctxTimout, cancel := context.WithTimeout(context.Background(), p.Timeout)
	defer cancel()
	defer ticker.Stop()

	l.Info("waiting for pipeline to complete", "timeout", p.Timeout)
	for {
		select {
		case <-ctxTimout.Done():
			return errors.New("timeout waiting for pipeline to complete after " + p.Timeout.String())
		case <-ticker.C:
			// Check pipeline status
			status, err := g.GetPipelineStatus(p.ProjectID, pipeline.ID)
			if err != nil {
				return err
			}

			l.Info("pipeline status",
				"status", status,
				"triggered_by", pipeline.User.Name,
			)

			// https://docs.gitlab.com/ee/api/pipelines.html
			// created, waiting_for_resource, preparing, pending,
			// running, success, failed, canceled, skipped, manual, scheduled
			if status == "success" ||
				status == "failed" ||
				status == "canceled" ||
				status == "skipped" {
				l.Info("pipeline completed", "status", status)
				if p.IsGitHub {
					// update status
					if err := p.SetOutput(map[string]string{"status": status}); err != nil {
						return err
					}
				}
				return nil
			}

			if ctxTimout.Err() != nil {
				if p.IsGitHub {
					// update status
					if err := p.SetOutput(map[string]string{"status": status}); err != nil {
						return err
					}
				}
				return ctxTimout.Err()
			}
		}
	}
}

func (p Plugin) SetOutput(data map[string]string) error {
	githubOutput := os.Getenv("GITHUB_OUTPUT")
	if githubOutput == "" {
		return errors.New("GITHUB_OUTPUT is not set")
	}

	file, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", githubOutput, err)
	}
	defer file.Close()

	for k, v := range data {
		_, err = fmt.Fprintf(file, "%s=%s\n", k, v)
		if err != nil {
			return fmt.Errorf("failed to write to file %s: %w", githubOutput, err)
		}
	}

	return nil
}

// Validate checks the plugin configuration.
func (p Plugin) Validate() error {
	if len(p.Host) == 0 {
		return errors.New("missing host")
	}
	if len(p.Token) == 0 {
		return errors.New("missing token")
	}
	if len(p.ProjectID) == 0 {
		return errors.New("missing project id")
	}
	return nil
}

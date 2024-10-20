package main

import (
	"errors"
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
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	l := slog.New(slog.NewTextHandler(os.Stderr, nil)).
		With("project_id", p.ProjectID)

	if len(p.Host) == 0 {
		return errors.New("missing host")
	}
	if len(p.Token) == 0 {
		return errors.New("missing token")
	}
	if len(p.ProjectID) == 0 {
		return errors.New("missing project id")
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

	// Wait for pipeline to complete
	if !p.Wait {
		return nil
	}

	// Wait for pipeline to complete
	ticker := time.NewTicker(p.Interval)
	defer ticker.Stop()

	l.Info("waiting for pipeline to complete", "timeout", p.Timeout)
	for {
		select {
		case <-time.After(p.Timeout):
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
				return nil
			}
		}
	}
}

package main

import (
	"crypto/tls"
	"net/http"

	"github.com/xanzy/go-gitlab"
)

type (
	// Gitlab contain Auth and BaseURL
	Gitlab struct {
		client *gitlab.Client
	}
)

// NewGitlab is initial Gitlab object
func NewGitlab(host, token string, insecure, debug bool) (*Gitlab, error) {
	httpClient := http.DefaultClient
	if insecure {
		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	g, err := gitlab.NewClient(
		token,
		gitlab.WithBaseURL(host),
		gitlab.WithHTTPClient(httpClient),
	)
	if err != nil {
		return nil, err
	}

	return &Gitlab{
		client: g,
	}, nil
}

// CreatePipeline triggers the creation of a new pipeline in a specified GitLab project.
//
// Parameters:
//   - projectID: The ID of the GitLab project where the pipeline will be created.
//   - ref: The branch or tag name to create the pipeline for.
//   - variables: A map of environment variables to set for the pipeline.
//
// Returns:
//   - *gitlab.Pipeline: The created pipeline object.
//   - error: An error object if the pipeline creation fails, otherwise nil.
func (g *Gitlab) CreatePipeline(projectID string, ref string, variables map[string]string) (*gitlab.Pipeline, error) {
	allenvs := make([]*gitlab.PipelineVariableOptions, 0)
	options := &gitlab.CreatePipelineOptions{
		Ref:       &ref,
		Variables: &allenvs,
	}
	for k, v := range variables {
		allenvs = append(allenvs, &gitlab.PipelineVariableOptions{
			Key:   &k,
			Value: &v,
		})
	}
	pipeline, _, err := g.client.Pipelines.CreatePipeline(projectID, options)
	if err != nil {
		return nil, err
	}

	return pipeline, nil
}

func (g *Gitlab) GetPipelineStatus(projectID string, pipelineID int) (string, error) {
	pipeline, _, err := g.client.Pipelines.GetPipeline(projectID, pipelineID)
	if err != nil {
		return "", err
	}

	return pipeline.Status, nil
}

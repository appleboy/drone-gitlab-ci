/*
Package main provides functionality to interact with GitLab's API, specifically for creating and managing pipelines.

The Gitlab struct contains an authenticated client to communicate with the GitLab API.

Functions:

- NewGitlab: Initializes a new Gitlab client with the provided host, token, and configuration options.
- CreatePipeline: Triggers the creation of a new pipeline in a specified GitLab project.
- GetPipelineStatus: Retrieves the status of a specific pipeline for a given project.

Types:

- Gitlab: Contains the authenticated client for GitLab API interactions.

Dependencies:

- "crypto/tls"
- "net/http"
- "gitlab.com/gitlab-org/api/client-go"
*/
package main

import (
	"crypto/tls"
	"net/http"

	"github.com/appleboy/com/convert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type (
	// Gitlab contain Auth and BaseURL
	Gitlab struct {
		client *gitlab.Client
	}
)

// NewGitlab initializes a new Gitlab client with the provided host, token, and configuration options.
func NewGitlab(host, token string, insecure bool) (*Gitlab, error) {
	httpClient := http.DefaultClient
	if insecure {
		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // #nosec G402
				},
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
func (g *Gitlab) CreatePipeline(
	projectID, ref string,
	variables map[string]string,
) (*gitlab.Pipeline, error) {
	allenvs := make([]*gitlab.PipelineVariableOptions, 0)
	options := &gitlab.CreatePipelineOptions{
		Ref:       convert.ToPtr(ref),
		Variables: convert.ToPtr(allenvs),
	}
	for k, v := range variables {
		// Usage of single iteration variable in range loop
		key, value := k, v
		allenvs = append(allenvs, &gitlab.PipelineVariableOptions{
			Key:   convert.ToPtr(key),
			Value: convert.ToPtr(value),
		})
	}
	pipeline, _, err := g.client.Pipelines.CreatePipeline(projectID, options)
	if err != nil {
		return nil, err
	}

	return pipeline, nil
}

// GetPipelineStatus retrieves the status of a specific pipeline for a given project.
// It takes the project ID as a string and the pipeline ID as an int64.
// It returns the status of the pipeline as a string and an error if any occurs during the retrieval process.
func (g *Gitlab) GetPipelineStatus(projectID string, pipelineID int64) (string, error) {
	pipeline, _, err := g.client.Pipelines.GetPipeline(projectID, pipelineID)
	if err != nil {
		return "", err
	}

	return pipeline.Status, nil
}

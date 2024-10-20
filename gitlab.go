package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

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

func (g *Gitlab) CreatePipeline(projectID string, ref string, variables map[string]string) error {
	allenvs := make([]*gitlab.PipelineVariableOptions, 0)
	options := &gitlab.CreatePipelineOptions{
		Ref:       &ref,
		Variables: &allenvs,
	}
	for _, variable := range variables {
		kvPair := strings.Split(variable, "=")
		if len(kvPair) != 2 {
			log.Println("gitlab-ci error: invalid environment variable: ", variable)
			continue
		}
		allenvs = append(allenvs, &gitlab.PipelineVariableOptions{
			Key:   &kvPair[0],
			Value: &kvPair[1],
		})
	}
	pipeline, _, err := g.client.Pipelines.CreatePipeline(projectID, options)
	if err != nil {
		return err
	}

	log.Println("gitlab-ci: pipeline ID: ", pipeline.ID)
	log.Println("gitlab-ci: Build SHA: ", pipeline.SHA)
	log.Println("gitlab-ci: Build Ref: ", pipeline.Ref)
	log.Println("gitlab-ci: Build Status: ", pipeline.Status)
	log.Println("gitlab-ci: Build WebURL: ", pipeline.WebURL)
	log.Println("gitlab-ci: Build CreatedAt: ", pipeline.CreatedAt)
	log.Println("gitlab-ci: Build UpdatedAt: ", pipeline.UpdatedAt)
	log.Println("gitlab-ci: Build StartedAt: ", pipeline.StartedAt)
	log.Println("gitlab-ci: Build FinishedAt: ", pipeline.FinishedAt)
	log.Println("gitlab-ci: Build CommittedAt: ", pipeline.CommittedAt)
	log.Println("gitlab-ci: Build Duration: ", pipeline.Duration)

	return nil
}

package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	gitlab "github.com/xanzy/go-gitlab"
)

type (
	// Plugin values.
	Plugin struct {
		Host            string
		Token           string
		Ref             string
		ID              string
		Debug           bool
		Environment     []string
		WaitOnCompleted bool
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {

	if len(p.Host) == 0 || len(p.Token) == 0 || len(p.ID) == 0 {
		return errors.New("missing gitlab-ci config")
	}

	git := gitlab.NewClient(nil, p.Token)
	err := git.SetBaseURL(fmt.Sprintf("%s/api/v4", p.Host))
	if err != nil {
		log.Println("failed setting base url: ", err.Error())
		return err
	}

	options := &gitlab.CreatePipelineOptions{
		Ref:       &p.Ref,
		Variables: make([]*gitlab.PipelineVariable, 0),
	}
	for _, variable := range p.Environment {
		kvPair := strings.Split(variable, "=")
		options.Variables = append(options.Variables, &gitlab.PipelineVariable{
			Key:   kvPair[0],
			Value: kvPair[1],
		})
	}
	pipeline, _, err := git.Pipelines.CreatePipeline(p.ID, options)
	if err != nil {
		log.Println("gitlab-ci error: ", err.Error())
		return err
	}

	log.Println("build id:", pipeline.ID)
	log.Println("build sha:", pipeline.SHA)
	log.Println("build ref:", pipeline.Ref)
	log.Println("build status:", pipeline.Status)

	if p.WaitOnCompleted {
		// sit and watch the pipeline finish
		for {
			pipeline, _, err = git.Pipelines.GetPipeline(p.ID, pipeline.ID)
			if err != nil {
				log.Println("gitlab-ci error: ", err.Error())
				return err
			}
			switch pipeline.Status {
			case "success":
				return nil
			case "failed", "canceled", "skipped":
				return fmt.Errorf("gitlab-ci pipeline status: %s which is not a success, object: %#v", pipeline.Status, *pipeline)
			case "pending", "running":
				time.Sleep(30 * time.Second)
			}

		}
	}
	return nil
}

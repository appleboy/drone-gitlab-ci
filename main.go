package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// Version set at compile-time
var Version string

func main() {
	// Load env-file if it exists first
	if filename, found := os.LookupEnv("PLUGIN_ENV_FILE"); found {
		_ = godotenv.Load(filename)
	}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	copyright := fmt.Sprintf("Copyright (c) %v Bo-Yi Wu", time.Now().Year())
	app := &cli.App{
		Name:      "gitlab-ci plugin",
		Usage:     "trigger gitlab-ci jobs",
		Copyright: copyright,
		Authors: []*cli.Author{
			{
				Name:  "Bo-Yi Wu",
				Email: "appleboy.tw@gmail.com",
			},
		},
		Action:  run,
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"l"},
				Usage:   "gitlab-ci base url",
				EnvVars: []string{"PLUGIN_HOST", "GITLAB_HOST", "INPUT_HOST"},
				Value:   "https://gitlab.com",
			},
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"t"},
				Usage:   "gitlab-ci token",
				EnvVars: []string{"PLUGIN_TOKEN", "GITLAB_TOKEN", "INPUT_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
				Usage:   "gitlab-ci project id",
				EnvVars: []string{"PLUGIN_ID", "GITLAB_PROJECT_ID", "INPUT_PROJECT_ID"},
			},
			&cli.StringFlag{
				Name:    "ref",
				Aliases: []string{"r"},
				Usage:   "gitlab-ci valid refs are only the branches and tags",
				EnvVars: []string{"PLUGIN_REF", "GITLAB_REF", "INPUT_REF"},
				Value:   "main",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "debug mode",
				EnvVars: []string{"PLUGIN_DEBUG", "GITLAB_DEBUG", "INPUT_DEBUG"},
			},
			&cli.StringSliceFlag{
				Name:    "variables",
				Usage:   "gitlab-ci variables",
				EnvVars: []string{"PLUGIN_VARIABLES", "GITLAB_VARIABLES", "INPUT_VARIABLES"},
			},
			&cli.BoolFlag{
				Name:    "insecure",
				Usage:   "allow connections to SSL sites without certs",
				EnvVars: []string{"PLUGIN_INSECURE", "GITLAB_INSECURE", "INPUT_INSECURE"},
			},
			&cli.DurationFlag{
				Name:    "timeout",
				Usage:   "timeout waiting for pipeline to complete",
				EnvVars: []string{"PLUGIN_TIMEOUT", "GITLAB_TIMEOUT", "INPUT_TIMEOUT"},
				Value:   time.Minute * 60,
			},
		},
	}

	// Override a template
	cli.AppHelpTemplate = `
________                                         ________.__  __  .__        ___.             _________ .___
\______ \_______  ____   ____   ____            /  _____/|__|/  |_|  | _____ \_ |__           \_   ___ \|   |
 |    |  \_  __ \/  _ \ /    \_/ __ \   ______ /   \  ___|  \   __\  | \__  \ | __ \   ______ /    \  \/|   |
 |    |   \  | \(  <_> )   |  \  ___/  /_____/ \    \_\  \  ||  | |  |__/ __ \| \_\ \ /_____/ \     \___|   |
/_______  /__|   \____/|___|  /\___  >          \______  /__||__| |____(____  /___  /          \______  /___|
        \/                  \/     \/                  \/                   \/    \/                  \/

                                                                                    version: {{.Version}}
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
REPOSITORY:
    Github: https://github.com/appleboy/drone-line
`

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	variables := make(map[string]string)

	for _, v := range c.StringSlice("variables") {
		s := strings.Split(v, "=")
		if len(s) == 2 {
			variables[s[0]] = s[1]
		}
	}

	plugin := Plugin{
		Host:      c.String("host"),
		Token:     c.String("token"),
		Ref:       c.String("ref"),
		ID:        c.String("id"),
		Debug:     c.Bool("debug"),
		Variables: variables,
		Insecure:  c.Bool("insecure"),
		Timeout:   c.Duration("timeout"),
	}

	return plugin.Exec()
}

package main

import (
	"log"
	"os"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version set at compile-time
var Version string

func main() {
	copyright := fmt.Sprintf("Copyright (c) %v Bo-Yi Wu", time.Now().Year())
	app := cli.NewApp()
	app.Name = "gitlab-ci plugin"
	app.Usage = "trigger gitlab-ci jobs"
	app.Copyright = copyright
	app.Authors = []cli.Author{
		{
			Name:  "Bo-Yi Wu",
			Email: "appleboy.tw@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host,l",
			Usage:  "gitlab-ci base url",
			EnvVar: "PLUGIN_HOST,GITLBA_HOST",
			Value:  "https://gitlab.com",
		},
		cli.StringFlag{
			Name:   "token,t",
			Usage:  "gitlab-ci token",
			EnvVar: "PLUGIN_TOKEN,GITLBA_TOKEN",
		},
		cli.StringFlag{
			Name:   "id,i",
			Usage:  "gitlab-ci project id",
			EnvVar: "PLUGIN_ID,GITLBA_PROJECT_ID",
		},
		cli.StringFlag{
			Name:   "ref,r",
			Usage:  "gitlab-ci valid refs are only the branches and tags",
			EnvVar: "PLUGIN_REF,GITLBA_REF",
			Value:  "master",
		},
		cli.BoolFlag{
			Name:   "debug,d",
			Usage:  "debug mode",
			EnvVar: "PLUGIN_DEBUG,GITLBA_DEBUG",
		},
		cli.StringFlag{
			Name:   "env-file",
			Usage:  "source env file",
			EnvVar: "ENV_FILE",
			Value:  ".env",
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
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Host:  c.String("host"),
		Token: c.String("token"),
		Ref:   c.String("ref"),
		ID:    c.String("id"),
		Debug: c.Bool("debug"),
	}

	return plugin.Exec()
}

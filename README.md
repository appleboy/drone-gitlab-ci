# drone-gitlab-ci

![logo](./images/logo.png)

[繁體中文](./README.zh-tw.md) | [简体中文](./README.zh-cn.md) | [English](./README.md)

[![Lint and Testing](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci)
[![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci)

## What is drone-gitlab-ci?

**drone-gitlab-ci** is a CLI tool and CI/CD plugin that triggers [GitLab CI](https://about.gitlab.com/solutions/continuous-integration/) pipelines from any environment. It serves as a bridge between different CI/CD platforms, enabling you to integrate GitLab CI into your existing workflows seamlessly.

## Why use this?

In modern software development, teams often use multiple CI/CD platforms across different projects or services. This creates challenges when you need to:

- **Cross-Platform Integration**: Trigger GitLab CI pipelines from other CI/CD systems (Drone, GitHub Actions, Jenkins, etc.)
- **Microservices Orchestration**: Automatically trigger dependent service pipelines after deployment
- **Unified Pipeline Management**: Centralize pipeline triggering logic across diverse infrastructure
- **Hybrid Cloud Workflows**: Connect CI/CD processes across different platforms and environments

```txt
┌─────────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   Drone CI      │     │                  │     │                 │
│   GitHub Actions│────▶│ drone-gitlab-ci  │────▶│   GitLab CI     │
│   Jenkins       │     │                  │     │   Pipeline      │
│   Local Dev     │     └──────────────────┘     └─────────────────┘
└─────────────────┘
```

## Use Cases

| Scenario                       | Description                                                          |
| ------------------------------ | -------------------------------------------------------------------- |
| **Drone CI → GitLab CI**       | Use as a Drone plugin to trigger GitLab pipelines after Drone builds |
| **GitHub Actions → GitLab CI** | Trigger GitLab deployments from GitHub repositories                  |
| **Jenkins → GitLab CI**        | Integrate GitLab CI into existing Jenkins workflows                  |
| **Local Development**          | Manually trigger pipelines during development and testing            |
| **Microservices**              | Service A deployment triggers Service B pipeline automatically       |

## Features

- Trigger GitLab CI pipelines via API
- Pass custom variables to pipelines
- Wait for pipeline completion with configurable timeout
- Support for self-hosted GitLab instances
- Cross-platform binary (Windows, Linux, macOS)
- Docker image available
- Native support for Drone CI, GitHub Actions, and other CI/CD platforms

## Table of Contents

- [drone-gitlab-ci](#drone-gitlab-ci)
  - [What is drone-gitlab-ci?](#what-is-drone-gitlab-ci)
  - [Why use this?](#why-use-this)
  - [Use Cases](#use-cases)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Quick Start](#quick-start)
  - [Installation](#installation)
    - [Download Pre-built Binaries](#download-pre-built-binaries)
    - [Install with Go](#install-with-go)
    - [Build from Source](#build-from-source)
  - [Configuration](#configuration)
    - [GitLab Setup](#gitlab-setup)
      - [1. Create a Personal Access Token](#1-create-a-personal-access-token)
      - [2. Get Your Project ID](#2-get-your-project-id)
    - [Parameters](#parameters)
  - [Usage](#usage)
    - [Command Line](#command-line)
    - [Docker](#docker)
    - [Drone CI](#drone-ci)
    - [GitHub Actions](#github-actions)
  - [Development](#development)
    - [Run Tests](#run-tests)
    - [Build Binary](#build-binary)
    - [Code Quality](#code-quality)
  - [License](#license)

## Quick Start

Trigger a GitLab CI pipeline in 3 steps:

```bash
# 1. Download the binary (or use Docker)
go install github.com/appleboy/drone-gitlab-ci@latest

# 2. Set your credentials
export GITLAB_TOKEN=your-gitlab-token
export GITLAB_PROJECT_ID=your-project-id

# 3. Trigger the pipeline
drone-gitlab-ci --ref main
```

## Installation

### Download Pre-built Binaries

Download from the [release page](https://github.com/appleboy/drone-gitlab-ci/releases). Supported platforms:

- Windows (amd64/386)
- Linux (amd64/386)
- macOS (amd64/arm64)

### Install with Go

```bash
go install github.com/appleboy/drone-gitlab-ci@latest
```

### Build from Source

```bash
git clone https://github.com/appleboy/drone-gitlab-ci.git
cd drone-gitlab-ci
make build
```

## Configuration

### GitLab Setup

#### 1. Create a Personal Access Token

You need a GitLab token to authenticate API requests. See [GitLab Token Overview](https://docs.gitlab.com/ee/security/tokens/index.html#personal-access-tokens).

1. Go to GitLab → User Settings → Access Tokens
2. Create a token with `api` scope
3. Save the token securely

![token](./images/user_token.png)

#### 2. Get Your Project ID

Find your project ID in **Settings → General → General project settings**.

![projectID](./images/projectID.png)

For more details, see [Pipeline trigger tokens API](https://docs.gitlab.com/ee/api/pipeline_triggers.html).

### Parameters

| Parameter  | Flag                 | Environment Variables                                     | Description                   | Required | Default              |
| ---------- | -------------------- | --------------------------------------------------------- | ----------------------------- | -------- | -------------------- |
| Host       | `--host`             | `GITLAB_HOST`, `PLUGIN_HOST`, `INPUT_HOST`                | GitLab instance URL           | No       | `https://gitlab.com` |
| Token      | `--token`, `-t`      | `GITLAB_TOKEN`, `PLUGIN_TOKEN`, `INPUT_TOKEN`             | GitLab access token           | **Yes**  | -                    |
| Project ID | `--project-id`, `-p` | `GITLAB_PROJECT_ID`, `PLUGIN_ID`, `INPUT_PROJECT_ID`      | GitLab project ID             | **Yes**  | -                    |
| Ref        | `--ref`, `-r`        | `GITLAB_REF`, `PLUGIN_REF`, `INPUT_REF`                   | Branch or tag to trigger      | No       | `main`               |
| Variables  | `--variables`        | `GITLAB_VARIABLES`, `PLUGIN_VARIABLES`, `INPUT_VARIABLES` | Variables to pass (KEY=VALUE) | No       | -                    |
| Wait       | `--wait`, `-w`       | `GITLAB_WAIT`, `PLUGIN_WAIT`, `INPUT_WAIT`                | Wait for pipeline completion  | No       | `false`              |
| Timeout    | `--timeout`          | `GITLAB_TIMEOUT`, `PLUGIN_TIMEOUT`, `INPUT_TIMEOUT`       | Timeout for waiting           | No       | `60m`                |
| Interval   | `--interval`, `-i`   | `GITLAB_INTERVAL`, `PLUGIN_INTERVAL`, `INPUT_INTERVAL`    | Polling interval              | No       | `5s`                 |
| Insecure   | `--insecure`, `-k`   | `GITLAB_INSECURE`, `PLUGIN_INSECURE`, `INPUT_INSECURE`    | Skip SSL verification         | No       | `false`              |
| Debug      | `--debug`, `-d`      | `GITLAB_DEBUG`, `PLUGIN_DEBUG`, `INPUT_DEBUG`             | Enable debug output           | No       | `false`              |
| GitHub     | `--github`           | `GITHUB_ACTIONS`, `PLUGIN_GITHUB`, `INPUT_GITHUB`         | Enable GitHub Actions output  | No       | `false`              |

## Usage

### Command Line

Basic usage:

```bash
drone-gitlab-ci \
  --host https://gitlab.com \
  --token your-token \
  --project-id 12345 \
  --ref main
```

With variables and wait for completion:

```bash
drone-gitlab-ci \
  --host https://gitlab.com \
  --token your-token \
  --project-id 12345 \
  --ref main \
  --variables "DEPLOY_ENV=production" \
  --variables "VERSION=1.0.0" \
  --wait \
  --timeout 30m
```

Enable debug mode:

```bash
drone-gitlab-ci \
  --host https://gitlab.com \
  --token your-token \
  --project-id 12345 \
  --ref main \
  --debug
```

### Docker

Basic usage:

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com \
  -e GITLAB_TOKEN=your-token \
  -e GITLAB_PROJECT_ID=12345 \
  -e GITLAB_REF=main \
  appleboy/drone-gitlab-ci
```

With wait and debug:

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com \
  -e GITLAB_TOKEN=your-token \
  -e GITLAB_PROJECT_ID=12345 \
  -e GITLAB_REF=main \
  -e GITLAB_WAIT=true \
  -e GITLAB_DEBUG=true \
  appleboy/drone-gitlab-ci
```

### Drone CI

Add to your `.drone.yml`:

```yaml
kind: pipeline
name: default

steps:
  - name: trigger-gitlab
    image: appleboy/drone-gitlab-ci
    settings:
      host: https://gitlab.com
      token:
        from_secret: gitlab_token
      project_id: 12345
      ref: main
      wait: true
      timeout: 30m
      variables:
        - DEPLOY_ENV=production
        - VERSION=${DRONE_TAG}
```

### GitHub Actions

Create `.github/workflows/trigger-gitlab.yml`:

```yaml
name: Trigger GitLab CI

on:
  push:
    branches: [main]

jobs:
  trigger:
    runs-on: ubuntu-latest
    steps:
      - name: Trigger GitLab Pipeline
        uses: docker://appleboy/drone-gitlab-ci
        with:
          host: https://gitlab.com
          token: ${{ secrets.GITLAB_TOKEN }}
          project_id: 12345
          ref: main
          github: true
          wait: true
```

## Development

### Run Tests

```bash
make test
```

### Build Binary

```bash
make build
```

### Code Quality

```bash
make lint
```

## License

MIT License - see [LICENSE](LICENSE) file for details.

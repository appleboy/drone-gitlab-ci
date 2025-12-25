# drone-gitlab-ci

![logo](./images/logo.png)

[繁體中文](./README.zh-tw.md) | [简体中文](./README.zh-cn.md) | [English](./README.md)

[![Lint and Testing](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci)
[![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci)

## 什么是 drone-gitlab-ci？

**drone-gitlab-ci** 是一个 CLI 工具和 CI/CD 插件，可以从任何环境触发 [GitLab CI](https://about.gitlab.com/solutions/continuous-integration/) Pipeline。它作为不同 CI/CD 平台之间的桥梁，让您能够将 GitLab CI 无缝集成到现有的工作流程中。

## 为什么使用这个工具？

在现代软件开发中，团队经常在不同的项目或服务中使用多个 CI/CD 平台。这会在以下情况产生挑战：

- **跨平台集成**：从其他 CI/CD 系统（Drone、GitHub Actions、Jenkins 等）触发 GitLab CI Pipeline
- **微服务编排**：部署后自动触发依赖服务的 Pipeline
- **统一 Pipeline 管理**：在多样化的基础架构中集中管理 Pipeline 触发逻辑
- **混合云工作流程**：连接不同平台和环境之间的 CI/CD 流程

```txt
┌─────────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   Drone CI      │     │                  │     │                 │
│   GitHub Actions│────▶│ drone-gitlab-ci  │────▶│   GitLab CI     │
│   Jenkins       │     │                  │     │   Pipeline      │
│   本地开发       │     └──────────────────┘     └─────────────────┘
└─────────────────┘
```

## 使用场景

| 场景                           | 说明                                                 |
| ------------------------------ | ---------------------------------------------------- |
| **Drone CI → GitLab CI**       | 作为 Drone 插件，在 Drone 构建后触发 GitLab Pipeline |
| **GitHub Actions → GitLab CI** | 从 GitHub 仓库触发 GitLab 部署                       |
| **Jenkins → GitLab CI**        | 将 GitLab CI 集成到现有的 Jenkins 工作流程           |
| **本地开发**                   | 在开发和测试期间手动触发 Pipeline                    |
| **微服务**                     | 服务 A 部署后自动触发服务 B 的 Pipeline              |

## 功能特点

- 通过 API 触发 GitLab CI Pipeline
- 传递自定义变量到 Pipeline
- 等待 Pipeline 完成，可配置超时时间
- 支持自托管 GitLab 实例
- 跨平台可执行文件（Windows、Linux、macOS）
- 提供 Docker 镜像
- 原生支持 Drone CI、GitHub Actions 和其他 CI/CD 平台

## 目录

- [drone-gitlab-ci](#drone-gitlab-ci)
  - [什么是 drone-gitlab-ci？](#什么是-drone-gitlab-ci)
  - [为什么使用这个工具？](#为什么使用这个工具)
  - [使用场景](#使用场景)
  - [功能特点](#功能特点)
  - [目录](#目录)
  - [快速开始](#快速开始)
  - [安装](#安装)
    - [下载预编译可执行文件](#下载预编译可执行文件)
    - [使用 Go 安装](#使用-go-安装)
    - [从源码编译](#从源码编译)
  - [配置](#配置)
    - [GitLab 设置](#gitlab-设置)
      - [1. 创建个人访问令牌](#1-创建个人访问令牌)
      - [2. 获取项目 ID](#2-获取项目-id)
    - [参数说明](#参数说明)
  - [使用方式](#使用方式)
    - [命令行](#命令行)
    - [Docker](#docker)
    - [Drone CI](#drone-ci)
    - [GitHub Actions](#github-actions)
  - [开发](#开发)
    - [运行测试](#运行测试)
    - [编译可执行文件](#编译可执行文件)
    - [代码质量](#代码质量)
  - [许可证](#许可证)

## 快速开始

三个步骤触发 GitLab CI Pipeline：

```bash
# 1. 下载可执行文件（或使用 Docker）
go install github.com/appleboy/drone-gitlab-ci@latest

# 2. 设置认证信息
export GITLAB_TOKEN=your-gitlab-token
export GITLAB_PROJECT_ID=your-project-id

# 3. 触发 Pipeline
drone-gitlab-ci --ref main
```

## 安装

### 下载预编译可执行文件

从 [发布页面](https://github.com/appleboy/drone-gitlab-ci/releases) 下载。支持的平台：

- Windows (amd64/386)
- Linux (amd64/386)
- macOS (amd64/arm64)

### 使用 Go 安装

```bash
go install github.com/appleboy/drone-gitlab-ci@latest
```

### 从源码编译

```bash
git clone https://github.com/appleboy/drone-gitlab-ci.git
cd drone-gitlab-ci
make build
```

## 配置

### GitLab 设置

#### 1. 创建个人访问令牌

您需要 GitLab 令牌来验证 API 请求。请参阅 [GitLab 令牌概述](https://docs.gitlab.com/ee/security/tokens/index.html#personal-access-tokens)。

1. 前往 GitLab → 用户设置 → 访问令牌
2. 创建具有 `api` 范围的令牌
3. 安全保存令牌

![token](./images/user_token.png)

#### 2. 获取项目 ID

在 **设置 → 通用 → 通用项目设置** 中找到您的项目 ID。

![projectID](./images/projectID.png)

更多详情请参阅 [Pipeline trigger tokens API](https://docs.gitlab.com/ee/api/pipeline_triggers.html)。

### 参数说明

| 参数       | 标志                 | 环境变量                                                  | 说明                     | 必填   | 默认值               |
| ---------- | -------------------- | --------------------------------------------------------- | ------------------------ | ------ | -------------------- |
| Host       | `--host`             | `GITLAB_HOST`, `PLUGIN_HOST`, `INPUT_HOST`                | GitLab 实例 URL          | 否     | `https://gitlab.com` |
| Token      | `--token`, `-t`      | `GITLAB_TOKEN`, `PLUGIN_TOKEN`, `INPUT_TOKEN`             | GitLab 访问令牌          | **是** | -                    |
| Project ID | `--project-id`, `-p` | `GITLAB_PROJECT_ID`, `PLUGIN_ID`, `INPUT_PROJECT_ID`      | GitLab 项目 ID           | **是** | -                    |
| Ref        | `--ref`, `-r`        | `GITLAB_REF`, `PLUGIN_REF`, `INPUT_REF`                   | 要触发的分支或标签       | 否     | `main`               |
| Variables  | `--variables`        | `GITLAB_VARIABLES`, `PLUGIN_VARIABLES`, `INPUT_VARIABLES` | 传递的变量 (KEY=VALUE)   | 否     | -                    |
| Wait       | `--wait`, `-w`       | `GITLAB_WAIT`, `PLUGIN_WAIT`, `INPUT_WAIT`                | 等待 Pipeline 完成       | 否     | `false`              |
| Timeout    | `--timeout`          | `GITLAB_TIMEOUT`, `PLUGIN_TIMEOUT`, `INPUT_TIMEOUT`       | 等待超时时间             | 否     | `60m`                |
| Interval   | `--interval`, `-i`   | `GITLAB_INTERVAL`, `PLUGIN_INTERVAL`, `INPUT_INTERVAL`    | 轮询间隔                 | 否     | `5s`                 |
| Insecure   | `--insecure`, `-k`   | `GITLAB_INSECURE`, `PLUGIN_INSECURE`, `INPUT_INSECURE`    | 跳过 SSL 验证            | 否     | `false`              |
| Debug      | `--debug`, `-d`      | `GITLAB_DEBUG`, `PLUGIN_DEBUG`, `INPUT_DEBUG`             | 启用调试输出             | 否     | `false`              |
| GitHub     | `--github`           | `GITHUB_ACTIONS`, `PLUGIN_GITHUB`, `INPUT_GITHUB`         | 启用 GitHub Actions 输出 | 否     | `false`              |

## 使用方式

### 命令行

基本用法：

```bash
drone-gitlab-ci \
  --host https://gitlab.com \
  --token your-token \
  --project-id 12345 \
  --ref main
```

带变量并等待完成：

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

启用调试模式：

```bash
drone-gitlab-ci \
  --host https://gitlab.com \
  --token your-token \
  --project-id 12345 \
  --ref main \
  --debug
```

### Docker

基本用法：

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com \
  -e GITLAB_TOKEN=your-token \
  -e GITLAB_PROJECT_ID=12345 \
  -e GITLAB_REF=main \
  appleboy/drone-gitlab-ci
```

启用等待和调试：

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

添加到您的 `.drone.yml`：

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

创建 `.github/workflows/trigger-gitlab.yml`：

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

## 开发

### 运行测试

```bash
make test
```

### 编译可执行文件

```bash
make build
```

### 代码质量

```bash
make lint
```

## 许可证

MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

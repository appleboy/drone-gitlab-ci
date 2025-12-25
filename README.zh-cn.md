# drone-gitlab-ci

![logo](./images/logo.png)

[繁體中文](./README.zh-tw.md) | [简体中文](./README.zh-cn.md) | [English](./README.md)

[![Lint and Testing](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci)
[![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci)

[Drone](https://www.drone.io/) 插件用于触发 [gitlab-ci](https://about.gitlab.com/solutions/continuous-integration/) 任务。

## GitLab 设置

请参阅 [Pipeline trigger tokens API](https://docs.gitlab.com/ee/api/pipeline_triggers.html) 的详细文档。您可以创建个人访问令牌来进行身份验证：

1. GitLab API。
2. GitLab 仓库。
3. GitLab 注册表。

请参阅 [GitLab token 概述](https://docs.gitlab.com/ee/security/tokens/index.html#personal-access-tokens)。

![token](./images/user_token.png)

如何获取项目 ID？前往您的项目 `设置 ➔ 一般` 下的一般项目。

![projectID](./images/projectID.png)

## 构建或下载二进制文件

可以从 [发布页面](https://github.com/appleboy/drone-gitlab-ci/releases) 下载预编译的二进制文件。支持以下操作系统类型。

- Windows amd64/386
- Linux amd64/386
- Darwin amd64/386

安装 `Go`

```sh
go install github.com/appleboy/drone-gitlab-ci
```

或者使用以下命令构建二进制文件：

```sh
make build
```

## 用法

有三种方法可以触发 gitlab-ci 任务。

### 从二进制文件使用

触发任务。

```bash
drone-gitlab-ci \
  --host https://gitlab.com/ \
  --token XXXXXXXX \
  --ref master \
  --project-id gitlab-ci-project-id
```

启用调试模式。

```bash
drone-gitlab-ci \
  --host https://gitlab.com/ \
  --token XXXXXXXX \
  --ref master \
  --project-id gitlab-ci-project-id \
  --debug
```

### 从 docker 使用

触发任务。

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com/
  -e GITLAB_TOKEN=xxxxx
  -e GITLAB_REF=master
  -e GITLAB_ID=gitlab-ci-project-id
  appleboy/drone-gitlab-ci
```

启用调试模式。

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com/ \
  -e GITLAB_TOKEN=xxxxx \
  -e GITLAB_REF=master \
  -e GITLAB_ID=gitlab-ci-project-id \
  -e GITLAB_DEBUG=true \
  appleboy/drone-gitlab-ci
```

### 从 drone ci 使用

从工作目录执行：

```sh
docker run --rm \
  -e PLUGIN_HOST=https://gitlab.com/ \
  -e PLUGIN_TOKEN=xxxxx \
  -e PLUGIN_REF=master \
  -e PLUGIN_ID=gitlab-ci-project-id \
  -e PLUGIN_DEBUG=true \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  appleboy/drone-gitlab-ci
```

您可以在 [此处](DOCS.md) 获取有关如何在 drone 中使用 scp 插件的更多 [信息](DOCS.md)。

## 测试

使用以下命令测试包：

```sh
make test
```

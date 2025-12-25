# drone-gitlab-ci

![logo](./images/logo.png)

[繁體中文](./README.zh-tw.md) | [简体中文](./README.zh-cn.md) | [English](./README.md)

[![Lint and Testing](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci)
[![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci)

[Drone](https://www.drone.io/) 插件用於觸發 [gitlab-ci](https://about.gitlab.com/solutions/continuous-integration/) 任務。

## GitLab 設定

請參閱 [Pipeline trigger tokens API](https://docs.gitlab.com/ee/api/pipeline_triggers.html) 的詳細文檔。您可以創建個人訪問令牌來進行身份驗證：

1. GitLab API。
2. GitLab 儲存庫。
3. GitLab 註冊表。

請參閱 [GitLab token 概述](https://docs.gitlab.com/ee/security/tokens/index.html#personal-access-tokens)。

![token](./images/user_token.png)

如何獲取項目 ID？前往您的項目 `設定 ➔ 一般` 下的一般項目。

![projectID](./images/projectID.png)

## 構建或下載二進制文件

可以從 [發佈頁面](https://github.com/appleboy/drone-gitlab-ci/releases) 下載預編譯的二進制文件。支持以下操作系統類型。

- Windows amd64/386
- Linux amd64/386
- Darwin amd64/386

安裝 `Go`

```sh
go install github.com/appleboy/drone-gitlab-ci
```

或使用以下命令構建二進制文件：

```sh
make build
```

## 用法

有三種方式觸發 gitlab-ci 任務。

### 從二進制文件使用

觸發任務。

```bash
drone-gitlab-ci \
  --host https://gitlab.com/ \
  --token XXXXXXXX \
  --ref master \
  --project-id gitlab-ci-project-id
```

啟用調試模式。

```bash
drone-gitlab-ci \
  --host https://gitlab.com/ \
  --token XXXXXXXX \
  --ref master \
  --project-id gitlab-ci-project-id \
  --debug
```

### 從 docker 使用

觸發任務。

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com/
  -e GITLAB_TOKEN=xxxxx
  -e GITLAB_REF=master
  -e GITLAB_ID=gitlab-ci-project-id
  appleboy/drone-gitlab-ci
```

啟用調試模式。

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com/ \
  -e GITLAB_TOKEN=xxxxx \
  -e GITLAB_REF=master \
  -e GITLAB_ID=gitlab-ci-project-id \
  -e GITLAB_DEBUG=true \
  appleboy/drone-gitlab-ci
```

### 從 drone ci 使用

從工作目錄執行：

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

您可以獲取更多有關如何在 drone 中使用 scp 插件的[信息](DOCS.md)。

## 測試

使用以下命令測試包：

```sh
make test
```

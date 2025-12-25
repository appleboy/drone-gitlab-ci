# drone-gitlab-ci

![logo](./images/gitlab-ci.png)

[繁體中文](./README.zh-tw.md) | [简体中文](./README.zh-cn.md) | [English](./README.md)

[![Lint and Testing](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-gitlab-ci/actions/workflows/trivy.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci)
[![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci)

## 什麼是 drone-gitlab-ci？

**drone-gitlab-ci** 是一個 CLI 工具和 CI/CD 插件，可以從任何環境觸發 [GitLab CI](https://about.gitlab.com/solutions/continuous-integration/) Pipeline。它作為不同 CI/CD 平台之間的橋樑，讓您能夠將 GitLab CI 無縫整合到現有的工作流程中。

## 為什麼使用這個工具？

在現代軟體開發中，團隊經常在不同的專案或服務中使用多個 CI/CD 平台。這會在以下情況產生挑戰：

- **跨平台整合**：從其他 CI/CD 系統（Drone、GitHub Actions、Jenkins 等）觸發 GitLab CI Pipeline
- **微服務編排**：部署後自動觸發相依服務的 Pipeline
- **統一 Pipeline 管理**：在多樣化的基礎架構中集中管理 Pipeline 觸發邏輯
- **混合雲工作流程**：連接不同平台和環境之間的 CI/CD 流程

```txt
┌─────────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   Drone CI      │     │                  │     │                 │
│   GitHub Actions│────▶│ drone-gitlab-ci  │────▶│   GitLab CI     │
│   Jenkins       │     │                  │     │   Pipeline      │
│   本地開發       │     └──────────────────┘     └─────────────────┘
└─────────────────┘
```

## 使用情境

| 情境                           | 說明                                                 |
| ------------------------------ | ---------------------------------------------------- |
| **Drone CI → GitLab CI**       | 作為 Drone 插件，在 Drone 建置後觸發 GitLab Pipeline |
| **GitHub Actions → GitLab CI** | 從 GitHub 儲存庫觸發 GitLab 部署                     |
| **Jenkins → GitLab CI**        | 將 GitLab CI 整合到現有的 Jenkins 工作流程           |
| **本地開發**                   | 在開發和測試期間手動觸發 Pipeline                    |
| **微服務**                     | 服務 A 部署後自動觸發服務 B 的 Pipeline              |

## 功能特點

- 透過 API 觸發 GitLab CI Pipeline
- 傳遞自訂變數到 Pipeline
- 等待 Pipeline 完成，可設定逾時時間
- 支援自架 GitLab 實例
- 跨平台執行檔（Windows、Linux、macOS）
- 提供 Docker 映像
- 原生支援 Drone CI、GitHub Actions 和其他 CI/CD 平台

## 目錄

- [drone-gitlab-ci](#drone-gitlab-ci)
  - [什麼是 drone-gitlab-ci？](#什麼是-drone-gitlab-ci)
  - [為什麼使用這個工具？](#為什麼使用這個工具)
  - [使用情境](#使用情境)
  - [功能特點](#功能特點)
  - [目錄](#目錄)
  - [快速開始](#快速開始)
  - [安裝](#安裝)
    - [下載預編譯執行檔](#下載預編譯執行檔)
    - [使用 Go 安裝](#使用-go-安裝)
    - [從原始碼編譯](#從原始碼編譯)
  - [設定](#設定)
    - [GitLab 設定](#gitlab-設定)
      - [1. 建立個人存取權杖](#1-建立個人存取權杖)
      - [2. 取得專案 ID](#2-取得專案-id)
    - [參數說明](#參數說明)
  - [使用方式](#使用方式)
    - [命令列](#命令列)
    - [Docker](#docker)
    - [Drone CI](#drone-ci)
    - [GitHub Actions](#github-actions)
  - [開發](#開發)
    - [執行測試](#執行測試)
    - [編譯執行檔](#編譯執行檔)
    - [程式碼品質](#程式碼品質)
  - [授權](#授權)

## 快速開始

三個步驟觸發 GitLab CI Pipeline：

```bash
# 1. 下載執行檔（或使用 Docker）
go install github.com/appleboy/drone-gitlab-ci@latest

# 2. 設定認證資訊
export GITLAB_TOKEN=your-gitlab-token
export GITLAB_PROJECT_ID=your-project-id

# 3. 觸發 Pipeline
drone-gitlab-ci --ref main
```

## 安裝

### 下載預編譯執行檔

從 [發佈頁面](https://github.com/appleboy/drone-gitlab-ci/releases) 下載。支援的平台：

- Windows (amd64/386)
- Linux (amd64/386)
- macOS (amd64/arm64)

### 使用 Go 安裝

```bash
go install github.com/appleboy/drone-gitlab-ci@latest
```

### 從原始碼編譯

```bash
git clone https://github.com/appleboy/drone-gitlab-ci.git
cd drone-gitlab-ci
make build
```

## 設定

### GitLab 設定

#### 1. 建立個人存取權杖

您需要 GitLab 權杖來驗證 API 請求。請參閱 [GitLab 權杖概述](https://docs.gitlab.com/ee/security/tokens/index.html#personal-access-tokens)。

1. 前往 GitLab → 使用者設定 → 存取權杖
2. 建立具有 `api` 範圍的權杖
3. 安全保存權杖

![token](./images/user_token.png)

#### 2. 取得專案 ID

在 **設定 → 一般 → 一般專案設定** 中找到您的專案 ID。

![projectID](./images/projectID.png)

更多詳情請參閱 [Pipeline trigger tokens API](https://docs.gitlab.com/ee/api/pipeline_triggers.html)。

### 參數說明

| 參數       | 旗標                 | 環境變數                                                  | 說明                     | 必填   | 預設值               |
| ---------- | -------------------- | --------------------------------------------------------- | ------------------------ | ------ | -------------------- |
| Host       | `--host`             | `GITLAB_HOST`, `PLUGIN_HOST`, `INPUT_HOST`                | GitLab 實例 URL          | 否     | `https://gitlab.com` |
| Token      | `--token`, `-t`      | `GITLAB_TOKEN`, `PLUGIN_TOKEN`, `INPUT_TOKEN`             | GitLab 存取權杖          | **是** | -                    |
| Project ID | `--project-id`, `-p` | `GITLAB_PROJECT_ID`, `PLUGIN_ID`, `INPUT_PROJECT_ID`      | GitLab 專案 ID           | **是** | -                    |
| Ref        | `--ref`, `-r`        | `GITLAB_REF`, `PLUGIN_REF`, `INPUT_REF`                   | 要觸發的分支或標籤       | 否     | `main`               |
| Variables  | `--variables`        | `GITLAB_VARIABLES`, `PLUGIN_VARIABLES`, `INPUT_VARIABLES` | 傳遞的變數 (KEY=VALUE)   | 否     | -                    |
| Wait       | `--wait`, `-w`       | `GITLAB_WAIT`, `PLUGIN_WAIT`, `INPUT_WAIT`                | 等待 Pipeline 完成       | 否     | `false`              |
| Timeout    | `--timeout`          | `GITLAB_TIMEOUT`, `PLUGIN_TIMEOUT`, `INPUT_TIMEOUT`       | 等待逾時時間             | 否     | `60m`                |
| Interval   | `--interval`, `-i`   | `GITLAB_INTERVAL`, `PLUGIN_INTERVAL`, `INPUT_INTERVAL`    | 輪詢間隔                 | 否     | `5s`                 |
| Insecure   | `--insecure`, `-k`   | `GITLAB_INSECURE`, `PLUGIN_INSECURE`, `INPUT_INSECURE`    | 跳過 SSL 驗證            | 否     | `false`              |
| Debug      | `--debug`, `-d`      | `GITLAB_DEBUG`, `PLUGIN_DEBUG`, `INPUT_DEBUG`             | 啟用除錯輸出             | 否     | `false`              |
| GitHub     | `--github`           | `GITHUB_ACTIONS`, `PLUGIN_GITHUB`, `INPUT_GITHUB`         | 啟用 GitHub Actions 輸出 | 否     | `false`              |

## 使用方式

### 命令列

基本用法：

```bash
drone-gitlab-ci \
  --host https://gitlab.com \
  --token your-token \
  --project-id 12345 \
  --ref main
```

帶變數並等待完成：

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

啟用除錯模式：

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

啟用等待和除錯：

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

新增到您的 `.drone.yml`：

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

建立 `.github/workflows/trigger-gitlab.yml`：

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

## 開發

### 執行測試

```bash
make test
```

### 編譯執行檔

```bash
make build
```

### 程式碼品質

```bash
make lint
```

## 授權

MIT 授權 - 詳見 [LICENSE](LICENSE) 檔案。

<img src="images/logo.png">

# drone-gitlab-ci

[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci) [![Build Status](http://drone.wu-boy.com/api/badges/appleboy/drone-gitlab-ci/status.svg)](http://drone.wu-boy.com/appleboy/drone-gitlab-ci) [![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci) [![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-gitlab-ci.svg)](https://hub.docker.com/r/appleboy/drone-gitlab-ci/) [![](https://images.microbadger.com/badges/image/appleboy/drone-gitlab-ci.svg)](https://microbadger.com/images/appleboy/drone-gitlab-ci "Get your own image badge on microbadger.com")

[Drone](https://github.com/drone/drone) plugin for trigger [gitlab-ci](https://about.gitlab.com/gitlab-ci) jobs.

## Build or Download a binary

The pre-compiled binaries can be downloaded from [release page](https://github.com/appleboy/drone-gitlab-ci/releases). Support the following OS type.

* Windows amd64/386
* Linux amd64/386
* Darwin amd64/386

With `Go` installed

```
$ go get -u -v github.com/appleboy/drone-gitlab-ci
```

or build the binary with the following command:

```
$ make build
```

## Docker

Build the docker image with the following commands:

```
$ make docker
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-gitlab-ci' not found or does not exist..
```

## Usage

There are three ways to trigger gitlab-ci jobs.

* [usage from binary](#usage-from-binary)
* [usage from docker](#usage-from-docker)
* [usage from drone ci](#usage-from-drone-ci)

<a name="usage-from-binary"></a>
### Usage from binary

trigger job.

```bash
drone-gitlab-ci \
  --host https://gitlab.com/ \
  --token XXXXXXXX \
  --ref master \
  --id gitlab-ci-project-id
```

Enable debug mode.

```diff
drone-gitlab-ci \
  --host https://gitlab.com/ \
  --token XXXXXXXX \
  --ref master \
  --id gitlab-ci-project-id
+ --debug
```

<a name="usage-from-docker"></a>
### Usage from docker

trigger job.

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com/
  -e GITLAB_TOKEN=xxxxx
  -e GITLAB_REF=master
  -e GITLAB_ID=gitlab-ci-project-id
  appleboy/drone-gitlab-ci
```

Enable debug mode.

```bash
docker run --rm \
  -e GITLAB_HOST=https://gitlab.com/ \
  -e GITLAB_TOKEN=xxxxx \
  -e GITLAB_REF=master \
  -e GITLAB_ID=gitlab-ci-project-id \
  -e GITLAB_DEBUG=true \
  appleboy/drone-gitlab-ci
```

<a name="usage-from-drone-ci"></a>
### Usage from drone ci

Execute from the working directory:

```
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

You can get more [information](DOCS.md) about how to use scp plugin in drone.

## Testing

Test the package with the following command:

```
$ make test
```

<img src="images/logo.png">

# drone-gitlab-ci

[![GoDoc](https://godoc.org/github.com/appleboy/drone-gitlab-ci?status.svg)](https://godoc.org/github.com/appleboy/drone-gitlab-ci) [![Build Status](http://drone.wu-boy.com/api/badges/appleboy/drone-gitlab-ci/status.svg)](http://drone.wu-boy.com/appleboy/drone-gitlab-ci) [![codecov](https://codecov.io/gh/appleboy/drone-gitlab-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-gitlab-ci) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-gitlab-ci)](https://goreportcard.com/report/github.com/appleboy/drone-gitlab-ci) [![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-gitlab-ci.svg)](https://hub.docker.com/r/appleboy/drone-gitlab-ci/) [![](https://images.microbadger.com/badges/image/appleboy/drone-gitlab-ci.svg)](https://microbadger.com/images/appleboy/drone-gitlab-ci "Get your own image badge on microbadger.com")

[Drone](https://github.com/drone/drone) plugin for trigger [gitlab-ci](https://gitlab-ci.io/) jobs.

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

trigger single job.

```bash
drone-gitlab-ci \
  --host http://gitlab-ci.example.com/ \
  --user appleboy \
  --token XXXXXXXX \
  --job drone-gitlab-ci-plugin
```

trigger multiple jobs.

```bash
drone-gitlab-ci \
  --host http://gitlab-ci.example.com/ \
  --user appleboy \
  --token XXXXXXXX \
  --job drone-gitlab-ci-plugin-1 \
  --job drone-gitlab-ci-plugin-2
```

<a name="usage-from-docker"></a>
### Usage from docker

trigger single job.

```bash
docker run --rm \
  -e gitlab-ci_BASE_URL=http://gitlab-ci.example.com/
  -e gitlab-ci_USER=appleboy
  -e gitlab-ci_TOKEN=xxxxxxx
  -e gitlab-ci_JOB=drone-gitlab-ci-plugin
  appleboy/drone-gitlab-ci
```

trigger multiple jobs.

```bash
docker run --rm \
  -e gitlab-ci_BASE_URL=http://gitlab-ci.example.com/
  -e gitlab-ci_USER=appleboy
  -e gitlab-ci_TOKEN=xxxxxxx
  -e gitlab-ci_JOB=drone-gitlab-ci-plugin-1,drone-gitlab-ci-plugin-2
  appleboy/drone-gitlab-ci
```

<a name="usage-from-drone-ci"></a>
### Usage from drone ci

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_URL=http://example.com \
  -e PLUGIN_USER=xxxxxxx \
  -e PLUGIN_TOKEN=xxxxxxx \
  -e PLUGIN_JOB=xxxxxxx \
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

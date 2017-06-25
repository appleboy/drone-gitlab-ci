---
date: 2017-01-16T00:00:00+00:00
title: Gitlab-ci
author: appleboy
tags: [ infrastructure, trigger, gitlab-ci ]
repo: appleboy/drone-gitlab-ci
logo: gitlab-ci.svg
image: appleboy/drone-gitlab-ci
---

The Gitlab-ci plugin allows you to trigger Gitlab-ci job automatically. The below pipeline configuration demonstrates simple usage:

```yaml
pipeline:
  gitlab-ci:
    image: appleboy/drone-gitlab-ci
    url: http://example.com
    user: appleboy
    token: xxxxxxxxxx
    job: drone-gitlab-ci-plugin-job
```

Example configuration for success builds:

```diff
pipeline:
  gitlab-ci:
    image: appleboy/drone-gitlab-ci
    url: http://example.com
    user: appleboy
    token: xxxxxxxxxx
    job: drone-gitlab-ci-plugin-job
+   when:
+     status: [ success ]
```

Example configuration with multiple jobs:

```yaml
pipeline:
  gitlab-ci:
    image: appleboy/drone-gitlab-ci
    url: http://example.com
    user: appleboy
    token: xxxxxxxxxx
    job:
+     - drone-gitlab-ci-plugin-job-1
+     - drone-gitlab-ci-plugin-job-2
```

Example configuration with jobs in the folder:

```yaml
pipeline:
  gitlab-ci:
    image: appleboy/drone-gitlab-ci
    url: http://example.com
    user: appleboy
    token: xxxxxxxxxx
+   job: folder_name/job_name
```

It will trigger the URL of Gitlab-ci job like as `http://example.com/job/folder_name/job/job_name/`

# Parameter Reference

url
: gitlab-ci server base url.

user
: gitlab-ci user account

token
: gitlab-ci user token

job
: gitlab-ci job name

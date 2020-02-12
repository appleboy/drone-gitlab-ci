---
date: 2019-19-19T00:00:00+00:00
title: Gitlab-CI
author: appleboy
tags: [ infrastructure, trigger, gitlab, gitlab-ci ]
repo: appleboy/drone-gitlab-ci
logo: gitlab-ci.svg
image: appleboy/drone-gitlab-ci
---

The Gitlab-ci plugin allows you to trigger Gitlab-ci job automatically. The below pipeline configuration demonstrates simple usage:

```yaml
- name: trigger gitlab job
  image: appleboy/drone-gitlab-ci
  settings:
    host: https://gitlab.com
    token: xxxxxxxxxx
    ref: master
    id: gitlab-project-id
```

Example configuration with debug mode:

```diff
  - name: trigger gitlab job
    image: appleboy/drone-gitlab-ci
    settings:
      host: https://gitlab.com
      token: xxxxxxxxxx
      ref: master
      id: gitlab-project-id
+     debug: true
```

Example configuration using credentials from secrets:

```diff
  - name: trigger gitlab job
    image: appleboy/drone-gitlab-ci
    settings:
      host: https://gitlab.com
+     token:
+       from_secret: gitlab_token
      id: gitlab-project-id
```

Example configuration with environment variables set on the gitlab pipeline:

```diff
  - name: trigger gitlab job
    image: appleboy/drone-gitlab-ci
    settings:
      host: https://gitlab.com
      token:
        from_secret: gitlab_token
      id: gitlab-project-id
+     gitlab-env:
+       - VARIABLE_NAME=VALUE
```

Example config where the job will wait on the gitlab pipeline completing:

```diff
  - name: trigger gitlab job
    image: appleboy/drone-gitlab-ci
    settings:
      host: https://gitlab.com
      token:
        from_secret: gitlab_token
      id: gitlab-project-id
+     wait: true
```

## Parameter Reference

host
: gitlab-ci server base url.

token
: gitlab-ci user token

ref
: valid refs are only the branches and tags

id
: gitlab-ci project id

debug
: enable debug mode

gitlab-env
: list of strings in the form `VARIABLE_NAME=VALUE` which will be passed to the pipeline.

wait
: bool - if true, will wait for pipeline to finish before reporting success or failure.

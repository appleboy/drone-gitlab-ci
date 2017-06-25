FROM centurylink/ca-certs

ADD drone-gitlab-ci /

ENTRYPOINT ["/drone-gitlab-ci"]

# go-gitlab-client

[![Build Status](https://travis-ci.org/plouc/go-gitlab-client.png?branch=master)](https://travis-ci.org/plouc/go-gitlab-client)

**go-gitlab-client** is a simple client written in golang to consume gitlab API.

It also provides an handy CLI to easily interact with gitlab through CLI.

- [lib](#lib)
  - [install](#install-lib)
  - [update](#update)
  - [documentation](#documentation)
  - [supported APIs](#supported-apis)
- [cli](#cli)
  - [features](#cli-features)
  - [install](#install-cli)

## lib

### Install lib

To install go-gitlab-client, use `go get`:

```
go get github.com/plouc/go-gitlab-client/gitlab
```

Import the `go-gitlab-client` package into your code:

```go
package whatever

import (
    "github.com/plouc/go-gitlab-client/gogitlab"
)
```

### Update

To update `go-gitlab-client`, use `go get -u`:

    go get -u github.com/plouc/go-gitlab-client

### Documentation

Visit the docs at http://godoc.org/github.com/plouc/go-gitlab-client

### Supported APIs

#### Projects

[gitlab api doc](http://doc.gitlab.com/ce/api/projects.html)

- [x] List all projects
- [ ] List user projects
- [x] Get single project
- [x] Remove project
- [x] Star a project
- [x] Unstar a project

#### Repositories

[gitlab api doc](http://doc.gitlab.com/ce/api/repositories.html)

- list project repository tags
- list repository commits
- list project hooks
- add/get/edit/rm project hook

#### Users

[gitlab api doc](http://api.gitlab.org/users.html)

- [x] List users
- [x] Single user
- [x] Current user

#### Groups

[gitlab api doc](https://docs.gitlab.com/ce/api/groups.html)

- [x] List groups
- [ ] List a groups's subgroups
- [ ] List a group's projects
- [x] Details of a group
- [x] New group
- [ ] Transfer project to group
- [ ] Update group
- [x] Remove group
- [x] Search for group
- [x] Group members

#### Deploy Keys

[gitlab api doc](http://doc.gitlab.com/ce/api/deploy_keys.html)

- list project deploy keys
- add/get/rm project deploy key

#### Builds

[gitlab api doc](http://doc.gitlab.com/ce/api/builds.html)

- List project builds
- Get a single build
- List commit builds
- Get build artifacts
- Cancel a build
- Retry a build
- Erase a build

#### Runners

[gitlab api doc](https://docs.gitlab.com/ee/api/runners.html)

- [x] List owned runners
- [x] List all runners
- [x] Get runner's details
- [ ] Update runner's details
- [ ] Remove a runner
- [ ] List runner's jobs
- [ ] List project's runners
- [ ] Enable a runner in project
- [ ] Disable a runner from project
- [ ] Register a new Runner
- [ ] Delete a registered Runner
- [ ] Verify authentication for a registered Runner

#### Branches

[gitlab api doc](https://docs.gitlab.com/ee/api/branches.html)

- [x] List repository branches
- [x] Get single repository branch
- [x] Protect repository branch
- [x] Unprotect repository branch
- [x] Create repository branch
- [x] Delete repository branch
- [x] Delete merged branches

#### Project hooks

[gitlab api doc](https://docs.gitlab.com/ee/api/projects.html#hooks)

- [x] List project hooks
- [x] Get project hook
- [x] Add project hook
- [ ] Edit project hook
- [x] Delete project hook

#### Project-level variables [gitlab api doc](https://docs.gitlab.com/ee/api/project_level_variables.html)

- [x] List project variables
- [x] Show project variable details
- [x] Create project variable
- [ ] Update project variable
- [x] Remove project variable

#### Group-level variables

[gitlab api doc](https://docs.gitlab.com/ee/api/group_level_variables.html)

- [x] List group variables
- [x] Show group variable details
- [x] Create group variable
- [ ] Update group variable
- [x] Remove group variable

#### Pipelines

[gitlab api doc](https://docs.gitlab.com/ee/api/pipelines.html)

- [x] List project pipelines
- [x] Get a single pipeline
- [ ] Create a new pipeline
- [ ] Retry jobs in a pipeline
- [ ] Cancel a pipeline's jobs

#### Project badges

[gitlab api doc](https://docs.gitlab.com/ee/api/project_badges.html)

- [x] List all badges of a project
- [x] Get a badge of a project
- [x] Add a badge to a project
- [ ] Edit a badge of a project
- [x] Remove a badge from a project
- [ ] Preview a badge from a project

#### Namespaces

[gitlab api doc](https://docs.gitlab.com/ee/api/namespaces.html)

- [x] List namespaces
- [x] Search for namespace
- [x] Get namespace by ID

## CLI

### CLI features

- normalized operations: `ls`, `get`, `add`, `update`
- resource aliases for easy retrieval
- `text`, `yaml` & `json` output
- saving output to file
- interactive pagination mode
- interactive resource creation

### install CLI

@todo

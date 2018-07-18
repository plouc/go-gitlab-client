# go-gitlab-client

[![Build Status](https://travis-ci.org/plouc/go-gitlab-client.png?branch=v2)](https://travis-ci.org/plouc/go-gitlab-client)

**go-gitlab-client** is a client written in golang to consume gitlab API.

It also provides an handy CLI to easily interact with gitlab API.

- [lib](#lib)
  - [install](#install-lib)
  - [update](#update)
  - [documentation](#documentation)
  - [supported APIs](#supported-apis)
- [CLI](#cli)
  - [features](#cli-features)
  - [install](#install-cli)
  - [commands](#cli-commands)

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
    "github.com/plouc/go-gitlab-client/gitlab"
)
```

### Update

To update `go-gitlab-client`, use `go get -u`:

    go get -u github.com/plouc/go-gitlab-client/gitlab

### Documentation

Visit the docs at http://godoc.org/github.com/plouc/go-gitlab-client/gitlab

### Supported APIs

#### Branches

[gitlab api doc](https://docs.gitlab.com/ee/api/branches.html)

- [x] List repository branches
- [x] Get single repository branch
- [x] Protect repository branch
- [x] Unprotect repository branch
- [x] Create repository branch
- [x] Delete repository branch
- [x] Delete merged branches

#### Project-level variables

[gitlab api doc](https://docs.gitlab.com/ee/api/project_level_variables.html)

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

#### Commits

[gitlab api doc](http://doc.gitlab.com/ce/api/commits.html)

- [x] List repository commits
- [ ] Create a commit with multiple files and actions
- [x] Get a single commit
- [x] Get references a commit is pushed to
- [ ] Cherry pick a commit
- [ ] Get the diff of a commit
- [ ] Get the comments of a commit
- [ ] Post comment to commit
- [x] List the statuses of a commit
- [ ] Post the build status to a commit
- [ ] List Merge Requests associated with a commit

#### Deploy Keys

[gitlab api doc](http://doc.gitlab.com/ce/api/deploy_keys.html)

- list project deploy keys
- add/get/rm project deploy key

#### Environments

[gitlab api doc](https://docs.gitlab.com/ee/api/environments.html)

- [x] List environments
- [x] Create a new environment
- [ ] Edit an existing environment
- [x] Delete an environment
- [ ] Stop an environment

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

#### Jobs

[gitlab api doc](http://doc.gitlab.com/ce/api/jobs.html)

- [x] List project jobs
- [x] List pipeline jobs
- [x] Get a single job
- [ ] Get job artifacts
- [ ] Download the artifacts archive
- [ ] Download a single artifact file
- [x] Get a trace file
- [x] Cancel a job
- [x] Retry a job
- [x] Erase a job
- [ ] Keep artifacts
- [ ] Play a job

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

#### SSH Keys

[gitlab api doc](http://api.gitlab.org/users.html#list-ssh-keys)

- [x] List SSH keys
- [x] List SSH keys for user
- [x] Single SSH key
- [x] Add SSH key
- [x] Add SSH key for user
- [x] Delete SSH key for current user
- [x] Delete SSH key for given user

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

#### Project hooks

[gitlab api doc](https://docs.gitlab.com/ee/api/projects.html#hooks)

- [x] List project hooks
- [x] Get project hook
- [x] Add project hook
- [ ] Edit project hook
- [x] Delete project hook

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

#### Merge requests

[gitlab api doc](https://docs.gitlab.com/ee/api/merge_requests.html)

- [x] List merge requests
- [x] List project merge requests
- [x] List group merge requests
- [x] Get single MR
- [ ] Get single MR participants
- [ ] Get single MR commits
- [ ] Get single MR changes
- [ ] List MR pipelines
- [ ] Create MR
- [ ] Update MR
- [ ] Delete a merge request
- [ ] Accept MR
- [ ] Cancel Merge When Pipeline Succeeds
- [ ] Comments on merge requests
- [ ] List issues that will close on merge
- [ ] Subscribe to a merge request
- [ ] Unsubscribe from a merge request
- [ ] Create a todo
- [ ] Get MR diff versions
- [ ] Get a single MR diff version
- [ ] Set a time estimate for a merge request
- [ ] Reset the time estimate for a merge request
- [ ] Add spent time for a merge request
- [ ] Reset spent time for a merge request
- [ ] Get time tracking stats
- [ ] Approvals

#### Notes

[gitlab api doc](https://docs.gitlab.com/ee/api/notes.html)

- Issues
  - [x] List project issue notes
  - [x] Get single issue note
  - [x] Create new issue note
  - [ ] Modify existing issue note
  - [ ] Delete an issue note
- Snippets
  - [x] List all snippet notes
  - [x] Get single snippet note
  - [x] Create new snippet note
  - [ ] Modify existing snippet note
  - [ ] Delete a snippet note
- Merge Requests
  - [x] List all merge request notes
  - [x] Get single merge request note
  - [x] Create new merge request note
  - [ ] Modify existing merge request note
  - [ ] Delete a merge request note
- Epics
  - [x] List all epic notes
  - [x] Get single epic note
  - [x] Create new epic note
  - [ ] Modify existing epic note
  - [ ] Delete an epic note

## CLI

**go-gitlab-client** provides a CLI to easily interact with GitLab API, **glc**.

### install CLI

**glc** is a single binary with no external dependencies, released for several platforms.
Go to the [releases page](https://github.com/plouc/go-gitlab-client/releases),
download the package for your OS, and copy the binary to somewhere on your PATH.
Please make sure to rename the binary to `glc` and make it executable.

You can also install completion for bash or zsh, please run `glc help completion`
for more info.

### CLI features

- normalized operations: `ls`, `get`, `add`, `update`
- resource aliases for easy retrieval
- `text`, `yaml` & `json` output
- saving output to file
- interactive pagination mode
- interactive resource creation

### CLI commands

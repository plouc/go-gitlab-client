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


- [glc add](#glc-add)	*Add resource*
- [glc add alias](#glc-add-alias)	*Create resource alias*
- [glc add group](#glc-add-group)	*Create a new group*
- [glc add group-epic-note](#glc-add-group-epic-note)	*Add group epic note*
- [glc add group-var](#glc-add-group-var)	*Create a new group variable*
- [glc add project](#glc-add-project)	*Create a new project*
- [glc add project-badge](#glc-add-project-badge)	*Create project badge*
- [glc add project-branch](#glc-add-project-branch)	*Create project branch*
- [glc add project-environment](#glc-add-project-environment)	*Create project environment*
- [glc add project-hook](#glc-add-project-hook)	*Create a new hook for given project*
- [glc add project-issue-note](#glc-add-project-issue-note)	*Add project issue note*
- [glc add project-merge-request-note](#glc-add-project-merge-request-note)	*Add project issue note*
- [glc add project-protected-branch](#glc-add-project-protected-branch)	*Protect project branch*
- [glc add project-snippet-note](#glc-add-project-snippet-note)	*Add project snippet note*
- [glc add project-star](#glc-add-project-star)	*Stars a given project*
- [glc add project-var](#glc-add-project-var)	*Create a new project variable*
- [glc ci-info](#glc-ci-info)	*Print information about CI environment*
- [glc completion](#glc-completion)	*Output shell completion code for the specified shell (bash or zsh)*
- [glc doc](#glc-doc)	*Generate CLI documentation in markdown format*
- [glc get](#glc-get)	*Get resource details*
- [glc get current-user](#glc-get-current-user)	*Get current user*
- [glc get group](#glc-get-group)	*Get all details of a group*
- [glc get group-var](#glc-get-group-var)	*Get the details of a group's specific variable*
- [glc get namespace](#glc-get-namespace)	*Get a single namespace*
- [glc get project](#glc-get-project)	*Get a specific project*
- [glc get project-badge](#glc-get-project-badge)	*Get project badge info*
- [glc get project-branch](#glc-get-project-branch)	*Get project branch info*
- [glc get project-hook](#glc-get-project-hook)	*Get project hook info*
- [glc get project-job](#glc-get-project-job)	*Get project job info*
- [glc get project-job cancel](#glc-get-project-job-cancel)	*Cancel project job*
- [glc get project-job retry](#glc-get-project-job-retry)	*Retry project job*
- [glc get project-job-trace](#glc-get-project-job-trace)	*Get project job trace*
- [glc get project-merge-request](#glc-get-project-merge-request)	*Get project merge request info*
- [glc get project-merge-request-note](#glc-get-project-merge-request-note)	*Get project merge request note*
- [glc get project-pipeline](#glc-get-project-pipeline)	*Get project pipeline details*
- [glc get project-var](#glc-get-project-var)	*Get the details of a project's specific variable*
- [glc get runner](#glc-get-runner)	*Get details of a runner*
- [glc get user](#glc-get-user)	*Get a single user*
- [glc init](#glc-init)	*Init glc config*
- [glc list](#glc-list)	*List resource*
- [glc list aliases](#glc-list-aliases)	*List resource aliases*
- [glc list group-epic-notes](#glc-list-group-epic-notes)	*List group epic notes*
- [glc list group-merge-requests](#glc-list-group-merge-requests)	*List group merge requests*
- [glc list group-variables](#glc-list-group-variables)	*Get list of a group's variables*
- [glc list groups](#glc-list-groups)	*List groups*
- [glc list merge-requests](#glc-list-merge-requests)	*List merge requests*
- [glc list namespaces](#glc-list-namespaces)	*List namespaces*
- [glc list project-badges](#glc-list-project-badges)	*List project badges*
- [glc list project-branches](#glc-list-project-branches)	*List project branches*
- [glc list project-commits](#glc-list-project-commits)	*List project repository commits*
- [glc list project-environments](#glc-list-project-environments)	*List project environments*
- [glc list project-hooks](#glc-list-project-hooks)	*List project's hooks*
- [glc list project-issue-notes](#glc-list-project-issue-notes)	*List project issue notes*
- [glc list project-jobs](#glc-list-project-jobs)	*List project jobs*
- [glc list project-members](#glc-list-project-members)	*List project members*
- [glc list project-merge-request-commits](#glc-list-project-merge-request-commits)	*List project merge request commits*
- [glc list project-merge-request-notes](#glc-list-project-merge-request-notes)	*List project merge request notes*
- [glc list project-merge-requests](#glc-list-project-merge-requests)	*List project merge requests*
- [glc list project-pipeline-jobs](#glc-list-project-pipeline-jobs)	*List project pipeline jobs*
- [glc list project-pipelines](#glc-list-project-pipelines)	*List project pipelines*
- [glc list project-protected-branches](#glc-list-project-protected-branches)	*List project protected branches*
- [glc list project-snippet-notes](#glc-list-project-snippet-notes)	*List project snippet notes*
- [glc list project-variables](#glc-list-project-variables)	*Get list of a project's variables*
- [glc list projects](#glc-list-projects)	*List projects*
- [glc list runners](#glc-list-runners)	*List runners*
- [glc list ssh-keys](#glc-list-ssh-keys)	*List current user ssh keys*
- [glc list user-ssh-keys](#glc-list-user-ssh-keys)	*List specific user ssh keys*
- [glc list users](#glc-list-users)	*List users*
- [glc rm](#glc-rm)	*Remove resource*
- [glc rm alias](#glc-rm-alias)	*Remove resource alias*
- [glc rm group](#glc-rm-group)	*Remove group*
- [glc rm group-epic-note](#glc-rm-group-epic-note)	*Remove group epic note*
- [glc rm group-var](#glc-rm-group-var)	*Remove a group's variable*
- [glc rm project](#glc-rm-project)	*Remove project*
- [glc rm project-badge](#glc-rm-project-badge)	*Remove project badge*
- [glc rm project-branch](#glc-rm-project-branch)	*Remove project branch*
- [glc rm project-environment](#glc-rm-project-environment)	*Remove project environment*
- [glc rm project-hook](#glc-rm-project-hook)	*Remove project hook*
- [glc rm project-issue-note](#glc-rm-project-issue-note)	*Remove project issue note*
- [glc rm project-merge-request-note](#glc-rm-project-merge-request-note)	*Remove project merge request note*
- [glc rm project-merged-branches](#glc-rm-project-merged-branches)	*Remove project merged branches*
- [glc rm project-protected-branch](#glc-rm-project-protected-branch)	*Unprotect project branch*
- [glc rm project-snippet-note](#glc-rm-project-snippet-note)	*Remove project snippet note*
- [glc rm project-star](#glc-rm-project-star)	*Unstars a given project*
- [glc rm project-var](#glc-rm-project-var)	*Remove a project's variable*
- [glc version](#glc-version)	*Print the version number of glc*



#### glc add

Add resource

##### Synopsis

Add resource

##### Options

```
  -h, --help   help for add
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add alias](#glc-add-alias)	*Create resource alias*
- [glc add group](#glc-add-group)	*Create a new group*
- [glc add group-epic-note](#glc-add-group-epic-note)	*Add group epic note*
- [glc add group-var](#glc-add-group-var)	*Create a new group variable*
- [glc add project](#glc-add-project)	*Create a new project*
- [glc add project-badge](#glc-add-project-badge)	*Create project badge*
- [glc add project-branch](#glc-add-project-branch)	*Create project branch*
- [glc add project-environment](#glc-add-project-environment)	*Create project environment*
- [glc add project-hook](#glc-add-project-hook)	*Create a new hook for given project*
- [glc add project-issue-note](#glc-add-project-issue-note)	*Add project issue note*
- [glc add project-merge-request-note](#glc-add-project-merge-request-note)	*Add project issue note*
- [glc add project-protected-branch](#glc-add-project-protected-branch)	*Protect project branch*
- [glc add project-snippet-note](#glc-add-project-snippet-note)	*Add project snippet note*
- [glc add project-star](#glc-add-project-star)	*Stars a given project*
- [glc add project-var](#glc-add-project-var)	*Create a new project variable*



#### glc add alias

Create resource alias

##### Synopsis

Create resource alias

```
glc add alias ALIAS RESOURCE_TYPE [...resource ids] [flags]
```

##### Options

```
  -h, --help   help for alias
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add group

Create a new group

##### Synopsis

Create a new group

```
glc add group [flags]
```

##### Options

```
  -h, --help   help for group
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add group-epic-note

Add group epic note

##### Synopsis

Add group epic note

```
glc add group-epic-note GROUP_ID EPIC_ID [flags]
```

##### Options

```
  -h, --help   help for group-epic-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add group-var

Create a new group variable

##### Synopsis

Create a new group variable

```
glc add group-var GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project

Create a new project

##### Synopsis

Create a new project

```
glc add project [flags]
```

##### Options

```
  -h, --help   help for project
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-badge

Create project badge

##### Synopsis

Create project badge

```
glc add project-badge PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-badge
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-branch

Create project branch

##### Synopsis

Create project branch

```
glc add project-branch PROJECT_ID [flags]
```

##### Options

```
  -b, --branch string   Name of the branch
  -h, --help            help for project-branch
  -r, --ref string      	The branch name or commit SHA to create branch from
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-environment

Create project environment

##### Synopsis

Create project environment

```
glc add project-environment PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-environment
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-hook

Create a new hook for given project

##### Synopsis

Create a new hook for given project

```
glc add project-hook PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-hook
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-issue-note

Add project issue note

##### Synopsis

Add project issue note

```
glc add project-issue-note PROJECT_ID ISSUE_IID [flags]
```

##### Options

```
  -h, --help   help for project-issue-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-merge-request-note

Add project issue note

##### Synopsis

Add project issue note

```
glc add project-merge-request-note PROJECT_ID MERGE_REQUEST_IID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-protected-branch

Protect project branch

##### Synopsis

Protect project branch

```
glc add project-protected-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-protected-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-snippet-note

Add project snippet note

##### Synopsis

Add project snippet note

```
glc add project-snippet-note PROJECT_ID SNIPPET_ID [flags]
```

##### Options

```
  -h, --help   help for project-snippet-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-star

Stars a given project

##### Synopsis

Stars a given project

```
glc add project-star PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-star
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc add project-var

Create a new project variable

##### Synopsis

Create a new project variable

```
glc add project-var PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc add](#glc-add)	*Add resource*



#### glc ci-info

Print information about CI environment

##### Synopsis

Print information about CI environment

```
glc ci-info [flags]
```

##### Options

```
  -h, --help   help for ci-info
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc completion

Output shell completion code for the specified shell (bash or zsh)

##### Synopsis

Output shell completion code for the specified shell (bash or zsh).
The shell code must be evaluated to provide interactive
completion of glc commands.
This can be done by sourcing it from the .bash_profile or .zshrc.
For bash you can run:

  echo "source <(kubectl completion bash)" >> ~/.bashrc


```
glc completion [shell] [flags]
```

##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc doc

Generate CLI documentation in markdown format

##### Synopsis

Generate CLI documentation in markdown format

```
glc doc [flags]
```

##### Options

```
  -h, --help   help for doc
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc get

Get resource details

##### Synopsis

Get resource details

##### Options

```
  -h, --help   help for get
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get current-user](#glc-get-current-user)	*Get current user*
- [glc get group](#glc-get-group)	*Get all details of a group*
- [glc get group-var](#glc-get-group-var)	*Get the details of a group's specific variable*
- [glc get namespace](#glc-get-namespace)	*Get a single namespace*
- [glc get project](#glc-get-project)	*Get a specific project*
- [glc get project-badge](#glc-get-project-badge)	*Get project badge info*
- [glc get project-branch](#glc-get-project-branch)	*Get project branch info*
- [glc get project-hook](#glc-get-project-hook)	*Get project hook info*
- [glc get project-job](#glc-get-project-job)	*Get project job info*
- [glc get project-job-trace](#glc-get-project-job-trace)	*Get project job trace*
- [glc get project-merge-request](#glc-get-project-merge-request)	*Get project merge request info*
- [glc get project-merge-request-note](#glc-get-project-merge-request-note)	*Get project merge request note*
- [glc get project-pipeline](#glc-get-project-pipeline)	*Get project pipeline details*
- [glc get project-var](#glc-get-project-var)	*Get the details of a project's specific variable*
- [glc get runner](#glc-get-runner)	*Get details of a runner*
- [glc get user](#glc-get-user)	*Get a single user*



#### glc get current-user

Get current user

##### Synopsis

Get current user

```
glc get current-user [flags]
```

##### Options

```
  -h, --help   help for current-user
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get group

Get all details of a group

##### Synopsis

Get all details of a group

```
glc get group GROUP_ID [flags]
```

##### Options

```
  -h, --help                     help for group
  -x, --with-custom-attributes   Include custom attributes (admins only)
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get group-var

Get the details of a group's specific variable

##### Synopsis

Get the details of a group's specific variable

```
glc get group-var GROUP_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for group-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get namespace

Get a single namespace

##### Synopsis

Get a single namespace

```
glc get namespace NAMESPACE_ID [flags]
```

##### Options

```
  -h, --help   help for namespace
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project

Get a specific project

##### Synopsis

Get a specific project

```
glc get project PROJECT_ID [flags]
```

##### Options

```
  -h, --help         help for project
  -s, --statistics   Include project statistics
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-badge

Get project badge info

##### Synopsis

Get project badge info

```
glc get project-badge PROJECT_ID BADGE_ID [flags]
```

##### Options

```
  -h, --help   help for project-badge
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-branch

Get project branch info

##### Synopsis

Get project branch info

```
glc get project-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-hook

Get project hook info

##### Synopsis

Get project hook info

```
glc get project-hook PROJECT_ID HOOK_ID [flags]
```

##### Options

```
  -h, --help   help for project-hook
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-job

Get project job info

##### Synopsis

Get project job info

```
glc get project-job PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for project-job
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*
- [glc get project-job cancel](#glc-get-project-job-cancel)	*Cancel project job*
- [glc get project-job retry](#glc-get-project-job-retry)	*Retry project job*



#### glc get project-job cancel

Cancel project job

##### Synopsis

Cancel project job

```
glc get project-job cancel PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for cancel
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get project-job](#glc-get-project-job)	*Get project job info*



#### glc get project-job retry

Retry project job

##### Synopsis

Retry project job

```
glc get project-job retry PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for retry
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get project-job](#glc-get-project-job)	*Get project job info*



#### glc get project-job-trace

Get project job trace

##### Synopsis

Get project job trace

```
glc get project-job-trace PROJECT_ID JOB_ID [flags]
```

##### Options

```
  -h, --help   help for project-job-trace
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-merge-request

Get project merge request info

##### Synopsis

Get project merge request info

```
glc get project-merge-request PROJECT_ID MERGE_REQUEST_IID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-merge-request-note

Get project merge request note

##### Synopsis

Get project merge request note

```
glc get project-merge-request-note PROJECT_ID MERGE_REQUEST_IID NOTE_ID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-pipeline

Get project pipeline details

##### Synopsis

Get project pipeline details

```
glc get project-pipeline PROJECT_ID PIPELINE_ID [flags]
```

##### Options

```
  -h, --help   help for project-pipeline
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get project-var

Get the details of a project's specific variable

##### Synopsis

Get the details of a project's specific variable

```
glc get project-var PROJECT_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for project-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get runner

Get details of a runner

##### Synopsis

Get details of a runner

```
glc get runner RUNNER_ID [flags]
```

##### Options

```
  -h, --help   help for runner
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc get user

Get a single user

##### Synopsis

Get a single user

```
glc get user USER_ID [flags]
```

##### Options

```
  -h, --help   help for user
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc get](#glc-get)	*Get resource details*



#### glc init

Init glc config

##### Synopsis

Init glc config

```
glc init [flags]
```

##### Options

```
  -h, --help   help for init
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```



#### glc list

List resource

##### Synopsis

List resource

##### Options

```
  -h, --help           help for list
  -p, --page int       Page (default 1)
  -l, --per-page int   Items per page (default 10)
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list aliases](#glc-list-aliases)	*List resource aliases*
- [glc list group-epic-notes](#glc-list-group-epic-notes)	*List group epic notes*
- [glc list group-merge-requests](#glc-list-group-merge-requests)	*List group merge requests*
- [glc list group-variables](#glc-list-group-variables)	*Get list of a group's variables*
- [glc list groups](#glc-list-groups)	*List groups*
- [glc list merge-requests](#glc-list-merge-requests)	*List merge requests*
- [glc list namespaces](#glc-list-namespaces)	*List namespaces*
- [glc list project-badges](#glc-list-project-badges)	*List project badges*
- [glc list project-branches](#glc-list-project-branches)	*List project branches*
- [glc list project-commits](#glc-list-project-commits)	*List project repository commits*
- [glc list project-environments](#glc-list-project-environments)	*List project environments*
- [glc list project-hooks](#glc-list-project-hooks)	*List project's hooks*
- [glc list project-issue-notes](#glc-list-project-issue-notes)	*List project issue notes*
- [glc list project-jobs](#glc-list-project-jobs)	*List project jobs*
- [glc list project-members](#glc-list-project-members)	*List project members*
- [glc list project-merge-request-commits](#glc-list-project-merge-request-commits)	*List project merge request commits*
- [glc list project-merge-request-notes](#glc-list-project-merge-request-notes)	*List project merge request notes*
- [glc list project-merge-requests](#glc-list-project-merge-requests)	*List project merge requests*
- [glc list project-pipeline-jobs](#glc-list-project-pipeline-jobs)	*List project pipeline jobs*
- [glc list project-pipelines](#glc-list-project-pipelines)	*List project pipelines*
- [glc list project-protected-branches](#glc-list-project-protected-branches)	*List project protected branches*
- [glc list project-snippet-notes](#glc-list-project-snippet-notes)	*List project snippet notes*
- [glc list project-variables](#glc-list-project-variables)	*Get list of a project's variables*
- [glc list projects](#glc-list-projects)	*List projects*
- [glc list runners](#glc-list-runners)	*List runners*
- [glc list ssh-keys](#glc-list-ssh-keys)	*List current user ssh keys*
- [glc list user-ssh-keys](#glc-list-user-ssh-keys)	*List specific user ssh keys*
- [glc list users](#glc-list-users)	*List users*



#### glc list aliases

List resource aliases

##### Synopsis

List resource aliases

```
glc list aliases [flags]
```

##### Options

```
  -h, --help   help for aliases
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list group-epic-notes

List group epic notes

##### Synopsis

List group epic notes

```
glc list group-epic-notes GROUP_ID EPIC_ID [flags]
```

##### Options

```
  -h, --help   help for group-epic-notes
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list group-merge-requests

List group merge requests

##### Synopsis

List group merge requests

```
glc list group-merge-requests GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group-merge-requests
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list group-variables

Get list of a group's variables

##### Synopsis

Get list of a group's variables

```
glc list group-variables GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group-variables
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list groups

List groups

##### Synopsis

List groups

```
glc list groups [flags]
```

##### Options

```
      --all                      Show all the groups you have access to (defaults to false for authenticated users, true for admin)
  -h, --help                     help for groups
      --owned                    Limit to groups owned by the current user
  -s, --search string            Return the list of authorized groups matching the search criteria
      --statistics               Include group statistics (admins only)
  -x, --with-custom-attributes   Include custom attributes in response (admins only)
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list merge-requests

List merge requests

##### Synopsis

List merge requests

```
glc list merge-requests [flags]
```

##### Options

```
  -h, --help   help for merge-requests
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list namespaces

List namespaces

##### Synopsis

List namespaces

```
glc list namespaces [flags]
```

##### Options

```
  -h, --help            help for namespaces
  -s, --search string   Returns a list of namespaces the user is authorized to see based on the search criteria
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-badges

List project badges

##### Synopsis

List project badges

```
glc list project-badges PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-badges
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-branches

List project branches

##### Synopsis

List project branches

```
glc list project-branches PROJECT_ID [flags]
```

##### Options

```
  -h, --help            help for project-branches
  -s, --search string   Search term
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-commits

List project repository commits

##### Synopsis

List project repository commits

```
glc list project-commits PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-commits
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-environments

List project environments

##### Synopsis

List project environments

```
glc list project-environments PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-environments
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-hooks

List project's hooks

##### Synopsis

List project's hooks

```
glc list project-hooks PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-hooks
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-issue-notes

List project issue notes

##### Synopsis

List project issue notes

```
glc list project-issue-notes PROJECT_ID ISSUE_IID [flags]
```

##### Options

```
  -h, --help   help for project-issue-notes
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-jobs

List project jobs

##### Synopsis

List project jobs

```
glc list project-jobs PROJECT_ID [flags]
```

##### Options

```
  -h, --help           help for project-jobs
      --pretty         Use custom output formatting
  -s, --scope string   Scope
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-members

List project members

##### Synopsis

List project members

```
glc list project-members PROJECT_ID [flags]
```

##### Options

```
  -h, --help           help for project-members
  -q, --query string   Search term
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-merge-request-commits

List project merge request commits

##### Synopsis

List project merge request commits

```
glc list project-merge-request-commits PROJECT_ID MERGE_REQUEST_IID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request-commits
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-merge-request-notes

List project merge request notes

##### Synopsis

List project merge request notes

```
glc list project-merge-request-notes PROJECT_ID MERGE_REQUEST_IID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request-notes
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-merge-requests

List project merge requests

##### Synopsis

List project merge requests

```
glc list project-merge-requests PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-merge-requests
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-pipeline-jobs

List project pipeline jobs

##### Synopsis

List project pipeline jobs

```
glc list project-pipeline-jobs PROJECT_ID PIPELINE_ID [flags]
```

##### Options

```
  -h, --help           help for project-pipeline-jobs
      --pretty         Use custom output formatting
  -s, --scope string   Scope
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-pipelines

List project pipelines

##### Synopsis

List project pipelines

```
glc list project-pipelines PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-pipelines
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-protected-branches

List project protected branches

##### Synopsis

List project protected branches

```
glc list project-protected-branches PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-protected-branches
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-snippet-notes

List project snippet notes

##### Synopsis

List project snippet notes

```
glc list project-snippet-notes PROJECT_ID SNIPPET_ID [flags]
```

##### Options

```
  -h, --help   help for project-snippet-notes
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list project-variables

Get list of a project's variables

##### Synopsis

Get list of a project's variables

```
glc list project-variables PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-variables
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list projects

List projects

##### Synopsis

List projects

```
glc list projects [flags]
```

##### Options

```
      --archived        Limit by archived status
  -h, --help            help for projects
      --membership      Limit by projects that the current user is a member of
      --owned           Limit by projects owned by the current user
  -s, --search string   Search term
      --starred         Limit by projects starred by the current user
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list runners

List runners

##### Synopsis

List runners

```
glc list runners [flags]
```

##### Options

```
      --all            Get a list of all runners in the GitLab instance (specific and shared). Access is restricted to users with admin privileges
  -h, --help           help for runners
  -s, --scope string   The scope of runners to show, one of: specific, shared, active, paused, online; showing all runners if none provided
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list ssh-keys

List current user ssh keys

##### Synopsis

List current user ssh keys

```
glc list ssh-keys [flags]
```

##### Options

```
  -h, --help   help for ssh-keys
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list user-ssh-keys

List specific user ssh keys

##### Synopsis

List specific user ssh keys

```
glc list user-ssh-keys USER_ID [flags]
```

##### Options

```
  -h, --help   help for user-ssh-keys
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc list users

List users

##### Synopsis

List users

```
glc list users [flags]
```

##### Options

```
      --active            Limit to active users
      --blocked           Limit to blocked users
  -h, --help              help for users
  -s, --search string     Search users by email or username
  -u, --username string   Search users by username
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
  -p, --page int                    Page (default 1)
  -l, --per-page int                Items per page (default 10)
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc list](#glc-list)	*List resource*



#### glc rm

Remove resource

##### Synopsis

Remove resource

##### Options

```
  -h, --help   help for rm
  -y, --yes    Do not ask for confirmation
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```

##### See also

- [glc rm alias](#glc-rm-alias)	*Remove resource alias*
- [glc rm group](#glc-rm-group)	*Remove group*
- [glc rm group-epic-note](#glc-rm-group-epic-note)	*Remove group epic note*
- [glc rm group-var](#glc-rm-group-var)	*Remove a group's variable*
- [glc rm project](#glc-rm-project)	*Remove project*
- [glc rm project-badge](#glc-rm-project-badge)	*Remove project badge*
- [glc rm project-branch](#glc-rm-project-branch)	*Remove project branch*
- [glc rm project-environment](#glc-rm-project-environment)	*Remove project environment*
- [glc rm project-hook](#glc-rm-project-hook)	*Remove project hook*
- [glc rm project-issue-note](#glc-rm-project-issue-note)	*Remove project issue note*
- [glc rm project-merge-request-note](#glc-rm-project-merge-request-note)	*Remove project merge request note*
- [glc rm project-merged-branches](#glc-rm-project-merged-branches)	*Remove project merged branches*
- [glc rm project-protected-branch](#glc-rm-project-protected-branch)	*Unprotect project branch*
- [glc rm project-snippet-note](#glc-rm-project-snippet-note)	*Remove project snippet note*
- [glc rm project-star](#glc-rm-project-star)	*Unstars a given project*
- [glc rm project-var](#glc-rm-project-var)	*Remove a project's variable*



#### glc rm alias

Remove resource alias

##### Synopsis

Remove resource alias

```
glc rm alias [alias] [resource type] [flags]
```

##### Options

```
  -h, --help   help for alias
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm group

Remove group

##### Synopsis

Remove group

```
glc rm group GROUP_ID [flags]
```

##### Options

```
  -h, --help   help for group
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm group-epic-note

Remove group epic note

##### Synopsis

Remove group epic note

```
glc rm group-epic-note GROUP_ID EPIC_ID NOTE_ID [flags]
```

##### Options

```
  -h, --help   help for group-epic-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm group-var

Remove a group's variable

##### Synopsis

Remove a group's variable

```
glc rm group-var GROUP_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for group-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project

Remove project

##### Synopsis

Remove project

```
glc rm project PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-badge

Remove project badge

##### Synopsis

Remove project badge

```
glc rm project-badge PROJECT_ID BADGE_ID [flags]
```

##### Options

```
  -h, --help   help for project-badge
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-branch

Remove project branch

##### Synopsis

Remove project branch

```
glc rm project-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-environment

Remove project environment

##### Synopsis

Remove project environment

```
glc rm project-environment PROJECT_ID ENVIRONMENT_ID [flags]
```

##### Options

```
  -h, --help   help for project-environment
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-hook

Remove project hook

##### Synopsis

Remove project hook

```
glc rm project-hook PROJECT_ID HOOK_ID [flags]
```

##### Options

```
  -h, --help   help for project-hook
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-issue-note

Remove project issue note

##### Synopsis

Remove project issue note

```
glc rm project-issue-note PROJECT_ID ISSUE_IID NOTE_ID [flags]
```

##### Options

```
  -h, --help   help for project-issue-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-merge-request-note

Remove project merge request note

##### Synopsis

Remove project merge request note

```
glc rm project-merge-request-note PROJECT_ID MERGE_REQUEST_IID NOTE_ID [flags]
```

##### Options

```
  -h, --help   help for project-merge-request-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-merged-branches

Remove project merged branches

##### Synopsis

Remove project merged branches

```
glc rm project-merged-branches PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-merged-branches
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-protected-branch

Unprotect project branch

##### Synopsis

Unprotect project branch

```
glc rm project-protected-branch PROJECT_ID BRANCH_NAME [flags]
```

##### Options

```
  -h, --help   help for project-protected-branch
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-snippet-note

Remove project snippet note

##### Synopsis

Remove project snippet note

```
glc rm project-snippet-note PROJECT_ID SNIPPET_ID NOTE_ID [flags]
```

##### Options

```
  -h, --help   help for project-snippet-note
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-star

Unstars a given project

##### Synopsis

Unstars a given project

```
glc rm project-star PROJECT_ID [flags]
```

##### Options

```
  -h, --help   help for project-star
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc rm project-var

Remove a project's variable

##### Synopsis

Remove a project's variable

```
glc rm project-var PROJECT_ID VAR_KEY [flags]
```

##### Options

```
  -h, --help   help for project-var
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
  -y, --yes                         Do not ask for confirmation
```

##### See also

- [glc rm](#glc-rm)	*Remove resource*



#### glc version

Print the version number of glc

##### Synopsis

Print the version number of glc

```
glc version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
  -a, --alias string                Use resource alias
  -c, --config string               Path to configuration file (default ".glc.yml")
      --host string                 GitLab host
  -i, --interactive                 enable interactive mode when applicable (eg. creation, pagination)
      --no-color                    disable color output
  -o, --output-destination string   Output result to file if specified
  -f, --output-format string        Output format, must be one of 'text', 'json', 'yaml'
      --silent                      silent mode
  -v, --verbose                     verbose output
```




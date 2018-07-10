# Contributing to go-gitlab-client

## Requirements

- docker
- docker-compose
- make
- curl

## Setting up working environment

Please run:

```
make setup
```

This will install required tools/dependencies and launch a docker-compose stack with a go image and a wiremock server.

Then you can run:

```
make dev
```

This will automatically run tests and format code when you modify code.

### Without docker

If you do not want to work inside docker, you can also run make targets
prefixed with an underscore, like `_install_go_deps`, non prefixed targets
are just proxies to run the former inside docker compose container.

Even if you do not want to use docker for development, you'll have to launch
the docker compose stack to run tests because of wiremock.

## Building CLI binary

The following command will generate CLI binary for various platforms:

```
make cli_build_all
```

You'll find generated files in `cli/build`.

Please note that because this project use docker as a working environment,
the generated binary used for integration tests (`cli/glc`) might not work
on your host as it's generated for the golang docker image.
You can use one of the generated build in `cli/build` according to your platform.

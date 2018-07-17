# Contributing to go-gitlab-client

- [Setup](#setting-up-working-environment)
- [Build](#building-cli-binary)

## Setting up working environment

Please run:

```
make install
```

## Building CLI binary

The following command will generate CLI binary for various platforms:

```
make cli_build_all
```

You'll find generated files in `cli/build`.
You can use one of the generated build in `cli/build` according to your platform.

You can also generate checksums for the binaries:

```
make cli_checksums
```

Which will then be available in `cli/build/checksums.txt`.


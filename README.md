# go-logger

A simple facade over Go's "log" standard library.

## Features

This facade:

1. Adds level-based logging.
   1. Trace, Debug, Info, Warn, Error, Fatal, Panic
1. Adds log.Printf()-like formatting for each level.
   1. Tracef, Debugf, Infof, Warnf, Errorf, Fatalf, Panicf
1. Adds "guards" that can be used to conditionally run code
   1. IsTrace, IsDebug, IsInfo, IsWarn, IsError, IsFatal, IsPanic
1. Implemented as [fluent interface](https://en.wikipedia.org/wiki/Fluent_interface)


This facade writes to the Go `log` standard library,
so `log` and `logger` functions can be used interchangeably.

For example use, see [main.go](main.go)

## Develop

The following instructions are used when modifying and building the Docker image.

### Prerequisites for development

:thinking: The following tasks need to be complete before proceeding.
These are "one-time tasks" which may already have been completed.

1. The following software programs need to be installed:
    1. [git](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/install-git.md)
    1. [make](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/install-make.md)

### Clone repository

For more information on environment variables,
see [Environment Variables](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md).

1. Set these environment variable values:

    ```console
    export GIT_ACCOUNT=docktermj
    export GIT_REPOSITORY=go-logger
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"
    ```
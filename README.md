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

### Clone repository

1. :pencil2: Set environment variable values.
   Example:

    ```console
    export GIT_ACCOUNT=docktermj
    export GIT_REPOSITORY=go-logger
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"
    export GIT_REPOSITORY_URL="git@github.com:${GIT_ACCOUNT}/${GIT_REPOSITORY}.git"
    ```

1. Clone repository.
   Example:

    ```console
    mkdir --parents ${GIT_ACCOUNT_DIR}
    cd  ${GIT_ACCOUNT_DIR}
    git clone ${GIT_REPOSITORY_URL}
    ```

### Test

1. Get dependencies.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make dependencies
    ```

1. Run tests.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test
    ```

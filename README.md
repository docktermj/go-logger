# go-logger

A simple facade over Go's "log" standard library.

### Features

1. Adds level-based logging.
   1. Debug, Info, Warn, Error, Fatal, Panic 
1. Adds log.Printf()-like formatting for each level.
   1. Debugf, Infof, Warnf, Errorf, Fatalf, Panicf
1. Adds "guards" that can be used to conditionally run code
   1. IsDebug, IsInfo, IsWarn, IsError, IsFatal, IsPanic

This facade writes to the Go `log` standard library,
so `log` and `logger` functions can be used interchangeably.

For example use, see [main.go](main.go)
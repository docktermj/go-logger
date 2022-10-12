package logger

import (
	"fmt"
	"log"
)

// Log level constants.
type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

const (
	LevelTraceName = "TRACE"
	LevelDebugName = "DEBUG"
	LevelInfoName  = "INFO"
	LevelWarnName  = "WARN"
	LevelErrorName = "ERROR"
	LevelFatalName = "FATAL"
	LevelPanicName = "PANIC"
)

type Logger struct {
	isPanic bool
	isFatal bool
	isError bool
	isWarn  bool
	isInfo  bool
	isDebug bool
	isTrace bool
}

var logger *Logger

const noFormat = ""

func init() {
	logger = New()
}

func New() *Logger {
	return new(Logger)
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (thisLogger *Logger) printf(debugLevelName string, format string, v ...interface{}) *Logger {
	var message string
	calldepth := 3
	if format == noFormat {
		v := append(v, 0)
		copy(v[1:], v[0:])
		v[0] = debugLevelName + " "
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(debugLevelName+" "+format, v...)
	}
	log.Output(calldepth, message)
	return thisLogger
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *Logger { return logger.SetLevel(level) }
func (thisLogger *Logger) SetLevel(level Level) *Logger {
	thisLogger.isPanic = level <= LevelPanic
	thisLogger.isFatal = level <= LevelFatal
	thisLogger.isError = level <= LevelError
	thisLogger.isWarn = level <= LevelWarn
	thisLogger.isInfo = level <= LevelInfo
	thisLogger.isDebug = level <= LevelDebug
	thisLogger.isTrace = level <= LevelTrace
	return thisLogger
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return logger.IsPanic() }
func (thisLogger *Logger) IsPanic() bool {
	return thisLogger.isPanic
}

func IsFatal() bool { return logger.IsFatal() }
func (thisLogger *Logger) IsFatal() bool {
	return thisLogger.isFatal
}

func IsError() bool { return logger.IsError() }
func (thisLogger *Logger) IsError() bool {
	return thisLogger.isError
}

func IsWarn() bool { return logger.IsWarn() }
func (thisLogger *Logger) IsWarn() bool {
	return thisLogger.isWarn
}

func IsInfo() bool { return logger.IsInfo() }
func (thisLogger *Logger) IsInfo() bool {
	return thisLogger.isInfo
}

func IsDebug() bool { return logger.IsDebug() }
func (thisLogger *Logger) IsDebug() bool {
	return thisLogger.isDebug
}

func IsTrace() bool { return logger.IsTrace() }
func (thisLogger *Logger) IsTrace() bool {
	return thisLogger.isTrace
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Trace ------------------------------------------------------------------

func Trace(v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func (thisLogger *Logger) Trace(v ...interface{}) *Logger {
	if thisLogger.isTrace {
		thisLogger.printf(LevelTraceName, noFormat, v...)
	}
	return thisLogger
}

func Tracef(format string, v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

func (thisLogger *Logger) Tracef(format string, v ...interface{}) *Logger {
	if thisLogger.isTrace {
		thisLogger.printf(LevelTraceName, format, v...)
	}
	return thisLogger
}

// --- Debug ------------------------------------------------------------------

func Debug(v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func (thisLogger *Logger) Debug(v ...interface{}) *Logger {
	if thisLogger.isDebug {
		thisLogger.printf(LevelDebugName, noFormat, v...)
	}
	return thisLogger
}

func Debugf(format string, v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

func (thisLogger *Logger) Debugf(format string, v ...interface{}) *Logger {
	if thisLogger.isDebug {
		thisLogger.printf(LevelDebugName, format, v...)
	}
	return thisLogger
}

// --- Info -------------------------------------------------------------------

func Info(v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func (thisLogger *Logger) Info(v ...interface{}) *Logger {
	if thisLogger.isInfo {
		thisLogger.printf(LevelInfoName, noFormat, v...)
	}
	return thisLogger
}

func Infof(format string, v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

func (thisLogger *Logger) Infof(format string, v ...interface{}) *Logger {
	if thisLogger.isInfo {
		thisLogger.printf(LevelInfoName, format, v...)
	}
	return thisLogger
}

// --- Warn -------------------------------------------------------------------

func Warn(v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func (thisLogger *Logger) Warn(v ...interface{}) *Logger {
	if thisLogger.isWarn {
		thisLogger.printf(LevelWarnName, noFormat, v...)
	}
	return thisLogger
}

func Warnf(format string, v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

func (thisLogger *Logger) Warnf(format string, v ...interface{}) *Logger {
	if thisLogger.isWarn {
		thisLogger.printf(LevelWarnName, format, v...)
	}
	return thisLogger
}

// --- Error ------------------------------------------------------------------

func Error(v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func (thisLogger *Logger) Error(v ...interface{}) *Logger {
	if thisLogger.isError {
		thisLogger.printf(LevelErrorName, noFormat, v...)
	}
	return thisLogger
}

func Errorf(format string, v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

func (thisLogger *Logger) Errorf(format string, v ...interface{}) *Logger {
	if thisLogger.isError {
		thisLogger.printf(LevelErrorName, format, v...)
	}
	return thisLogger
}

// --- Fatal ------------------------------------------------------------------

func Fatal(v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func (thisLogger *Logger) Fatal(v ...interface{}) *Logger {
	if thisLogger.isFatal {
		thisLogger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return thisLogger
}

func Fatalf(format string, v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

func (thisLogger *Logger) Fatalf(format string, v ...interface{}) *Logger {
	if thisLogger.isFatal {
		thisLogger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return thisLogger
}

// --- Panic ------------------------------------------------------------------

func Panic(v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func (thisLogger *Logger) Panic(v ...interface{}) *Logger {
	if thisLogger.isPanic {
		thisLogger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return thisLogger
}

func Panicf(format string, v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

func (thisLogger *Logger) Panicf(format string, v ...interface{}) *Logger {
	if thisLogger.isPanic {
		thisLogger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return thisLogger
}

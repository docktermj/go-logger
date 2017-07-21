package logger

import (
	"fmt"
	"log"
)

// Log level constants.
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

const (
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
}

var logger *Logger

func init() {
	logger = New()
}

func New() *Logger {
	return new(Logger)
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (this *Logger) printf(debugLevelName string, format string, v ...interface{}) *Logger {
	calldepth := 3
	message := fmt.Sprintf(debugLevelName+" "+format, v...)
	log.Output(calldepth, message)
	return this
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *Logger { return logger.SetLevel(level) }
func (this *Logger) SetLevel(level Level) *Logger {
	this.isPanic = level <= LevelPanic
	this.isFatal = level <= LevelFatal
	this.isError = level <= LevelError
	this.isWarn = level <= LevelWarn
	this.isInfo = level <= LevelInfo
	this.isDebug = level <= LevelDebug
	return this
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return logger.IsPanic() }
func (this *Logger) IsPanic() bool {
	return this.isPanic
}

func IsFatal() bool { return logger.IsFatal() }
func (this *Logger) IsFatal() bool {
	return this.isFatal
}

func IsError() bool { return logger.IsError() }
func (this *Logger) IsError() bool {
	return this.isError
}

func IsWarn() bool { return logger.IsWarn() }
func (this *Logger) IsWarn() bool {
	return this.isWarn
}

func IsInfo() bool { return logger.IsInfo() }
func (this *Logger) IsInfo() bool {
	return this.isInfo
}

func IsDebug() bool { return logger.IsDebug() }
func (this *Logger) IsDebug() bool {
	return this.isDebug
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Debug ------------------------------------------------------------------

func Debug(message string) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, "%s", message)
	}
	return logger
}

func (this *Logger) Debug(message string) *Logger {
	if this.isDebug {
		this.printf(LevelDebugName, "%s", message)
	}
	return this
}

func Debugf(format string, v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

func (this *Logger) Debugf(format string, v ...interface{}) *Logger {
	if this.isDebug {
		this.printf(LevelDebugName, format, v...)
	}
	return this
}

// --- Info -------------------------------------------------------------------

func Info(message string) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, "%s", message)
	}
	return logger
}

func (this *Logger) Info(message string) *Logger {
	if this.isInfo {
		this.printf(LevelInfoName, "%s", message)
	}
	return this
}

func Infof(format string, v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

func (this *Logger) Infof(format string, v ...interface{}) *Logger {
	if this.isInfo {
		this.printf(LevelDebugName, format, v...)
	}
	return this
}

// --- Warn -------------------------------------------------------------------

func Warn(message string) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, "%s", message)
	}
	return logger
}

func (this *Logger) Warn(message string) *Logger {
	if this.isWarn {
		this.printf(LevelWarnName, "%s", message)
	}
	return this
}

func Warnf(format string, v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

func (this *Logger) Warnf(format string, v ...interface{}) *Logger {
	if this.isWarn {
		this.printf(LevelDebugName, format, v...)
	}
	return this
}

// --- Error ------------------------------------------------------------------

func Error(message string) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, "%s", message)
	}
	return logger
}

func (this *Logger) Error(message string) *Logger {
	if this.isError {
		this.printf(LevelErrorName, "%s", message)
	}
	return this
}

func Errorf(format string, v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

func (this *Logger) Errorf(format string, v ...interface{}) *Logger {
	if this.isError {
		this.printf(LevelErrorName, format, v...)
	}
	return this
}

// --- Fatal ------------------------------------------------------------------

func Fatal(message string) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, "%s", message)
		log.Fatal("")
	}
	return logger
}

func (this *Logger) Fatal(message string) *Logger {
	if this.isFatal {
		this.printf(LevelFatalName, "%s", message)
		log.Fatal("")
	}
	return this
}

func Fatalf(format string, v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

func (this *Logger) Fatalf(format string, v ...interface{}) *Logger {
	if this.isFatal {
		this.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return this
}

// --- Panic ------------------------------------------------------------------

func Panic(message string) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, "%s", message)
		log.Panic("")
	}
	return logger
}

func (this *Logger) Panic(message string) *Logger {
	if this.isPanic {
		this.printf(LevelPanicName, "%s", message)
		log.Panic("")
	}
	return this
}

func Panicf(format string, v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

func (this *Logger) Panicf(format string, v ...interface{}) *Logger {
	if this.isPanic {
		this.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return this
}
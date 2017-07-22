package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestLevels(test *testing.T) {
	assert.True(test, LevelDebug < LevelInfo, "Debug")
	assert.True(test, LevelInfo < LevelWarn, "Info")
	assert.True(test, LevelWarn < LevelError, "Warn")
	assert.True(test, LevelError < LevelFatal, "Error")
	assert.True(test, LevelFatal < LevelPanic, "Fatal")
}

func TestIsDebug(test *testing.T) {
	SetLevel(LevelDebug)
	assert.True(test, IsDebug(), "Debug")
	assert.True(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

func TestIsInfo(test *testing.T) {
	SetLevel(LevelInfo)
	assert.False(test, IsDebug(), "Debug")
	assert.True(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

func TestIsWarn(test *testing.T) {
	SetLevel(LevelWarn)
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

func TestIsError(test *testing.T) {
	SetLevel(LevelError)
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.False(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

func TestIsFatal(test *testing.T) {
	SetLevel(LevelFatal)
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.False(test, IsWarn(), "Warn")
	assert.False(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

func TestIsPanic(test *testing.T) {
	SetLevel(LevelPanic)
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.False(test, IsWarn(), "Warn")
	assert.False(test, IsError(), "Error")
	assert.False(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

func TestDebug(test *testing.T) {
	SetLevel(LevelDebug)
	assert.NotZero(test, Debug("test"), "string")
	assert.NotZero(test, Debugf("test %s", "something"), "format")
}

func TestInfo(test *testing.T) {
	SetLevel(LevelInfo)
	assert.NotZero(test, Info("test"), "string")
	assert.NotZero(test, Infof("test %s", "something"), "format")
}

func TestWarn(test *testing.T) {
	SetLevel(LevelWarn)
	assert.NotZero(test, Warn("test"), "string")
	assert.NotZero(test, Warnf("test %s", "something"), "format")
}

func TestError(test *testing.T) {
	SetLevel(LevelError)
	assert.NotZero(test, Error("test"), "string")
	assert.NotZero(test, Errorf("test %s", "something"), "format")
}

func TestFluentInterface(test *testing.T) {
	SetLevel(LevelDebug)
	Debug("debug").Info("info").Warn("warn").Error("error")
}

func TestVaradic(test *testing.T) {
	SetLevel(LevelDebug)
	_, err := time.LoadLocation("bob")
	Info("Should be error: ", err)
}

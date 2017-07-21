package main

import (
	"log"
	"time"

	"github.com/docktermj/go-logger/logger"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func complexProcess() string {
	time.Sleep(10 * time.Second)
	return "slept"
}

func main() {

	//	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)

	logger.SetLevel(logger.LevelInfo)

	log.Println("Test Debug")
	logger.Debug("debug works")
	logger.Debugf("debug A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Info")
	logger.Info("info works")
	logger.Infof("info A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Warn")
	logger.Warn("warn works")
	logger.Warnf("warn A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Error")
	logger.Error("error works")
	logger.Errorf("error A: %s B: %s C: %d", "aaa", "bbb", 35)

	// Test the efficiency of the logger
	log.Println("Start timed test")
	log.Printf("%s", complexProcess())
	log.Println("1")
	if logger.IsDebug() {
		logger.Debugf("%s", complexProcess())
	}
	log.Println("Done")

	// Note:  the first Fatal or Panic issued will exit the program.

	log.Println("Test Fatal")
	//	logger.Fatal("fatal works")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Panicl")
	//	logger.Fatal("fatal works")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("End")

}

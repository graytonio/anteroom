package lib

import (
	"io/ioutil"
	"log"
	"os"
)

var error_logger *log.Logger
var debug_logger *log.Logger
var status_logger *log.Logger

func init() {
	error_logger = log.New(os.Stderr, "", 0)
	debug_logger = log.New(ioutil.Discard, "", log.LstdFlags)
	status_logger = log.New(os.Stdout, "", 0)
}

func EnableDebugLogger() {
	debug_logger.SetOutput(os.Stdout)
}

func LogStatus(msg string) {
	status_logger.Println(msg)
}

func LogError(msg string) {
	error_logger.Println(msg)
}

func LogDebug(msg string) {
	debug_logger.Println(msg)
}

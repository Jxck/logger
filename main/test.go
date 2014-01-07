package main

import (
	"github.com/jxck/logger"
)

func init() {
	logger.LogLevel(4)
}

func main() {
	logger.Error("%v", "error")
	logger.Warn("%v", "warn")
	logger.Info("%v", "info")
	logger.Debug("%v", "logger")
}

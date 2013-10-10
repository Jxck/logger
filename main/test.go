package main

import (
	"github.com/jxck/logger"
)

func init() {
	logger.Init(4)
}

func main() {
	logger.Info("%v", "info")
	logger.Debug("%v", "logger")
	logger.Warn("%v", "warn")
	logger.Error("%v", "error")
}

package main

import (
	"github.com/jxck/debug"
)

func init() {
	debug.Init(3)
}

func main() {
	debug.Info("%v", "info")
	debug.Debug("%v", "debug")
	debug.Warn("%v", "warn")
	debug.Error("%v", "error")
}

package main

import (
	"github.com/jxck/debug"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	Log := debug.NewLogger(4)
	Log.Info("%v", "info")
	Log.Debug("%v", "debug")
	Log.Warn("%v", "warn")
	Log.Error("%v", "error")
}

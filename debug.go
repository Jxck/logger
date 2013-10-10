package debug

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	// 1 ERR
	// 2 WARNING
	// 3 INFO
	// 4 DEBUG
	LogLevel int
}

var Level map[int]string = map[int]string{
	1: "ERROR",
	2: "WARN",
	3: "INFO",
	4: "DEBUG",
}

func NewLogger(level int) *Logger {
	return &Logger{
		LogLevel: level,
	}
}

func getPath() (dir, path string, line int) {
	_, fullpath, line, _ := runtime.Caller(3)
	f := strings.Split(fullpath, "/")
	dir = f[len(f)-2]
	file := f[len(f)-1]
	path = dir + "/" + file
	return
}

func out(level int, format string, a ...interface{}) {
	_, path, line := getPath()

	dest := os.Stdout
	if level >= 4 {
		dest = os.Stderr
	}

	fmt.Fprintf(dest,
		"%s [%v:%v] %v\n",
		Level[level],
		path,
		line,
		fmt.Sprintf(format, a...))
}

func (l *Logger) Error(format string, a ...interface{}) {
	if l.LogLevel >= 1 {
		out(1, format, a...)
	}
}

func (l *Logger) Warn(format string, a ...interface{}) {
	if l.LogLevel >= 2 {
		out(2, format, a...)
	}
}

func (l *Logger) Info(format string, a ...interface{}) {
	if l.LogLevel >= 3 {
		out(3, format, a...)
	}
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if l.LogLevel >= 4 {
		out(4, format, a...)
	}
}

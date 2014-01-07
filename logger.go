package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var logger *Logger

type Logger struct {
	// 1 ERR
	// 2 WARNING
	// 3 INFO
	// 4 DEBUG
	Level int
}

var Level map[int]string = map[int]string{
	1: "ERROR",
	2: "WARN",
	3: "INFO",
	4: "DEBUG",
}

func NewLogger(level int) *Logger {
	return &Logger{
		Level: level,
	}
}

func LogLevel(level int) {
	logger = NewLogger(level)
}

func getPath() (dir, path string, line int) {
	_, fullpath, line, _ := runtime.Caller(4)
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
	if l.Level >= 1 {
		out(1, format, a...)
	}
}

func (l *Logger) Warn(format string, a ...interface{}) {
	if l.Level >= 2 {
		out(2, format, a...)
	}
}

func (l *Logger) Info(format string, a ...interface{}) {
	if l.Level >= 3 {
		out(3, format, a...)
	}
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if l.Level >= 4 {
		out(4, format, a...)
	}
}

func Error(format string, a ...interface{}) {
	logger.Error(format, a...)
}

func Warn(format string, a ...interface{}) {
	logger.Warn(format, a...)
}

func Info(format string, a ...interface{}) {
	logger.Info(format, a...)
}

func Debug(format string, a ...interface{}) {
	logger.Debug(format, a...)
}

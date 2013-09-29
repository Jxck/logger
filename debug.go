package http2

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Debug(format string, a ...interface{}) {
	env := os.Getenv("DEBUG")

	_, fullpath, line, _ := runtime.Caller(1)
	f := strings.Split(fullpath, "/")
	dir := f[len(f)-2]
	file := f[len(f)-1]
	path := dir + "/" + file

	if strings.Contains(env, dir) {
		fmt.Printf("[%v:%v] ", path, line)
		fmt.Printf(format, a...)
		fmt.Printf("\n")
	}
}

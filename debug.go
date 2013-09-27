package http2

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Debug(str string) {
	env := os.Getenv("DEBUG")
	if strings.Contains(env, "http2") {
		_, file, line, _ := runtime.Caller(1)
		f := strings.Split(file, "/")
		filename := f[len(f)-2] + "/" + f[len(f)-1]
		fmt.Printf("%v:%v %v\n", filename, line, str)
	}
}

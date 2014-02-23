# logger

logger for github.com/jxck/http2

## log level

```
1 ERR
2 WARN
3 NOTICE
4 INFO
5 DEBUG
6 TRACE
```

## Sample

```go
package main

import (
	"flag"
	. "github.com/jxck/logger"
)

var loglevel int

func init() {
	flag.IntVar(&loglevel, "l", 0, "log level (1 ERR, 2 WARNING, 3 INFO, 4 DEBUG)")
	flag.Parse()
	LogLevel(loglevel)
}

func main() {
	Error("%v", "error")
	Warn("%v", "warn")
	Notice("%v", "notice")
	Info("%v", "info")
	Debug("%v", "logger")
}
```

## License

The MIT License (MIT)
Copyright (c) 2013 Jxck

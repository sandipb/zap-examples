# Using global loggers

zap also seems to offer global loggers - `zap.L()` for the mandatory structure logger, and `zap.S()`, for the _Sugared_ logger.

From what I have seen, these loggers are not meant to be used out of the box (like `log.Print()` in the standard library), but rather their purpose seems only to provide a shared logger throughout the code. If you really want to use it, you need to [replace](https://godoc.org/go.uber.org/zap#ReplaceGlobals) the core with that of a different logger. You are also provided a way to _undo_ a replacement. Out of the box, the global loggers have no output.

```console
$ go run src/globallogger/main.go

*** Using the global logger out of the box


*** After replacing the global logger with a development logger

2018-05-02T16:24:40.992-0700    INFO    globallogger/main.go:17 An info message {"iteration": 1}

*** After replacing the global logger with a production logger

{"level":"info","ts":1525303480.993161,"caller":"globallogger/main.go:22","msg":"An info message","iteration":1}

*** After undoing the last replacement of the global logger

2018-05-02T16:24:40.993-0700    INFO    globallogger/main.go:26 An info message {"iteration": 1}
```


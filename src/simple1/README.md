# Using presets

Zap recommends using presets for the simplest of cases. 

It makes three presets available:

- Example
- Development
- Production

I try out their presets in this piece of code. Here is the output.

```console
$ go run src/simple1/main.go

*** Using the Example logger

{"level":"debug","msg":"This is a DEBUG message"}
{"level":"info","msg":"This is an INFO message"}
{"level":"info","msg":"This is an INFO message with fields","region":"us-west","id":2}
{"level":"warn","msg":"This is a WARN message"}
{"level":"error","msg":"This is an ERROR message"}
{"level":"dpanic","msg":"This is a DPANIC message"}

*** Using the Development logger

2018-05-02T13:52:44.332-0700    DEBUG   simple1/main.go:28      This is a DEBUG message
2018-05-02T13:52:44.332-0700    INFO    simple1/main.go:29      This is an INFO message
2018-05-02T13:52:44.332-0700    INFO    simple1/main.go:30      This is an INFO message with fields     {"region": "us-west", "id": 2}
2018-05-02T13:52:44.332-0700    WARN    simple1/main.go:31      This is a WARN messagemain.main
        /Users/snbhatta/dev/zap-examples/src/simple1/main.go:31
runtime.main
        /Users/snbhatta/.gradle/language/golang/1.9.2/go/src/runtime/proc.go:195
2018-05-02T13:52:44.332-0700    ERROR   simple1/main.go:32      This is an ERROR message
main.main
        /Users/snbhatta/dev/zap-examples/src/simple1/main.go:32
runtime.main
        /Users/snbhatta/.gradle/language/golang/1.9.2/go/src/runtime/proc.go:195

*** Using the Production logger

{"level":"info","ts":1525294364.332839,"caller":"simple1/main.go:43","msg":"This is an INFO message"}
{"level":"info","ts":1525294364.332864,"caller":"simple1/main.go:44","msg":"This is an INFO message with fields","region":"us-west","id":2}
{"level":"warn","ts":1525294364.3328729,"caller":"simple1/main.go:45","msg":"This is a WARN message"}
{"level":"error","ts":1525294364.332882,"caller":"simple1/main.go:46","msg":"This is an ERROR message","stacktrace":"main.main\n\t/Users/snbhatta/dev/zap-examples/src/simple1/main.go:46\nruntime.main\n\
t/Users/snbhatta/.gradle/language/golang/1.9.2/go/src/runtime/proc.go:195"}
{"level":"dpanic","ts":1525294364.332895,"caller":"simple1/main.go:48","msg":"This is a DPANIC message","stacktrace":"main.main\n\t/Users/snbhatta/dev/zap-examples/src/simple1/main.go:48\nruntime.main\n
\t/Users/snbhatta/.gradle/language/golang/1.9.2/go/src/runtime/proc.go:195"}


*** Using the Sugar logger

2018-05-02T18:13:22.376-0700    INFO    simple1/main.go:56      Info() uses sprint
2018-05-02T18:13:22.376-0700    INFO    simple1/main.go:57      Infof() uses sprintf
2018-05-02T18:13:22.376-0700    INFO    simple1/main.go:58      Infow() allows tags     {"name": "Legolas", "type": 1}
```

# Observations

- Both `Example` and `Production` loggers use the [JSON encoder](https://godoc.org/go.uber.org/zap/zapcore#NewJSONEncoder). `Development` uses the [Console](https://godoc.org/go.uber.org/zap/zapcore#NewConsoleEncoder) encoder.
- The `logger.DPanic()` function causes a panic in `Development` logger but not in `Example` or `Production`.
- The `Development` logger:
    * Adds a stack trace from Warn level and up. 
    * Always prints the package/file/line number
    * Tacks extra fields as a json string at the end of the line
    * level names are uppercase
    * timestamp is in ISO8601 with seconds
- The `Production` logger:
    * Doesn't log messages at debug level
    * Adds stack trace as a json field for Error, DPanic levels, but not for Warn.
    * Always adds the caller as a json field
    * timestamp is in epoch format
    * level is in lower case

# Using the "sugar" logger

The default logger expects structured tags.

```go
logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
```

This is the fastest option for an application where performance is key.

However, for a just [a small additional penalty](https://github.com/uber-go/zap#performance), 
which actually is still slightly better than the standard library, you can use 
the _sugar_ logger, which uses a reflection based type detection to give you
a simpler syntax to add tags of mixed types.

```go
slogger := logger.Sugar()
slogger.Info("Info() uses sprint")
slogger.Infof("Infof() uses %s", "sprintf")
slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)
```

Output:

```
2018-05-02T18:13:22.376-0700    INFO    simple1/main.go:56      Info() uses sprint
2018-05-02T18:13:22.376-0700    INFO    simple1/main.go:57      Infof() uses sprintf
2018-05-02T18:13:22.376-0700    INFO    simple1/main.go:58      Infow() allows tags     {"name": "Legolas", "type": 1}
```

You can switch from a sugar logger to a standard logger any time using the 
`.Desugar()` method on the logger.
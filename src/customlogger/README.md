# Creating custom loggers

Instead of using the presets, you can create custom loggers with the exact 
combinations of features that you want.

# Using the zap config struct to create a logger

Loggers can be created using a configuration struct `zap.Config`. You are expected 
to fill in the struct with require values, and then call the `.Build()` method
on the struct to get your logger.

```go
cfg := zap.Config{...}
logger, err := cfg.Build()
```

There are no sane defaults for the struct. You have to at the minimum provide 
values for the three classes of settings that zap needs.

* _encoder_: Just adding a `Encoding: "xxx"` field is a minimum. Using `json` 
   here as the value will create a default JSON encoder. You can customize the 
   encoder (which almost certainly you have to, because the defaults aren't very 
   useful), by adding a `zapcore.EncoderConfig` struct to the `EncoderConfig` 
   field.
* _level enabler_: This is a data type which allows zap to determine whether a 
   message at a particular level should be displayed. In the zap config struct, 
   you provide such a type using the `AtomicLevel` wrapper in the `Level` field.
* _sink_: This is the destination of the log messages. You can specify multiple
   output paths using the `OutputPaths` field which accepts a list of path names.
   Magic values like `"stderr"` and `"stdout"` can be used for the usual 
   purposes.
  
# Customizing the encoder

Just mentioning an encoder type in the struct is not enough. By default the
JSON encoder only outputs fields specifically provided in the log messages.

```go
logger, _ = zap.Config{
    Encoding:    "json",
    Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
    OutputPaths: []string{"stdout"},
}.Build()

logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
```

Will output:

```
{"region":"us-west","id":2}
```

Even the message is not printed!

To add the message in the JSON encoder, you need to specify the JSON key name
which will display this value in the output.

```go
logger, _ = zap.Config{
    Encoding:    "json",
    Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
    OutputPaths: []string{"stdout"},
    EncoderConfig: zapcore.EncoderConfig{
        MessageKey: "message",
    },
}.Build()

logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
```

Will output:

```
{"message":"This is an INFO message with fields","region":"us-west","id":2}
```

zap can add more metadata to the message like level name, timestamp, caller,
stacktrace, etc. Unless you specifically mention the JSON key in the output 
corresponding to these metadata, it is not added.

Some of these field names *have* to be paired with an _encoder_ else zap just
burns and dies (!!).

For example:

```go
cfg := zap.Config{
    Encoding:         "json",
    Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
    OutputPaths:      []string{"stderr"},
    ErrorOutputPaths: []string{"stderr"},
    EncoderConfig: zapcore.EncoderConfig{
        MessageKey: "message",

        LevelKey:    "level",
        EncodeLevel: zapcore.CapitalLevelEncoder,

        TimeKey:    "time",
        EncodeTime: zapcore.ISO8601TimeEncoder,

        CallerKey:    "caller",
        EncodeCaller: zapcore.ShortCallerEncoder,
    },
}
logger, _ = cfg.Build()
logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
```

Will output:

```
{"level":"INFO","time":"2018-05-02T16:37:54.998-0700","caller":"customlogger/main.go:91","message":"This is an INFO message with fields","region":"us-west","id":2}
```

Each of the encoder can be customized to fit your requirements, and some have
different implementations provided by zap. 

- timestamp can be output in either ISO 8601 format, or as an epoch timestamp.
- level can be capital or lowercase or even colored (even though it is probably 
  only visible in the console output). Weirdly, the colors escape codes are 
  not stripped in the JSON output.
- caller can be shown in short and full formats.

# Changing logger behavior on the fly

loggers can be cloned from an existing logger with certain modification to their
behavior. This can often be useful for example, when you want to reduce code 
duplication by fixing a standard set of fields the logger will always output.

* `logger.AddCaller()` adds caller annotation
* `logger.AddStacktrace()` adds stacktraces for messages at and above a given 
  level
* `logger.Fields()` adds specified fields to all messages output by the new logger
* `logger.WrapCore()` allows you to modify or even completely replace the 
  underlying _core_ in the logger which combines the encoder, level and sink.

```go
fmt.Printf("\n*** Using a JSON encoder, at debug level, sending output to stdout, all possible keys specified\n\n")

cfg := zap.Config{
    Encoding:         "json",
    Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
    OutputPaths:      []string{"stderr"},
    ErrorOutputPaths: []string{"stderr"},
    EncoderConfig: zapcore.EncoderConfig{
        MessageKey: "message",

        LevelKey:    "level",
        EncodeLevel: zapcore.CapitalLevelEncoder,

        TimeKey:    "time",
        EncodeTime: zapcore.ISO8601TimeEncoder,

        CallerKey:    "caller",
        EncodeCaller: zapcore.ShortCallerEncoder,
    },
}
logger, _ = cfg.Build()

logger.Info("This is an INFO message")

fmt.Printf("\n*** Same logger with console logging enabled instead\n\n")

logger.WithOptions(
    zap.WrapCore(
        func(zapcore.Core) zapcore.Core {
            return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
        })).Info("This is an INFO message")

```

Output:

```
*** Using a JSON encoder, at debug level, sending output to stdout, all possible keys specified

{"level":"INFO","time":"2018-05-02T16:37:54.998-0700","caller":"customlogger/main.go:90","message":"This is an INFO message"}

*** Same logger with console logging enabled instead

2018-05-02T16:37:54.998-0700    INFO    customlogger/main.go:99 This is an INFO message

```
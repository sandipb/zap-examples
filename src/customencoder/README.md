# Using custom encoders for standard fields

You can use custom encoders for formatting time, level, caller, etc. One caveat
is that you need these encoders to be as efficient as possible so as to not
negate the memory/speed advantages of zap itself. After all, these functions 
are called for *every* line of log to be emitted!

That said, the examples below are just a demonstration. I am not claiming at all 
that they are efficient replacements of standard functionality. :)

## Customizing timestamp formatting

Here is an implementation which uses the syslog format often found in the wild.

```go
func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

...

cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder

logger, _ = cfg.Build()
logger.Info("This should have a syslog style timestamp")
```

Output:

```
May  2 18:54:55 INFO    This should have a syslog style timestamp
```

If you noticed in the implementation that the encoder is supposed to append
primitives to an array like object. zap uses this array to efficiently encode
output with minimal memory allocations.

## Customizing level formatting

```go
func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
...
cfg.EncoderConfig.EncodeLevel = CustomLevelEncoder

logger, _ = cfg.Build()
logger.Info("This should have a bracketed level name")
```

Output:

```
May  2 18:54:55 [INFO]  This should have a bracketed level name
```

*NOTE*: I am creating a single string from multiple substrings and appending 
it to the array. For some reason, if I appended the substrings separately, zap
is padding them with spaces.

Similar customization can be done for other metadata fields as well. You can
look [at the zap source](https://sourcegraph.com/github.com/uber-go/zap/-/blob/zapcore/encoder.go)
to find the general pattern of these implementations.
# Hierarchal Logging

One way of simulating hierarchal logging is to rely of Cloning the parent logger and 
replacing the settings you want to change and get the old reference point to the new logger.
The cloning process is never inplace. 

The keys are usage of zap.WrapCore, the introduced Clone and the Logger struct to wrap the underlying zap logger.

Hierarachal logging can be using when you have multiple components that require differnt default logging
levels.

When look at the examples, keep note of which ancestor loggers each children is cloned from.
Also, notice the extra wrapping required to retain the underlying zap logger semantics.
You can definitely want to investigate if wrapping Levels is required for more manipulations of Levels.

```console
$ go run src/hierarchal-logging/main.go

*** Create multiple hierachy of loggers on raw logger ***
You probably need to wrap SetLevel to deal with more level setting/resetting
parent child loggers are really copying parent logger then replace settings
{"level":"debug","ts":1527539781.3316631,"msg":"childlogger2 decends from parentLogger with Debug enabled. Debug Should print"}
{"level":"debug","ts":1527539781.3317158,"msg":"This child still inherits level from childLogger2. Debug Should print \n"}

```
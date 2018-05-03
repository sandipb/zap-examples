# Using the zap logging library

This repository provides some examples of using [Uber's zap](https://github.com/uber-go/zap) Go logging library

Install the zap library before trying out the examples:

```console
$ source env.sh

$ go get -u go.uber.org/zap

$ go run src/simple1/main.go
```

## Examples

* [Simplest usage using presets](./src/simple1)
* [Creating a custom logger](./src/customlogger)
* [Using the global logger](./src/globallogger)
* [Creating custom encoders for metadata fields](./src/customencoder)
package main

import (
	"fmt"
    "encoding/json"

	"go.uber.org/zap"
)

func main() {
	fmt.Printf("\n*** Using the Example logger\n\n")

	logger := zap.NewExample()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")  // would exit if uncommented
	logger.DPanic("This is a DPANIC message")
	//logger.Panic("This is a PANIC message")   // would exit if uncommented

	fmt.Println()

	fmt.Printf("*** Using the Development logger\n\n")

	logger, _ = zap.NewDevelopment()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")   // would exit if uncommented
	// logger.DPanic("This is a DPANIC message") // would exit if uncommented
	//logger.Panic("This is a PANIC message")    // would exit if uncommented

	fmt.Println()

	fmt.Printf("*** Using the Production logger\n\n")

	logger, _ = zap.NewProduction()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")   // would exit if uncommented
	logger.DPanic("This is a DPANIC message")
	// logger.Panic("This is a PANIC message")   // would exit if uncommented

	fmt.Println()

	fmt.Printf("*** Using the Sugar logger\n\n")

	logger, _ = zap.NewDevelopment()
	slogger := logger.Sugar()
	slogger.Info("Info() uses sprint")
	slogger.Infof("Infof() uses %s", "sprintf")
	slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)

    fmt.Printf("*** Build a logger from a json")
    rawJSONConfig := []byte(`{
      "level": "debug",
      "encoding": "console",
      "outputPaths": ["stdout", "/tmp/logs"],
      "errorOutputPaths": ["stderr"],
      "initialFields": {"foo": "bar"},
      "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "nameKey": "logger",
        "timeKey": "time",  
        "callerKey": "logger",
        "stacktraceKey": "stacktrace",          
        "callstackKey": "callstack",
        "errorKey": "error",
        "fileKey": "file",
        "levelEncoder": "capitalColor",
        "timeEncoder": "iso8601",
        "durationEncoder": "nanos",
        "callerEncoder": "full",
        "nameEncoder": "full",
        "sampling": {
            "initial": "5",
            "thereafter": "10"
        }
      }
    }`)
    config := zap.Config{}
    if err := json.Unmarshal(rawJSONConfig, &config); err != nil {
        panic(err)
    }
    logger, err := config.Build()
    if err != nil {
        panic(err)
    }
    logger.Debug("This is a DEBUG message")
    logger.Info("This should have an ISO8601 based time stamp")
    logger.Warn("This is a WARN message")
    logger.Error("This is an ERROR message")
    // logger.Fatal("This is a FATAL message")   // would exit if uncommented
    // logger.DPanic("This is a DPANIC message") // would exit if uncommented


}

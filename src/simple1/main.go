package main

import (
	"fmt"

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
}

package main

import (
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func main() {

	/*

		// This would cause a panic: "panic: no encoder name specified"

		cfg := zap.Config{}
		logger, err := cfg.Build()
		if err != nil {
			panic(err)
		}
	*/

	/*

		// causes a panic: "panic: runtime error: invalid memory address or nil pointer dereference"

		logger, err := zap.Config{Encoding: "json"}.Build()
	*/

	/*

		// No output

		logger, err := zap.Config{Encoding: "json", Level: zap.NewAtomicLevelAt(zapcore.InfoLevel)}.Build()

	*/

	fmt.Printf("\n*** Using a JSON encoder, at debug level, sending output to stdout, no key specified\n\n")

	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

	fmt.Printf("\n*** Using a JSON encoder, at debug level, sending output to stdout, message key only specified\n\n")

	logger, _ = zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
		},
	}.Build()

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

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

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

	fmt.Printf("\n*** Same logger with console logging enabled instead\n\n")

	logger.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
				return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
			})).Info("This is an INFO message")
}

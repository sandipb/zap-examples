package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func main() {

	cfg := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}

	fmt.Printf("\n*** Using standard ISO8601 time encoder\n\n")

	// avoiding copying of atomic values
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logger, _ := cfg.Build()
	logger.Info("This should have an ISO8601 based time stamp")

	fmt.Printf("\n*** Using a custom time encoder\n\n")

	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder

	logger, _ = cfg.Build()
	logger.Info("This should have a syslog style time stamp")

	fmt.Printf("\n*** Using a custom level encoder\n\n")

	cfg.EncoderConfig.EncodeLevel = CustomLevelEncoder

	logger, _ = cfg.Build()
	logger.Info("This should have a interesting level name")

}

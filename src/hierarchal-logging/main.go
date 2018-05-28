package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level
type Logger struct {
	logger  *zap.Logger
	encoder zapcore.Encoder
	writer  zapcore.WriteSyncer
}

func main() {
	fmt.Printf("\n*** Create multiple hierachy of loggers on raw logger *** \n")
	fmt.Printf("You probably need to wrap SetLevel to deal with more level setting/resetting\n")

	encoderConfig := zap.NewProductionEncoderConfig()
	atomLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(encoder, writer, &atomLevel)
	logger := zap.New(core)
	parentLogger := Logger{logger, encoder, writer}
	parentLogger.Debug("Should not print")
	fmt.Printf("parent child loggers are really copying parent logger then replace settings \n")
	childLogger1 := parentLogger.Clone()
	childLogger1.Debug("This child still inherits level from parent logger. Should not print \n")
	childLogger2 := parentLogger.NewLevel(zapcore.DebugLevel)
	childLogger2.Debug("childlogger2 decends from parentLogger with Debug enabled. Debug Should print")
	childLogger3 := childLogger2.Clone()
	childLogger3.Debug("This child still inherits level from childLogger2. Debug Should print \n")
    childLogger4 := parentLogger.Clone()
    childLogger4.Debug("This child still inherits level from original parentLogger. Debug Should not print \n") 
    // {"level":"debug","ts":1527539781.3316631,"msg":"childlogger2 decends from parentLogger with Debug enabled. Debug Should print"}
    // {"level":"debug","ts":1527539781.3317158,"msg":"This child still inherits level from childLogger2. Debug Should print \n"}
}

func (l *Logger) NewLevel(level zapcore.Level) *Logger {
	newLevel := zapcore.Level(level)
	newLogger := l.logger.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
				return zapcore.NewCore(l.encoder, l.writer, newLevel)
			}))
	return &Logger{newLogger, l.encoder, l.writer}
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *Logger) Clone() *Logger {
	copy := *l
	return &copy
}

package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func main() {
	fmt.Printf("Log sampling to reduce the presure on I/O and CPU by reducing logging\n")
	fmt.Printf("Using the built in sampler. You probably need to wrap the whole zapcore Sampler public methods if you need to write our own custom sampler\n")
	encoderConfig := zap.NewProductionEncoderConfig()
	atomLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(encoder, writer, &atomLevel)
	emitInitialMessages := 5
	thereAfterEachMessages := 100
	// You can embed zap.Logger inside your Logger struct for WithSamplingConfig and preserve zap.Logger interface
	// Or if you take zap.Logger as one of the parameters for WithSamplingConfig, you don't need to use embedded struct.
	logger := &Logger{zap.New(core)}

	fmt.Printf("We will first emit the first %v messages then one every each %v messages thereafter \n", emitInitialMessages, thereAfterEachMessages)
	logger = logger.WithSamplingConfig(time.Second, emitInitialMessages, thereAfterEachMessages)

	for i := 1; i < 110; i++ {
		logger.With(zap.Int("n", i)).Info("test at info")
	}
	logger.Sync()
	// Output:
	// {"level":"info","ts":1527544371.1515696,"msg":"test at info","n":1}
	// {"level":"info","ts":1527544371.1515965,"msg":"test at info","n":2}
	// {"level":"info","ts":1527544371.1516032,"msg":"test at info","n":3}
	// {"level":"info","ts":1527544371.1516066,"msg":"test at info","n":4}
	// {"level":"info","ts":1527544371.15161,"msg":"test at info","n":5}
	// {"level":"info","ts":1527544371.1518133,"msg":"test at info","n":105}
}

func (l *Logger) WithSamplingConfig(tick time.Duration, initial, thereAfter int) *Logger {
	if initial < 1 || thereAfter < 1 {
		// fmt.Printf("all arguments must be positive")
		return l
	}
	core := l.Core()
	newLogger := l.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
				return zapcore.NewSampler(core, tick, initial, thereAfter)
			}))
	return &Logger{newLogger}
}

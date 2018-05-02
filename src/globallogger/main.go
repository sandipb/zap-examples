package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	fmt.Printf("\n*** Using the global logger out of the box\n\n")

	zap.S().Infow("An info message", "iteration", 1)

	fmt.Printf("\n*** After replacing the global logger with a development logger\n\n")
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infow("An info message", "iteration", 1)

	fmt.Printf("\n*** After replacing the global logger with a production logger\n\n")
	logger, _ = zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	zap.S().Infow("An info message", "iteration", 1)

	fmt.Printf("\n*** After undoing the last replacement of the global logger\n\n")
	undo()
	zap.S().Infow("An info message", "iteration", 1)

}

package main

import (
	"go.uber.org/zap"
	"os"
	"reminder/internal/engine"
)

func init() {
	initLog()
}

func main() {
	// init routes and start listening
	if err := engine.NewRoutes().Start("0.0.0.0:8080"); err != nil {
		zap.L().Sugar().Errorf("internal server error, detail: %v", err)
		return
	}
	// listen to the signal and if signal received then quit
	go func() {
		sigChan := make(chan os.Signal, 1)
		if sig := <-sigChan; sig != nil {
			os.Exit(0)
		}
	}()
}

func initLog() {
	loggerConfig := zap.NewProductionConfig()
	logger, _ := loggerConfig.Build()
	zap.ReplaceGlobals(logger)
}

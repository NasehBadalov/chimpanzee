package log

import "go.uber.org/zap"

// InitLogger initializes zap logger to log into stdout and sets global logger
func InitLogger() error {
	// Configure logging
	// log to stdout and set to global logger

	lconf := zap.NewProductionConfig()
	lconf.OutputPaths = []string{"stdout"}
	logger, err := lconf.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}

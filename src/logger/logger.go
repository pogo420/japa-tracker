package logger

import (
	"japa-tracker/src/config"
	"log/slog"
	"os"
)

// Initializing logger
func InitializeLogger() *slog.Logger {

	// Setting log level
	var appLogLevel slog.Level
	switch config.AppConfig.LogLevel {
	case config.Info:
		appLogLevel = slog.LevelInfo
	case config.Debug:
		appLogLevel = slog.LevelDebug
	default:
		appLogLevel = slog.LevelInfo
	}

	appLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: appLogLevel,
	}))
	// Setting application default
	slog.SetDefault(appLogger)
	return appLogger
}

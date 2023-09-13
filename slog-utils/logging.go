package slog_utils

import (
	"log/slog"
	"os"
	"strings"
)

const (
	LevelTrace slog.Level = slog.LevelDebug - 4
)

var (
	programLevel = new(slog.LevelVar)
	rootLogger   *slog.Logger
)

func GetRootLogger() *slog.Logger {
	return rootLogger
}

func GetLoggerForCurrentPackage() *slog.Logger {
	if rootLogger == nil {
		return nil
	}
	callInfo := retrieveCallInfo()
	return rootLogger.With("pkg", callInfo.packageName)
}

func ToLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "TRACE":
		return LevelTrace
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func SetLevel(level slog.Level) {
	programLevel.Set(level)
}

func init() {
	rootLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: programLevel,
	}))
}

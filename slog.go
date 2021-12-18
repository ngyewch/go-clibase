package clibase

import (
	slog "github.com/go-eden/slf4go"
	"strings"
)

func ToSlogLevel(levelName string, defaultLevel slog.Level) slog.Level {
	levelName = strings.ToUpper(levelName)
	switch levelName {
	case "FATAL":
		return slog.FatalLevel
	case "PANIC":
		return slog.PanicLevel
	case "ERROR":
		return slog.ErrorLevel
	case "WARN":
		return slog.WarnLevel
	case "INFO":
		return slog.InfoLevel
	case "DEBUG":
		return slog.DebugLevel
	case "TRACE":
		return slog.TraceLevel
	default:
		return defaultLevel
	}
}

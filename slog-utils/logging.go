package slog_utils

import (
	"io"
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
	writer       = &replaceableWriter{
		w: os.Stdout,
	}
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

func SetWriter(w io.Writer) {
	writer.w = w
}

func init() {
	rootLogger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: programLevel,
	}))
}

type replaceableWriter struct {
	w io.Writer
}

func (w *replaceableWriter) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

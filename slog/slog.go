package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// Logger is a slog wrapper that implements the logx.Logger interface
type Logger struct {
	*slog.Logger
}

// pathExists checks if the path exists
func pathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, fmt.Errorf("file exists")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// NewLog creates a slog logger instance
func NewLog(c *SlogConfig) *Logger {
	// if the log directory does not exist, create it
	if c.Dir != "" {
		if ok, _ := pathExists(c.Dir); !ok {
			_ = os.MkdirAll(c.Dir, os.ModePerm)
		}
	}

	handler, err := GetHandler(c)
	if err != nil {
		panic(err)
	}

	logger := slog.New(handler)
	return &Logger{Logger: logger}
}

// getSlogLevel converts the string level to slog.Level
func getSlogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "fatal":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Debug implements the Debug method of the Logger interface
func (l *Logger) Debug(args ...any) {
	l.Logger.Debug(args[0].(string), args[1:]...)
}

// Info implements the Info method of the Logger interface
func (l *Logger) Info(args ...any) {
	l.Logger.Info(args[0].(string), args[1:]...)
}

// Warn implements the Warn method of the Logger interface
func (l *Logger) Warn(args ...any) {
	l.Logger.Warn(args[0].(string), args[1:]...)
}

// Error implements the Error method of the Logger interface
func (l *Logger) Error(args ...any) {
	l.Logger.Error(args[0].(string), args[1:]...)
}

// Fatal implements the Fatal method of the Logger interface
func (l *Logger) Fatal(args ...any) {
	l.Logger.Error(args[0].(string), args[1:]...)
	os.Exit(1)
}

// WithContext returns a logger with context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{Logger: l.Logger}
}

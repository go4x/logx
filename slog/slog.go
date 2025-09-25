// Package slog provides a slog logger implementation for the logx package.
// It supports Go's standard structured logging with features like log rotation,
// buffering, and performance optimization.
package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// Logger wraps slog.Logger to implement the logx.Logger interface.
type Logger struct {
	*slog.Logger
}

// pathExists checks if the given path exists.
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

// NewLog creates a new slog logger instance with the given configuration.
// Returns nil if configuration is invalid or handler creation fails.
func NewLog(c *SlogConfig) (*Logger, error) {
	if c.Dir == "" {
		return nil, fmt.Errorf("dir is required")
	}

	if ok, _ := pathExists(c.Dir); !ok {
		_ = os.MkdirAll(c.Dir, os.ModePerm)
	}

	handler, err := GetHandler(c)
	if err != nil {
		return nil, err
	}

	logger := slog.New(handler)
	return &Logger{Logger: logger}, nil
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
	if len(args) == 0 {
		return
	}
	msg := fmt.Sprint(args...)
	l.Logger.Debug(msg)
}

// Info implements the Info method of the Logger interface
func (l *Logger) Info(args ...any) {
	if len(args) == 0 {
		return
	}
	msg := fmt.Sprint(args...)
	l.Logger.Info(msg)
}

// Warn implements the Warn method of the Logger interface
func (l *Logger) Warn(args ...any) {
	if len(args) == 0 {
		return
	}
	msg := fmt.Sprint(args...)
	l.Logger.Warn(msg)
}

// Error implements the Error method of the Logger interface
func (l *Logger) Error(args ...any) {
	if len(args) == 0 {
		return
	}
	msg := fmt.Sprint(args...)
	l.Logger.Error(msg)
}

// Fatal implements the Fatal method of the Logger interface
// Note: This implementation logs the message as error and exits.
// Consider using a different approach for production code.
func (l *Logger) Fatal(args ...any) {
	if len(args) == 0 {
		l.Logger.Error("Fatal error occurred")
		os.Exit(1)
		return
	}
	msg := fmt.Sprint(args...)
	l.Logger.Error(msg)
	os.Exit(1)
}

// WithContext returns a logger with context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{Logger: l.Logger}
}

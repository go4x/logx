// Package logx provides a unified logging interface with support for multiple backends.
// It supports both zap and slog loggers with features like log rotation, buffering,
// and performance optimization.
//
// Example usage:
//
//	config := &logx.LoggerConfig{
//	    Type:         logx.LoggerTypeZap,
//	    Level:        "info",
//	    LogInConsole: true,
//	    Dir:          "logs",
//	    Format:       "json",
//	}
//	err := logx.Init(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	logx.Info("Hello, world!")
package logx

import (
	"errors"
	"fmt"

	"github.com/go4x/logx/slog"
	"github.com/go4x/logx/zap"
)

// LogFormat represents the format of log output.
type LogFormat string

const (
	// FormatText outputs logs in plain text format.
	FormatText LogFormat = "text"
	// FormatJSON outputs logs in JSON format.
	FormatJSON LogFormat = "json"
)

// LoggerType represents the type of logger backend.
type LoggerType string

const (
	// LoggerTypeSlog uses Go's standard slog logger.
	LoggerTypeSlog LoggerType = "slog"
	// LoggerTypeZap uses Uber's zap logger for high performance.
	LoggerTypeZap LoggerType = "zap"
)

// Logger defines the interface for logging operations.
// All logger implementations must satisfy this interface.
type Logger interface {
	// Debug logs a debug message.
	Debug(args ...any)
	// Info logs an informational message.
	Info(args ...any)
	// Warn logs a warning message.
	Warn(args ...any)
	// Error logs an error message.
	Error(args ...any)
	// Fatal logs a fatal message and exits the program.
	Fatal(args ...any)
}

// LoggerConfig holds the configuration for the logger.
type LoggerConfig struct {
	// Type specifies the logger backend type (slog or zap).
	Type LoggerType `mapstructure:"type" yaml:"type"`

	// Level specifies the minimum log level (debug, info, warn, error, fatal).
	Level string `mapstructure:"level" yaml:"level"`

	// LogInConsole determines whether to output logs to console.
	LogInConsole bool `mapstructure:"log-in-console" yaml:"log-in-console"`

	// Dir specifies the directory where log files are stored.
	Dir string `mapstructure:"dir" yaml:"dir"`

	// Format specifies the log output format (text or json).
	Format string `mapstructure:"format" yaml:"format"`

	// MaxAge specifies the maximum number of days to retain log files (0 means no limit).
	MaxAge int `mapstructure:"max-age" yaml:"max-age"`

	// MaxSize specifies the maximum size of a log file in MB (0 means no limit).
	MaxSize int `mapstructure:"max-size" yaml:"max-size"`

	// MaxBackups specifies the maximum number of log files to keep (0 means no limit).
	MaxBackups int `mapstructure:"max-backups" yaml:"max-backups"`

	// LocalTime determines whether to use local time (true) or UTC time (false).
	LocalTime bool `mapstructure:"local-time" yaml:"local-time"`

	// Compress determines whether to compress rotated log files.
	Compress bool `mapstructure:"compress" yaml:"compress"`

	// BufferSize specifies the buffer size for performance optimization (0 disables buffering).
	BufferSize int `mapstructure:"buffer-size" yaml:"buffer-size"`

	// FlushInterval specifies the interval in seconds to flush the buffer.
	FlushInterval int `mapstructure:"flush-interval" yaml:"flush-interval"`
}

// globalLogger is the global logger instance.
var globalLogger Logger

// Init initializes the logger according to the configuration.
// It returns an error if the configuration is invalid or initialization fails.
func Init(c *LoggerConfig) error {
	if c == nil {
		return errors.New("logger config cannot be nil")
	}

	// if the type is not specified, default to zap
	if c.Type == "" {
		c.Type = LoggerTypeZap
	}

	switch c.Type {
	case LoggerTypeSlog:
		return initSlogLogger(c)
	case LoggerTypeZap:
		return initZapLogger(c)
	default:
		return fmt.Errorf("unsupported logger type: %s", c.Type)
	}
}

// GetLogger returns the global logger instance.
func GetLogger() Logger {
	return globalLogger
}

// initSlogLogger initializes the slog logger with the given configuration.
func initSlogLogger(c *LoggerConfig) error {
	// convert LoggerConfig to SlogConfig
	slogConfig := &slog.SlogConfig{
		Level:         c.Level,
		LogInFile:     c.Dir != "", // Enable file logging if directory is specified
		LogInConsole:  c.LogInConsole,
		Dir:           c.Dir,
		Format:        c.Format,
		MaxAge:        c.MaxAge,
		MaxSize:       c.MaxSize,
		MaxBackups:    c.MaxBackups,
		LocalTime:     c.LocalTime,
		Compress:      c.Compress,
		BufferSize:    c.BufferSize,    // Use value from configuration
		FlushInterval: c.FlushInterval, // Use value from configuration
	}

	// create the slog logger
	logger := slog.NewLog(slogConfig)
	globalLogger = logger
	return nil
}

// initZapLogger initializes the zap logger with the given configuration.
func initZapLogger(c *LoggerConfig) error {
	// convert LoggerConfig to ZapConfig
	zapConfig := &zap.ZapConfig{
		Level:         c.Level,
		Format:        c.Format,
		Director:      c.Dir,
		MaxAge:        c.MaxAge,
		MaxSize:       c.MaxSize,
		MaxBackups:    c.MaxBackups,
		LogInConsole:  c.LogInConsole,
		ShowCaller:    true,
		LocalTime:     c.LocalTime,
		Compress:      c.Compress,
		StacktraceKey: "stacktrace",
		BufferSize:    c.BufferSize,    // Use value from configuration
		FlushInterval: c.FlushInterval, // Use value from configuration
	}

	// create the zap logger
	logger := zap.NewLog(zapConfig)
	globalLogger = logger
	return nil
}

// Debug logs a debug message using the global logger.
// If the global logger is not initialized, this function does nothing.
func Debug(args ...any) {
	if globalLogger != nil {
		globalLogger.Debug(args...)
	}
}

// Info logs an informational message using the global logger.
// If the global logger is not initialized, this function does nothing.
func Info(args ...any) {
	if globalLogger != nil {
		globalLogger.Info(args...)
	}
}

// Warn logs a warning message using the global logger.
// If the global logger is not initialized, this function does nothing.
func Warn(args ...any) {
	if globalLogger != nil {
		globalLogger.Warn(args...)
	}
}

// Error logs an error message using the global logger.
// If the global logger is not initialized, this function does nothing.
func Error(args ...any) {
	if globalLogger != nil {
		globalLogger.Error(args...)
	}
}

// Fatal logs a fatal message using the global logger and exits the program.
// If the global logger is not initialized, this function does nothing.
func Fatal(args ...any) {
	if globalLogger != nil {
		globalLogger.Fatal(args...)
	}
}

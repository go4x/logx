// Package logx is a log package
package logx

import (
	"errors"
	"fmt"

	"github.com/gophero/logx/slog"
	"github.com/gophero/logx/zap"
)

type LogFormat string

const (
	FormatText LogFormat = "text"
	FormatJSON LogFormat = "json"
)

// LoggerType 定义日志器类型
type LoggerType string

const (
	LoggerTypeSlog LoggerType = "slog"
	LoggerTypeZap  LoggerType = "zap"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type LoggerConfig struct {
	// logger type, slog or zap
	Type LoggerType `mapstructure:"type" yaml:"type"`

	// log level, debug, info, warn, error, fatal
	Level string `mapstructure:"level" yaml:"level"`

	// whether to write to console, default is false
	LogInConsole bool `mapstructure:"log-in-console" yaml:"log-in-console"`

	// log directory where the log file is stored
	Dir string `mapstructure:"dir" yaml:"dir"`

	// log format, text or json
	Format string `mapstructure:"format" yaml:"format"`

	// log retention time in days, 0 means no limit
	MaxAge int `mapstructure:"max-age" yaml:"max-age"`

	// log file size in MB, 0 means no limit
	MaxSize int `mapstructure:"max-size" yaml:"max-size"`

	// maximum number of log files to keep, 0 means no limit
	MaxBackups int `mapstructure:"max-backups" yaml:"max-backups"`

	// local time, default is true, if false, use UTC time
	LocalTime bool `mapstructure:"local-time" yaml:"local-time"`

	// compress, default is false
	Compress bool `mapstructure:"compress" yaml:"compress"`
}

// globalLogger is the global logger instance
var globalLogger Logger

// Init initialize the logger according to the configuration
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

// GetLogger get the global logger instance
func GetLogger() Logger {
	return globalLogger
}

// initSlogLogger initialize the slog logger
func initSlogLogger(c *LoggerConfig) error {
	// convert LoggerConfig to SlogConfig
	slogConfig := &slog.SlogConfig{
		Level:        c.Level,
		LogInConsole: c.LogInConsole,
		Dir:          c.Dir,
		Format:       c.Format,
		MaxAge:       c.MaxAge,
		MaxSize:      c.MaxSize,
		MaxBackups:   c.MaxBackups,
	}

	// create the slog logger
	logger := slog.NewLog(slogConfig)
	globalLogger = logger
	return nil
}

// initZapLogger initialize the zap logger
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
		LocalTime:     true,
		Compress:      false,
		StacktraceKey: "stacktrace",
	}

	// create the zap logger
	logger := zap.NewLog(zapConfig)
	globalLogger = logger
	return nil
}

// Debug is a wrapper for the global logger's Debug method
func Debug(args ...any) {
	globalLogger.Debug(args...)
}

// Info is a wrapper for the global logger's Info method
func Info(args ...any) {
	globalLogger.Info(args...)
}

// Warn is a wrapper for the global logger's Warn method
func Warn(args ...any) {
	globalLogger.Warn(args...)
}

// Error is a wrapper for the global logger's Error method
func Error(args ...any) {
	globalLogger.Error(args...)
}

// Fatal is a wrapper for the global logger's Fatal method
func Fatal(args ...any) {
	globalLogger.Fatal(args...)
}

// Package zap provides configuration structures and utilities for the zap logger.
package zap

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

const (
	// ZapFormatJSON specifies JSON format for log output.
	ZapFormatJSON = "json"
	// ZapFormatText specifies plain text format for log output.
	ZapFormatText = "text"

	// ZapEncodeLevelLower encodes log levels in lowercase.
	ZapEncodeLevelLower = "lower"
	// ZapEncodeLevelLowerColor encodes log levels in lowercase with color.
	ZapEncodeLevelLowerColor = "lowerColor"
	// ZapEncodeLevelCap encodes log levels in uppercase.
	ZapEncodeLevelCap = "cap"
	// ZapEncodeLevelCapColor encodes log levels in uppercase with color.
	ZapEncodeLevelCapColor = "capColor"
)

// ZapConfig holds the configuration for the zap logger.
type ZapConfig struct {
	// Level specifies the minimum log level (debug, info, warn, error, fatal).
	Level string `mapstructure:"level" yaml:"level"`
	// Prefix specifies the log message prefix.
	Prefix string `mapstructure:"prefix" yaml:"prefix"`
	// Format specifies the log output format (json or text).
	Format string `mapstructure:"format" yaml:"format"`
	// Director specifies the directory where log files are stored.
	Director string `mapstructure:"director" yaml:"director"`
	// EncodeLevel specifies how to encode log levels.
	EncodeLevel string `mapstructure:"encode-level" yaml:"encode-level"`
	// StacktraceKey specifies the key for stacktrace in log output.
	StacktraceKey string `mapstructure:"stacktrace-key" yaml:"stacktrace-key"`
	// MaxAge specifies the maximum number of days to retain log files.
	MaxAge int `mapstructure:"max-age" yaml:"max-age"`
	// MaxSize specifies the maximum size of a log file in MB.
	MaxSize int `mapstructure:"max-size" yaml:"max-size"`
	// MaxBackups specifies the maximum number of log files to keep.
	MaxBackups int `mapstructure:"max-backups" yaml:"max-backups"`
	// ShowCaller determines whether to show the caller information.
	ShowCaller bool `mapstructure:"show-caller" yaml:"show-caller"`
	// LogInConsole determines whether to output logs to console.
	LogInConsole bool `mapstructure:"log-in-console" yaml:"log-in-console"`
	// LocalTime determines whether to use local time (true) or UTC time (false).
	LocalTime bool `mapstructure:"local-time" yaml:"local-time"`
	// Compress determines whether to compress rotated log files.
	Compress bool `mapstructure:"compress" yaml:"compress"`
	// BufferSize specifies the buffer size for performance optimization.
	BufferSize int `mapstructure:"buffer-size" yaml:"buffer-size"`
	// FlushInterval specifies the interval in seconds to flush the buffer.
	FlushInterval int `mapstructure:"flush-interval" yaml:"flush-interval"`
}

// ZapEncodeLevel get zapcore.LevelEncoder
func (z *ZapConfig) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == ZapEncodeLevelLower: // lower case encoder
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == ZapEncodeLevelLowerColor: // lower case encoder with color
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == ZapEncodeLevelCap: // capital encoder
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == ZapEncodeLevelCapColor: // capital encoder with color
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.CapitalLevelEncoder
	}
}

// TransportLevel get zapcore.Level
func (z *ZapConfig) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

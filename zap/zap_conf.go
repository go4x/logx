package zap

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

const (
	ZapFormatJSON = "json"
	ZapFormatText = "text"

	ZapEncodeLevelLower      = "lower"
	ZapEncodeLevelLowerColor = "lowerColor"
	ZapEncodeLevelCap        = "cap"
	ZapEncodeLevelCapColor   = "capColor"
)

type ZapConfig struct {
	Level         string `mapstructure:"level" yaml:"level"`                   // 级别
	Prefix        string `mapstructure:"prefix" yaml:"prefix"`                 // 日志前缀
	Format        string `mapstructure:"format" yaml:"format"`                 // 输出，json或text
	Director      string `mapstructure:"director" yaml:"director"`             // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" yaml:"encode-level"`     // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	MaxAge        int    `mapstructure:"max-age" yaml:"max-age"`               // 日志留存时间
	MaxSize       int    `mapstructure:"max-size" yaml:"max-size"`             // 日志文件大小
	MaxBackups    int    `mapstructure:"max-backups" yaml:"max-backups"`       // 最大备份数量
	ShowCaller    bool   `mapstructure:"show-caller" yaml:"show-caller"`       // 显示行
	LogInConsole  bool   `mapstructure:"log-in-console" yaml:"log-in-console"` // 输出控制台
	LocalTime     bool   `mapstructure:"local-time" yaml:"local-time"`         // 本地时间
	Compress      bool   `mapstructure:"compress" yaml:"compress"`             // 是否压缩
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
		return zapcore.WarnLevel
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

// Package zap provides a zap logger implementation for the logx package.
// It supports high-performance logging with features like log rotation,
// buffering, and structured logging.
package zap

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Default is the default zap logger instance.
var Default = NewLog(&ZapConfig{Director: "logs"})

type loggerKey string

// LoggerKey is the context key for storing the logger.
const LoggerKey loggerKey = "zapLogger"

// Logger wraps zap.SugaredLogger to implement the logx.Logger interface.
type Logger struct {
	*zap.SugaredLogger
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

// NewLog creates a new zap logger instance with the given configuration.
func NewLog(c *ZapConfig) *Logger {
	zapObj = zapDef{c: c}

	// if the log directory does not exist, create it
	if ok, _ := pathExists(c.Director); !ok {
		fmt.Printf("create %v directory\n", c.Director)
		_ = os.Mkdir(c.Director, os.ModePerm)
	}

	cores := zapObj.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if c.ShowCaller {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return &Logger{SugaredLogger: logger.Sugar()}
}

var zapObj zapDef

// zapDef holds the zap logger definition.
type zapDef struct {
	c *ZapConfig
}

// GetEncoder returns a zapcore.Encoder based on the configuration.
func (z *zapDef) GetEncoder() zapcore.Encoder {
	switch z.c.Format {
	case ZapFormatJSON:
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	default:
		return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
	}
}

// GetEncoderConfig get zapcore.EncoderConfig
func (z *zapDef) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "log",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  z.c.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.c.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder, // Log time
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GetEncoderCore get zapcore.Core
func (z *zapDef) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := GetWriter(z.c, l.String()) // use file-rotatelogs to split logs
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder custom log output time format
func (z *zapDef) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	var prefix string
	if z.c.Prefix != "" {
		prefix = "[" + z.c.Prefix + "] "
	}
	encoder.AppendString(prefix + t.Format("2006/01/02 15:04:05.000"))
}

// GetZapCores get []zapcore.Core
func (z *zapDef) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.c.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority get zap.LevelEnablerFunc
func (z *zapDef) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // debug level
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // info level
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // warn level
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // error level
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic level
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic level
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // fatal level
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // debug level
			return level == zap.DebugLevel
		}
	}
}

// NewContext add fields to the specified context
func (l *Logger) NewContext(ctx context.Context, fields ...any) context.Context {
	return context.WithValue(ctx, LoggerKey, l.WithContext(ctx).With(fields...))
}

// WithContext return a zap instance from the specified context
func (l *Logger) WithContext(ctx context.Context) *Logger {
	if ctx == nil {
		return l
	}
	zl := ctx.Value(LoggerKey)
	ctxLogger, ok := zl.(*zap.SugaredLogger)
	if ok {
		return &Logger{ctxLogger}
	}
	return l
}

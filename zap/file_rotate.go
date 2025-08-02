package zap

import (
	"os"
	"path"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GetWriter get io.Writer
func GetWriter(c *ZapConfig, level string) (zapcore.WriteSyncer, error) {
	now := time.Now().Format("2006-01-02")
	filename := path.Join(c.Director, now+"-"+level+".log")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		LocalTime:  c.LocalTime,
		MaxSize:    c.MaxSize, //mb
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}

	if c.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger)), nil
	}
	return zapcore.AddSync(lumberjackLogger), nil
}

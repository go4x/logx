package slog

import (
	"io"
	"log/slog"
	"os"
	"path"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// GetHandler get slog.Handler
func GetHandler(c *SlogConfig) (slog.Handler, error) {
	now := time.Now().Format("2006-01-02")
	filename := path.Join(c.Dir, now+"-"+c.Level+".log")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		LocalTime:  c.LocalTime,
		MaxSize:    c.MaxSize, //mb
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}
	var writer io.Writer
	if c.LogInConsole {
		writer = io.MultiWriter(os.Stdout, lumberjackLogger)
	} else {
		writer = lumberjackLogger
	}
	switch c.Format {
	case "json":
		return slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: getSlogLevel(c.Level),
		}), nil
	default:
		return slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: getSlogLevel(c.Level),
		}), nil
	}
}

package slog

import (
	"log/slog"
	"testing"
)

func TestSlog(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("debug")
	slog.Debug("debug", "a", "b")
	slog.Debug("", "a", "b")
	slog.Info("info")
	slog.Info("info", "a", "b")
	slog.Warn("warn")
	slog.Warn("warn", "a", "b")
	slog.Error("error")
	slog.Error("error", "a", "b")
}

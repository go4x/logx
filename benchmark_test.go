package logx_test

import (
	"os"
	"testing"

	"github.com/go4x/logx"
)

// BenchmarkZapLogger benchmarks zap logger performance
func BenchmarkZapLogger(b *testing.B) {
	logDir := "benchmark_logs_zap"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		b.Fatalf("failed to initialize logger: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logx.Info("benchmark message")
		}
	})
}

// BenchmarkSlogLogger benchmarks slog logger performance
func BenchmarkSlogLogger(b *testing.B) {
	logDir := "benchmark_logs_slog"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeSlog,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		b.Fatalf("failed to initialize logger: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logx.Info("benchmark message")
		}
	})
}

// BenchmarkZapLoggerConsole benchmarks zap logger with console output
func BenchmarkZapLoggerConsole(b *testing.B) {
	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: true,
		Format:       "json",
	}

	err := logx.Init(cfg)
	if err != nil {
		b.Fatalf("failed to initialize logger: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logx.Info("benchmark message")
		}
	})
}

// BenchmarkSlogLoggerConsole benchmarks slog logger with console output
func BenchmarkSlogLoggerConsole(b *testing.B) {
	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeSlog,
		Level:        "info",
		LogInConsole: true,
		Format:       "json",
	}

	err := logx.Init(cfg)
	if err != nil {
		b.Fatalf("failed to initialize logger: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logx.Info("benchmark message")
		}
	})
}

// BenchmarkDifferentLevels benchmarks different log levels
func BenchmarkDifferentLevels(b *testing.B) {
	logDir := "benchmark_logs_levels"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "debug",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		b.Fatalf("failed to initialize logger: %v", err)
	}

	b.Run("Debug", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logx.Debug("debug message")
		}
	})

	b.Run("Info", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logx.Info("info message")
		}
	})

	b.Run("Warn", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logx.Warn("warn message")
		}
	})

	b.Run("Error", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logx.Error("error message")
		}
	})
}

// BenchmarkMixedArguments benchmarks logging with mixed argument types
func BenchmarkMixedArguments(b *testing.B) {
	logDir := "benchmark_logs_mixed"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		b.Fatalf("failed to initialize logger: %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logx.Info("message", 123, true, "string", nil)
		}
	})
}

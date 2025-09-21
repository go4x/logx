package logx_test

import (
	"os"
	"testing"
	"time"

	"github.com/go4x/logx"
)

// TestBufferConfiguration tests buffer configuration
func TestBufferConfiguration(t *testing.T) {
	logDir := "testlogs_buffer"
	defer os.RemoveAll(logDir)

	// Test zap with buffer configuration
	t.Run("ZapBuffer", func(t *testing.T) {
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
			t.Fatalf("failed to initialize zap logger: %v", err)
		}

		// Log some messages
		for i := 0; i < 10; i++ {
			logx.Info("buffer test message", "index", i)
		}

		// Wait a bit to ensure buffering works
		time.Sleep(100 * time.Millisecond)

		// Check if log file is created
		files, err := os.ReadDir(logDir)
		if err != nil {
			t.Fatalf("failed to read log directory: %v", err)
		}
		if len(files) == 0 {
			t.Errorf("no log files found in directory: %s", logDir)
		}
	})

	// Test slog with buffer configuration
	t.Run("SlogBuffer", func(t *testing.T) {
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
			t.Fatalf("failed to initialize slog logger: %v", err)
		}

		// Log some messages
		for i := 0; i < 10; i++ {
			logx.Info("buffer test message", "index", i)
		}

		// Wait a bit to ensure buffering works
		time.Sleep(100 * time.Millisecond)

		// Check if log file is created
		files, err := os.ReadDir(logDir)
		if err != nil {
			t.Fatalf("failed to read log directory: %v", err)
		}
		if len(files) == 0 {
			t.Errorf("no log files found in directory: %s", logDir)
		}
	})
}

// TestBufferPerformance tests buffer performance
func TestBufferPerformance(t *testing.T) {
	logDir := "testlogs_buffer_perf"
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
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Test performance with buffering
	start := time.Now()
	for i := 0; i < 1000; i++ {
		logx.Info("performance test", "index", i, "timestamp", time.Now())
	}
	elapsed := time.Since(start)

	t.Logf("Logged 1000 messages in %v", elapsed)
	t.Logf("Average time per message: %v", elapsed/1000)

	// Wait for buffer flush
	time.Sleep(2 * time.Second)

	// Check if log file is created
	files, err := os.ReadDir(logDir)
	if err != nil {
		t.Fatalf("failed to read log directory: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("no log files found in directory: %s", logDir)
	}
}

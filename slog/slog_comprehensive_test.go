package slog_test

import (
	"os"
	"testing"
	"time"

	"github.com/go4x/logx/slog"
)

// TestSlogConfigValidation tests slog configuration validation
func TestSlogConfigValidation(t *testing.T) {
	testCases := []struct {
		name   string
		config *slog.SlogConfig
		valid  bool
	}{
		{
			name: "ValidConfig",
			config: &slog.SlogConfig{
				Level:        "info",
				Format:       "json",
				Dir:          "testlogs",
				MaxAge:       7,
				MaxSize:      100,
				MaxBackups:   10,
				LogInConsole: true,
				LogInFile:    true,
				LocalTime:    true,
				Compress:     true,
			},
			valid: true,
		},
		{
			name:   "EmptyConfig",
			config: &slog.SlogConfig{},
			valid:  true, // Should work with defaults
		},
		{
			name: "MinimalConfig",
			config: &slog.SlogConfig{
				Level: "debug",
				Dir:   "testlogs",
			},
			valid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer os.RemoveAll("testlogs")

			logger := slog.NewLog(tc.config)
			if logger == nil {
				t.Error("expected logger to be created")
			}
		})
	}
}

// TestSlogFormats tests different output formats
func TestSlogFormats(t *testing.T) {
	defer os.RemoveAll("testlogs")

	formats := []string{"json", "text", ""}

	for _, format := range formats {
		t.Run("Format_"+format, func(t *testing.T) {
			config := &slog.SlogConfig{
				Level:        format,
				Format:       format,
				Dir:          "testlogs",
				LogInConsole: true,
			}

			logger := slog.NewLog(config)
			if logger == nil {
				t.Error("expected logger to be created")
			}

			logger.Info("format test", "format", format)
		})
	}
}

// TestSlogLogLevels tests different log levels
func TestSlogLogLevels(t *testing.T) {
	defer os.RemoveAll("testlogs")

	levels := []string{"debug", "info", "warn", "error"}

	for _, level := range levels {
		t.Run("Level_"+level, func(t *testing.T) {
			config := &slog.SlogConfig{
				Level:        level,
				Format:       "text",
				Dir:          "testlogs",
				LogInConsole: true,
			}

			logger := slog.NewLog(config)
			if logger == nil {
				t.Error("expected logger to be created")
			}

			// Test all log levels
			logger.Debug("debug message")
			logger.Info("info message")
			logger.Warn("warn message")
			logger.Error("error message")
		})
	}
}

// TestSlogConsoleAndFileOutput tests both console and file output
func TestSlogConsoleAndFileOutput(t *testing.T) {
	logDir := "testlogs_console_file"
	defer os.RemoveAll(logDir)

	config := &slog.SlogConfig{
		Level:        "info",
		Format:       "text",
		Dir:          logDir,
		LogInConsole: true,
		LogInFile:    true,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	logger.Info("console and file output test")
	time.Sleep(100 * time.Millisecond)

	// Check if log file is created
	files, err := os.ReadDir(logDir)
	if err != nil {
		t.Fatalf("failed to read log directory: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("no log files found in directory: %s", logDir)
	}
}

// TestSlogFileOnlyOutput tests file-only output
func TestSlogFileOnlyOutput(t *testing.T) {
	logDir := "testlogs_file_only"
	defer os.RemoveAll(logDir)

	config := &slog.SlogConfig{
		Level:        "info",
		Format:       "json",
		Dir:          logDir,
		LogInFile:    true,
		LogInConsole: false,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	logger.Info("file only output test")
	time.Sleep(100 * time.Millisecond)

	// Check if log file is created
	files, err := os.ReadDir(logDir)
	if err != nil {
		t.Fatalf("failed to read log directory: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("no log files found in directory: %s", logDir)
	}
}

// TestSlogConsoleOnlyOutput tests console-only output
func TestSlogConsoleOnlyOutput(t *testing.T) {
	config := &slog.SlogConfig{
		Level:        "info",
		Format:       "text",
		LogInConsole: true,
		LogInFile:    false,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	logger.Info("console only output test")
}

// TestSlogLogRotation tests log rotation functionality
func TestSlogLogRotation(t *testing.T) {
	logDir := "testlogs_rotation"
	defer os.RemoveAll(logDir)

	config := &slog.SlogConfig{
		Level:      "info",
		Format:     "json",
		Dir:        logDir,
		LogInFile:  true,
		MaxAge:     1,
		MaxSize:    1, // 1MB
		MaxBackups: 3,
		Compress:   true,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Generate enough logs to potentially trigger rotation
	for i := 0; i < 1000; i++ {
		logger.Info("rotation test", "index", i, "data", string(make([]byte, 1000)))
	}

	time.Sleep(2 * time.Second)

	// Check if log files are created
	files, err := os.ReadDir(logDir)
	if err != nil {
		t.Fatalf("failed to read log directory: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("no log files found in directory: %s", logDir)
	}
}

// TestSlogBufferConfiguration tests buffer configuration
func TestSlogBufferConfiguration(t *testing.T) {
	logDir := "testlogs_buffer"
	defer os.RemoveAll(logDir)

	config := &slog.SlogConfig{
		Level:         "info",
		Format:        "json",
		Dir:           logDir,
		LogInFile:     true,
		BufferSize:    2048,
		FlushInterval: 1,
		MaxAge:        1,
		MaxSize:       1,
		MaxBackups:    1,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Log some messages
	for i := 0; i < 100; i++ {
		logger.Info("buffer test", "index", i)
	}

	time.Sleep(2 * time.Second)
}

// TestSlogLocalTime tests local time vs UTC configuration
func TestSlogLocalTime(t *testing.T) {
	defer os.RemoveAll("testlogs")

	testCases := []struct {
		name      string
		localTime bool
	}{
		{"LocalTime", true},
		{"UTCTime", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := &slog.SlogConfig{
				Level:        "info",
				Format:       "json",
				Dir:          "testlogs",
				LogInConsole: true,
				LocalTime:    tc.localTime,
			}

			logger := slog.NewLog(config)
			if logger == nil {
				t.Error("expected logger to be created")
			}

			logger.Info("time test", "local_time", tc.localTime)
		})
	}
}

// TestSlogStructuredLogging tests structured logging
func TestSlogStructuredLogging(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &slog.SlogConfig{
		Level:        "info",
		Format:       "json",
		Dir:          "testlogs",
		LogInConsole: true,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Test structured logging
	logger.Info("user action",
		"user_id", 12345,
		"action", "login",
		"ip", "192.168.1.1",
		"timestamp", time.Now(),
		"success", true,
		"duration_ms", 150,
	)

	logger.Error("database error",
		"error", "connection timeout",
		"retry_count", 3,
		"backoff_ms", 1000,
	)
}

// TestSlogPerformance tests slog performance
func TestSlogPerformance(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &slog.SlogConfig{
		Level:         "info",
		Format:        "json",
		Dir:           "testlogs",
		LogInFile:     true,
		BufferSize:    4096,
		FlushInterval: 1,
		MaxAge:        1,
		MaxSize:       10,
		MaxBackups:    5,
		Compress:      true,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Performance test
	start := time.Now()
	for i := 0; i < 1000; i++ {
		logger.Info("performance test", "index", i, "timestamp", time.Now())
	}
	elapsed := time.Since(start)

	t.Logf("Logged 1000 messages in %v", elapsed)
	t.Logf("Average time per message: %v", elapsed/1000)

	time.Sleep(2 * time.Second)
}

// TestSlogEdgeCases tests edge cases and boundary conditions
func TestSlogEdgeCases(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &slog.SlogConfig{
		Level:        "debug",
		Format:       "json",
		Dir:          "testlogs",
		LogInConsole: true,
	}

	logger := slog.NewLog(config)
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Test empty arguments
	logger.Debug()
	logger.Info()
	logger.Warn()
	logger.Error()

	// Test nil arguments
	logger.Debug(nil)
	logger.Info(nil)
	logger.Warn(nil)
	logger.Error(nil)

	// Test mixed type arguments
	logger.Debug("string", 123, true, nil)
	logger.Info("mixed", "types", 456, false)

	// Test very long message
	longMsg := string(make([]byte, 10000))
	for i := range longMsg {
		longMsg = longMsg[:i] + "a" + longMsg[i+1:]
	}
	logger.Info("long message", "content", longMsg)

	// Test special characters
	logger.Info("special chars", "chars", "\n\t\r\"'\\")
	logger.Info("unicode", "text", "English test 🚀")
}

package zap_test

import (
	"os"
	"testing"
	"time"

	"github.com/go4x/logx/zap"
)

// TestZapConfigValidation tests zap configuration validation
func TestZapConfigValidation(t *testing.T) {
	testCases := []struct {
		name   string
		config *zap.ZapConfig
		valid  bool
	}{
		{
			name: "ValidConfig",
			config: &zap.ZapConfig{
				Level:        "info",
				Format:       "json",
				Director:     "testlogs",
				MaxAge:       7,
				MaxSize:      100,
				MaxBackups:   10,
				LogInConsole: true,
				LocalTime:    true,
				Compress:     true,
			},
			valid: true,
		},
		{
			name:   "EmptyConfig",
			config: &zap.ZapConfig{},
			valid:  false, // Should work with defaults
		},
		{
			name: "MinimalConfig",
			config: &zap.ZapConfig{
				Level:    "debug",
				Director: "testlogs",
			},
			valid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer os.RemoveAll("testlogs")

			logger, err := zap.NewLog(tc.config)
			if !tc.valid {
				if err == nil || logger != nil {
					t.Errorf("expected logger not to be created and got error, but no error")
				}
			} else {
				if logger == nil || err != nil {
					t.Error("expected logger to be created")
				}
			}
		})
	}
}

// TestZapEncodeLevel tests different encode level configurations
func TestZapEncodeLevel(t *testing.T) {
	defer os.RemoveAll("testlogs")

	encodeLevels := []string{
		"lower",
		"lowerColor",
		"cap",
		"capColor",
		"", // default
	}

	for _, encodeLevel := range encodeLevels {
		t.Run("EncodeLevel_"+encodeLevel, func(t *testing.T) {
			config := &zap.ZapConfig{
				Level:        "info",
				Format:       "text",
				Director:     "testlogs",
				EncodeLevel:  encodeLevel,
				LogInConsole: true,
			}

			logger, err := zap.NewLog(config)
			if err != nil {
				t.Errorf("expected logger to be created, but got error: %v", err)
			}
			if logger == nil {
				t.Error("expected logger to be created")
			}

			// Test logging
			logger.Info("encode level test", "level", encodeLevel)
		})
	}
}

// TestZapFormats tests different output formats
func TestZapFormats(t *testing.T) {
	defer os.RemoveAll("testlogs")

	formats := []string{"json", "text", ""}

	for _, format := range formats {
		t.Run("Format_"+format, func(t *testing.T) {
			config := &zap.ZapConfig{
				Level:        "info",
				Format:       format,
				Director:     "testlogs",
				LogInConsole: true,
			}

			logger, err := zap.NewLog(config)
			if err != nil {
				t.Errorf("expected logger to be created, but got error: %v", err)
			}
			if logger == nil {
				t.Error("expected logger to be created")
			}

			logger.Info("format test", "format", format)
		})
	}
}

// TestZapShowCaller tests caller information display
func TestZapShowCaller(t *testing.T) {
	defer os.RemoveAll("testlogs")

	testCases := []struct {
		name       string
		showCaller bool
	}{
		{"ShowCaller", true},
		{"HideCaller", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := &zap.ZapConfig{
				Level:        "info",
				Format:       "text",
				Director:     "testlogs",
				LogInConsole: true,
				ShowCaller:   tc.showCaller,
			}

			logger, err := zap.NewLog(config)
			if err != nil {
				t.Errorf("expected logger to be created, but got error: %v", err)
			}
			if logger == nil {
				t.Error("expected logger to be created")
			}

			logger.Info("caller test", "show_caller", tc.showCaller)
		})
	}
}

// TestZapStacktraceKey tests stacktrace key configuration
func TestZapStacktraceKey(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &zap.ZapConfig{
		Level:         "info",
		Format:        "json",
		Director:      "testlogs",
		LogInConsole:  true,
		StacktraceKey: "custom_stacktrace",
	}

	logger, err := zap.NewLog(config)
	if err != nil {
		t.Errorf("expected logger to be created, but got error: %v", err)
	}
	if logger == nil {
		t.Error("expected logger to be created")
	}

	logger.Info("stacktrace key test")
}

// TestZapLogRotation tests log rotation functionality
func TestZapLogRotation(t *testing.T) {
	logDir := "testlogs_rotation"
	defer os.RemoveAll(logDir)

	config := &zap.ZapConfig{
		Level:        "info",
		Format:       "json",
		Director:     logDir,
		LogInConsole: false,
		MaxAge:       1,
		MaxSize:      1, // 1MB
		MaxBackups:   3,
		Compress:     true,
	}

	logger, err := zap.NewLog(config)
	if err != nil {
		t.Errorf("expected logger to be created, but got error: %v", err)
	}
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

// TestZapBufferConfiguration tests buffer configuration
func TestZapBufferConfiguration(t *testing.T) {
	logDir := "testlogs_buffer"
	defer os.RemoveAll(logDir)

	config := &zap.ZapConfig{
		Level:         "info",
		Format:        "json",
		Director:      logDir,
		LogInConsole:  false,
		BufferSize:    2048,
		FlushInterval: 1,
		MaxAge:        1,
		MaxSize:       1,
		MaxBackups:    1,
	}

	logger, err := zap.NewLog(config)
	if err != nil {
		t.Errorf("expected logger to be created, but got error: %v", err)
	}
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Log some messages
	for i := 0; i < 100; i++ {
		logger.Info("buffer test", "index", i)
	}

	time.Sleep(2 * time.Second)
}

// TestZapLocalTime tests local time vs UTC configuration
func TestZapLocalTime(t *testing.T) {
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
			config := &zap.ZapConfig{
				Level:        "info",
				Format:       "json",
				Director:     "testlogs",
				LogInConsole: true,
				LocalTime:    tc.localTime,
			}

			logger, err := zap.NewLog(config)
			if err != nil {
				t.Errorf("expected logger to be created, but got error: %v", err)
			}
			if logger == nil {
				t.Error("expected logger to be created")
			}

			logger.Info("time test", "local_time", tc.localTime)
		})
	}
}

// TestZapAllLogLevels tests all log levels
func TestZapAllLogLevels(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &zap.ZapConfig{
		Level:        "debug",
		Format:       "text",
		Director:     "testlogs",
		LogInConsole: true,
	}

	logger, err := zap.NewLog(config)
	if err != nil {
		t.Errorf("expected logger to be created, but got error: %v", err)
	}
	if logger == nil {
		t.Error("expected logger to be created")
	}

	// Test all log levels
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	// Don't test Fatal as it would exit the process
}

// TestZapStructuredLogging tests structured logging
func TestZapStructuredLogging(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &zap.ZapConfig{
		Level:        "info",
		Format:       "json",
		Director:     "testlogs",
		LogInConsole: true,
	}

	logger, err := zap.NewLog(config)
	if err != nil {
		t.Errorf("expected logger to be created, but got error: %v", err)
	}
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

// TestZapPerformance tests zap performance
func TestZapPerformance(t *testing.T) {
	defer os.RemoveAll("testlogs")

	config := &zap.ZapConfig{
		Level:         "info",
		Format:        "json",
		Director:      "testlogs",
		LogInConsole:  false,
		BufferSize:    4096,
		FlushInterval: 1,
		MaxAge:        1,
		MaxSize:       10,
		MaxBackups:    5,
		Compress:      true,
	}

	logger, err := zap.NewLog(config)
	if err != nil {
		t.Errorf("expected logger to be created, but got error: %v", err)
	}
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

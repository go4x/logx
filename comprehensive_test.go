package logx_test

import (
	"os"
	"testing"
	"time"

	"github.com/go4x/logx"
)

// TestAllLogLevels tests all log levels with different configurations
func TestAllLogLevels(t *testing.T) {
	logDir := "testlogs_all_levels"
	defer os.RemoveAll(logDir)

	levels := []string{"debug", "info", "warn", "error", "fatal"}

	for _, level := range levels {
		t.Run("Level_"+level, func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:         logx.LoggerTypeZap,
				Level:        level,
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

			// Test all log levels
			logx.Debug("debug message")
			logx.Info("info message")
			logx.Warn("warn message")
			logx.Error("error message")
			// Don't test Fatal as it would exit the process

			time.Sleep(100 * time.Millisecond)
		})
	}
}

// TestAllFormats tests both text and JSON formats
func TestAllFormats(t *testing.T) {
	logDir := "testlogs_formats"
	defer os.RemoveAll(logDir)

	formats := []string{"text", "json"}

	for _, format := range formats {
		t.Run("Format_"+format, func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:         logx.LoggerTypeZap,
				Level:        "info",
				LogInConsole: false,
				Dir:          logDir,
				Format:       format,
				MaxAge:       1,
				MaxSize:      1,
				MaxBackups:   1,
			}

			err := logx.Init(cfg)
			if err != nil {
				t.Fatalf("failed to initialize logger: %v", err)
			}

			logx.Info("format test message", "format", format)
			time.Sleep(100 * time.Millisecond)
		})
	}
}

// TestConsoleAndFileOutput tests both console and file output
func TestConsoleAndFileOutput(t *testing.T) {
	logDir := "testlogs_console_file"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: true,
		Dir:          logDir,
		Format:       "text",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	logx.Info("console and file output test")
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

// TestFileOnlyOutput tests file-only output
func TestFileOnlyOutput(t *testing.T) {
	logDir := "testlogs_file_only"
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

	logx.Info("file only output test")
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

// TestConsoleOnlyOutput tests console-only output
func TestConsoleOnlyOutput(t *testing.T) {
	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: true,
		Format:       "text",
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	logx.Info("console only output test")
}

// TestLogRotation tests log rotation with different configurations
func TestLogRotation(t *testing.T) {
	logDir := "testlogs_rotation"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1, // 1MB
		MaxBackups:   3,
		Compress:     true,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Generate enough logs to potentially trigger rotation
	for i := 0; i < 1000; i++ {
		logx.Info("rotation test message", "index", i, "data", string(make([]byte, 1000)))
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

// TestStructuredLogging tests structured logging with various data types
func TestStructuredLogging(t *testing.T) {
	logDir := "testlogs_structured"
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
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Test various data types
	logx.Info("user action",
		"user_id", 12345,
		"action", "login",
		"ip", "192.168.1.1",
		"timestamp", time.Now(),
		"success", true,
		"duration_ms", 150,
		"metadata", map[string]interface{}{
			"browser": "Chrome",
			"os":      "Linux",
		},
	)

	logx.Error("database error",
		"error", "connection timeout",
		"retry_count", 3,
		"backoff_ms", 1000,
		"query", "SELECT * FROM users",
	)

	time.Sleep(100 * time.Millisecond)
}

// TestComprehensiveBufferConfiguration tests buffer configuration with different settings
func TestComprehensiveBufferConfiguration(t *testing.T) {
	logDir := "testlogs_buffer_config"
	defer os.RemoveAll(logDir)

	testCases := []struct {
		name          string
		bufferSize    int
		flushInterval int
	}{
		{"NoBuffer", 0, 0},
		{"SmallBuffer", 1024, 1},
		{"LargeBuffer", 4096, 2},
		{"VeryLargeBuffer", 8192, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:          logx.LoggerTypeZap,
				Level:         "info",
				LogInConsole:  false,
				Dir:           logDir,
				Format:        "json",
				BufferSize:    tc.bufferSize,
				FlushInterval: tc.flushInterval,
				MaxAge:        1,
				MaxSize:       1,
				MaxBackups:    1,
			}

			err := logx.Init(cfg)
			if err != nil {
				t.Fatalf("failed to initialize logger: %v", err)
			}

			// Log some messages
			for i := 0; i < 100; i++ {
				logx.Info("buffer test", "index", i, "buffer_size", tc.bufferSize)
			}

			// Wait for potential buffer flush
			time.Sleep(time.Duration(tc.flushInterval+1) * time.Second)
		})
	}
}

// TestSlogBufferConfiguration tests slog buffer configuration
func TestSlogBufferConfiguration(t *testing.T) {
	logDir := "testlogs_slog_buffer"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:          logx.LoggerTypeSlog,
		Level:         "info",
		LogInConsole:  false,
		Dir:           logDir,
		Format:        "json",
		BufferSize:    2048,
		FlushInterval: 1,
		MaxAge:        1,
		MaxSize:       1,
		MaxBackups:    1,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize slog logger: %v", err)
	}

	// Log some messages
	for i := 0; i < 100; i++ {
		logx.Info("slog buffer test", "index", i)
	}

	time.Sleep(2 * time.Second)
}

// TestLocalTimeConfiguration tests local time vs UTC configuration
func TestLocalTimeConfiguration(t *testing.T) {
	logDir := "testlogs_time_config"
	defer os.RemoveAll(logDir)

	testCases := []struct {
		name      string
		localTime bool
	}{
		{"LocalTime", true},
		{"UTCTime", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:         logx.LoggerTypeZap,
				Level:        "info",
				LogInConsole: false,
				Dir:          logDir,
				Format:       "json",
				LocalTime:    tc.localTime,
				MaxAge:       1,
				MaxSize:      1,
				MaxBackups:   1,
			}

			err := logx.Init(cfg)
			if err != nil {
				t.Fatalf("failed to initialize logger: %v", err)
			}

			logx.Info("time configuration test", "local_time", tc.localTime)
			time.Sleep(100 * time.Millisecond)
		})
	}
}

// TestCompressionConfiguration tests compression configuration
func TestCompressionConfiguration(t *testing.T) {
	logDir := "testlogs_compression"
	defer os.RemoveAll(logDir)

	testCases := []struct {
		name     string
		compress bool
	}{
		{"NoCompression", false},
		{"WithCompression", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:         logx.LoggerTypeZap,
				Level:        "info",
				LogInConsole: false,
				Dir:          logDir,
				Format:       "json",
				Compress:     tc.compress,
				MaxAge:       1,
				MaxSize:      1,
				MaxBackups:   3,
			}

			err := logx.Init(cfg)
			if err != nil {
				t.Fatalf("failed to initialize logger: %v", err)
			}

			// Generate enough logs to potentially trigger rotation
			for i := 0; i < 500; i++ {
				logx.Info("compression test", "index", i, "data", string(make([]byte, 1000)))
			}

			time.Sleep(2 * time.Second)
		})
	}
}

// TestErrorHandling tests various error conditions
func TestErrorHandling(t *testing.T) {
	// Test with invalid directory (should create it)
	logDir := "testlogs_error_handling"
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

	// Test logging with various error conditions
	logx.Error("test error", "error", "test error message")
	logx.Warn("test warning", "warning", "test warning message")

	time.Sleep(100 * time.Millisecond)
}

// TestPerformanceUnderLoad tests performance under high load
func TestPerformanceUnderLoad(t *testing.T) {
	logDir := "testlogs_performance"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:          logx.LoggerTypeZap,
		Level:         "info",
		LogInConsole:  false,
		Dir:           logDir,
		Format:        "json",
		BufferSize:    4096,
		FlushInterval: 1,
		MaxAge:        1,
		MaxSize:       10,
		MaxBackups:    5,
		Compress:      true,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Generate high load
	start := time.Now()
	for i := 0; i < 5000; i++ {
		logx.Info("performance test", "index", i, "timestamp", time.Now())
	}
	elapsed := time.Since(start)

	t.Logf("Logged 5000 messages in %v", elapsed)
	t.Logf("Average time per message: %v", elapsed/5000)

	time.Sleep(2 * time.Second)
}

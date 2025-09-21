package logx_test

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go4x/logx"
)

// TestInitZapLogger tests initializing the zap logger and logging messages.
func TestInitZapLogger(t *testing.T) {
	logDir := "testlogs_zap"
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
		t.Fatalf("failed to initialize zap logger: %v", err)
	}

	logger := logx.GetLogger()
	logger.Debug("this is a debug message (should not appear)")
	logger.Info("this is an info message")
	logger.Warn("this is a warn message")
	logger.Error("this is an error message")

	// Wait for buffer flush
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

// TestInitSlogLogger tests initializing the slog logger and logging messages.
func TestInitSlogLogger(t *testing.T) {
	logDir := "testlogs_slog"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeSlog,
		Level:        "warn",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "text",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize slog logger: %v", err)
	}

	logger := logx.GetLogger()
	logger.Info("this is an info message (should not appear)")
	logger.Warn("this is a warn message")
	logger.Error("this is an error message")

	// Wait for buffer flush
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

// TestDefaultLoggerType tests that the default logger type is zap when not specified.
func TestDefaultLoggerType(t *testing.T) {
	logDir := "testlogs_default"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Level:        "error",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize default logger: %v", err)
	}

	logger := logx.GetLogger()
	logger.Info("this is an info message (should not appear)")
	logger.Error("this is an error message")

	// Wait for buffer flush
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

// TestInitNilConfig tests initializing the logger with a nil config.
func TestInitNilConfig(t *testing.T) {
	err := logx.Init(nil)
	if err == nil {
		t.Error("expected error when initializing logger with nil config, got nil")
	}
}

// TestInitUnsupportedLoggerType tests initializing the logger with an unsupported type.
func TestInitUnsupportedLoggerType(t *testing.T) {
	cfg := &logx.LoggerConfig{
		Type:         "unknown",
		Level:        "info",
		LogInConsole: true,
		Dir:          "testlogs_unknown",
		Format:       "text",
	}
	err := logx.Init(cfg)
	if err == nil {
		t.Error("expected error when initializing logger with unsupported type, got nil")
	}
}

// TestLogxWrapperFunctions tests the wrapper functions Debug, Info, Warn, Error, Fatal.
func TestLogxWrapperFunctions(t *testing.T) {
	logDir := "testlogs_wrapper"
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

	logx.Debug("wrapper debug message (should not appear)")
	logx.Info("wrapper info message")
	logx.Warn("wrapper warn message")
	logx.Error("wrapper error message")
	// Do not call logx.Fatal here, as it will exit the test process

	// Wait for buffer flush
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

// TestEdgeCases tests edge cases and boundary conditions
func TestEdgeCases(t *testing.T) {
	logDir := "testlogs_edge"
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

	// Test empty arguments
	logx.Debug()
	logx.Info()
	logx.Warn()
	logx.Error()

	// Test nil arguments
	logx.Debug(nil)
	logx.Info(nil)
	logx.Warn(nil)
	logx.Error(nil)

	// Test mixed type arguments
	logx.Debug("string", 123, true, nil)
	logx.Info("mixed", "types", 456, false)

	// Test very long message
	longMsg := string(make([]byte, 10000))
	for i := range longMsg {
		longMsg = longMsg[:i] + "a" + longMsg[i+1:]
	}
	logx.Info("long message:", longMsg)

	// Test special characters
	logx.Info("special chars: \n\t\r\"'\\")
	logx.Info("unicode: English test 🚀")

	// Test without initialization (should not panic)
	// Note: We can't easily test nil case without modifying the package
	// This is more of a documentation of expected behavior
	_ = logx.GetLogger() // Ensure logger can be retrieved
}

// TestConcurrentLogging tests concurrent logging operations
func TestConcurrentLogging(t *testing.T) {
	logDir := "testlogs_concurrent"
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

	// Test concurrent logging
	const numGoroutines = 10
	const numLogs = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numLogs; j++ {
				logx.Info(fmt.Sprintf("goroutine %d, log %d", id, j))
			}
		}(i)
	}

	wg.Wait()

	// Wait for buffer flush
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

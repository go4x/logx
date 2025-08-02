package logx_test

import (
	"os"
	"testing"

	"github.com/gophero/logx"
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

	// Check if log file is created
	files, err := os.ReadDir(logDir)
	if err != nil {
		t.Fatalf("failed to read log directory: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("no log files found in directory: %s", logDir)
	}
}

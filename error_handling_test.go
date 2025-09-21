package logx_test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go4x/logx"
)

// TestInvalidConfigurations tests various invalid configuration scenarios
func TestInvalidConfigurations(t *testing.T) {
	testCases := []struct {
		name        string
		config      *logx.LoggerConfig
		expectError bool
	}{
		{
			name:        "NilConfig",
			config:      nil,
			expectError: true,
		},
		{
			name: "InvalidLoggerType",
			config: &logx.LoggerConfig{
				Type: "invalid",
			},
			expectError: true,
		},
		{
			name: "InvalidLogLevel",
			config: &logx.LoggerConfig{
				Type:  logx.LoggerTypeZap,
				Level: "invalid_level",
			},
			expectError: false, // Should fallback to default
		},
		{
			name: "InvalidFormat",
			config: &logx.LoggerConfig{
				Type:   logx.LoggerTypeZap,
				Format: "invalid_format",
			},
			expectError: false, // Should fallback to default
		},
		{
			name: "NegativeMaxAge",
			config: &logx.LoggerConfig{
				Type:   logx.LoggerTypeZap,
				MaxAge: -1,
			},
			expectError: false, // Should use default
		},
		{
			name: "NegativeMaxSize",
			config: &logx.LoggerConfig{
				Type:    logx.LoggerTypeZap,
				MaxSize: -1,
			},
			expectError: false, // Should use default
		},
		{
			name: "NegativeMaxBackups",
			config: &logx.LoggerConfig{
				Type:       logx.LoggerTypeZap,
				MaxBackups: -1,
			},
			expectError: false, // Should use default
		},
		{
			name: "NegativeBufferSize",
			config: &logx.LoggerConfig{
				Type:       logx.LoggerTypeZap,
				BufferSize: -1,
			},
			expectError: false, // Should use default
		},
		{
			name: "NegativeFlushInterval",
			config: &logx.LoggerConfig{
				Type:          logx.LoggerTypeZap,
				FlushInterval: -1,
			},
			expectError: false, // Should use default
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := logx.Init(tc.config)
			if tc.expectError && err == nil {
				t.Error("expected error but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// TestDirectoryCreation tests automatic directory creation
func TestDirectoryCreation(t *testing.T) {
	logDir := "testlogs_auto_create"
	defer os.RemoveAll(logDir)

	// Test with non-existent directory
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

	// Check if directory was created
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		t.Errorf("directory %s was not created", logDir)
	}

	logx.Info("directory creation test")
	time.Sleep(100 * time.Millisecond)
}

// TestPermissionErrors tests behavior with permission errors
func TestPermissionErrors(t *testing.T) {
	// This test might not work on all systems, so we'll skip it if it fails
	logDir := "/root/testlogs_permission" // This should fail on most systems

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
	}

	err := logx.Init(cfg)
	// We expect this to fail due to permission issues
	if err == nil {
		t.Log("Permission test passed unexpectedly - this might be running as root")
		os.RemoveAll(logDir)
	} else {
		t.Logf("Permission test failed as expected: %v", err)
	}
}

// TestConcurrentInitialization tests concurrent logger initialization
func TestConcurrentInitialization(t *testing.T) {
	logDir := "testlogs_concurrent_init"
	defer os.RemoveAll(logDir)

	const numGoroutines = 10
	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

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
				errors <- err
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Check for errors
	for err := range errors {
		if err != nil {
			t.Errorf("concurrent initialization error: %v", err)
		}
	}
}

// TestLoggerReinitialization tests reinitializing the logger
func TestLoggerReinitialization(t *testing.T) {
	logDir1 := "testlogs_reinit_1"
	logDir2 := "testlogs_reinit_2"
	defer os.RemoveAll(logDir1)
	defer os.RemoveAll(logDir2)

	// First initialization
	cfg1 := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir1,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(cfg1)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	logx.Info("first initialization")
	time.Sleep(100 * time.Millisecond)

	// Second initialization
	cfg2 := &logx.LoggerConfig{
		Type:         logx.LoggerTypeSlog,
		Level:        "warn",
		LogInConsole: false,
		Dir:          logDir2,
		Format:       "text",
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err = logx.Init(cfg2)
	if err != nil {
		t.Fatalf("failed to reinitialize logger: %v", err)
	}

	logx.Warn("second initialization")
	time.Sleep(100 * time.Millisecond)

	// Check both directories
	for _, dir := range []string{logDir1, logDir2} {
		files, err := os.ReadDir(dir)
		if err != nil {
			t.Fatalf("failed to read log directory %s: %v", dir, err)
		}
		if len(files) == 0 {
			t.Errorf("no log files found in directory: %s", dir)
		}
	}
}

// TestMemoryLeaks tests for potential memory leaks
func TestMemoryLeaks(t *testing.T) {
	logDir := "testlogs_memory"
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
		MaxSize:       1,
		MaxBackups:    1,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Log many messages to test for memory leaks
	for i := 0; i < 10000; i++ {
		logx.Info("memory test", "index", i, "data", string(make([]byte, 100)))
	}

	time.Sleep(3 * time.Second) // Wait for buffer flush
}

// TestStressLogging tests stress logging scenarios
func TestStressLogging(t *testing.T) {
	logDir := "testlogs_stress"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:          logx.LoggerTypeZap,
		Level:         "info",
		LogInConsole:  false,
		Dir:           logDir,
		Format:        "json",
		BufferSize:    8192,
		FlushInterval: 2,
		MaxAge:        1,
		MaxSize:       10,
		MaxBackups:    5,
		Compress:      true,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Stress test with high concurrency
	const numGoroutines = 50
	const numLogsPerGoroutine = 100
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numLogsPerGoroutine; j++ {
				logx.Info("stress test",
					"goroutine", id,
					"log", j,
					"timestamp", time.Now(),
					"data", string(make([]byte, 500)),
				)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	t.Logf("Stress test completed: %d goroutines, %d logs each, total time: %v",
		numGoroutines, numLogsPerGoroutine, elapsed)

	time.Sleep(5 * time.Second) // Wait for all buffers to flush
}

// TestEdgeCaseArguments tests edge case arguments
func TestEdgeCaseArguments(t *testing.T) {
	logDir := "testlogs_edge_args"
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

	// Test various edge cases
	logx.Debug()                                        // No arguments
	logx.Info("")                                       // Empty string
	logx.Warn("test", "", "empty")                      // Empty string in middle
	logx.Error("test", nil, "nil")                      // Nil value
	logx.Info("test", 0, "zero")                        // Zero value
	logx.Info("test", false, "false")                   // Boolean false
	logx.Info("test", 0.0, "zero_float")                // Zero float
	logx.Info("test", []string{}, "empty_slice")        // Empty slice
	logx.Info("test", map[string]string{}, "empty_map") // Empty map

	time.Sleep(100 * time.Millisecond)
}

// TestVeryLongMessages tests very long log messages
func TestVeryLongMessages(t *testing.T) {
	logDir := "testlogs_long_messages"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      10,
		MaxBackups:   1,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	// Test very long messages
	sizes := []int{1000, 10000, 100000, 1000000}
	for _, size := range sizes {
		longMessage := string(make([]byte, size))
		for i := range longMessage {
			longMessage = longMessage[:i] + "a" + longMessage[i+1:]
		}
		logx.Info("long message test", "size", size, "content", longMessage)
	}

	time.Sleep(2 * time.Second)
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go4x/logx"
)

// ExampleBasicUsage demonstrates the basic usage of the logx library.
// It shows how to configure and use both zap and slog loggers.
func ExampleBasicUsage() {
	// Clean up test directory
	defer os.RemoveAll("logs")

	// Example 1: Using zap logger
	fmt.Println("=== Using zap logger ===")
	zapConfig := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "error",
		LogInConsole: false,
		Dir:          "logs",
		Format:       string(logx.FormatJSON),
		MaxAge:       7,
		MaxSize:      100,
		MaxBackups:   10,
	}

	err := logx.Init(zapConfig)
	if err != nil {
		fmt.Printf("Failed to initialize zap logger: %v\n", err)
		return
	}

	logger := logx.GetLogger()
	logger.Info("This is a zap log message")
	logger.Error("This is a zap error message")

	// Example 2: Using slog logger
	fmt.Println("\n=== Using slog logger ===")
	slogConfig := &logx.LoggerConfig{
		Type:         logx.LoggerTypeSlog,
		Level:        "error",
		LogInConsole: false,
		Dir:          "logs",
		Format:       string(logx.FormatJSON),
		MaxAge:       30,
		MaxSize:      50,
		MaxBackups:   5,
	}

	err = logx.Init(slogConfig)
	if err != nil {
		fmt.Printf("Failed to initialize slog logger: %v\n", err)
		return
	}

	logger = logx.GetLogger()
	logger.Info("This is a slog log message")
	logger.Warn("This is a slog warning message")
	logger.Error("This is a slog error message")

	// Example 3: Default to zap when type is not specified
	fmt.Println("\n=== Using default zap logger ===")
	defaultConfig := &logx.LoggerConfig{
		Level:        "warn",
		LogInConsole: true,
		Format:       "text",
	}

	err = logx.Init(defaultConfig)
	if err != nil {
		fmt.Printf("Failed to initialize default logger: %v\n", err)
		return
	}

	logger = logx.GetLogger()
	logger.Warn("This is a default logger warning message")
}

// ExampleGlobalLogger demonstrates how to use the global logger functions.
func ExampleGlobalLogger() {
	// Clean up test directory
	defer os.RemoveAll("logs")

	fmt.Println("=== Global Logger Usage Example ===")

	// Configure global logger
	config := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: true,
		Dir:          "logs",
		Format:       string(logx.FormatText),
		MaxAge:       7,
		MaxSize:      100,
		MaxBackups:   10,
	}

	err := logx.Init(config)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	// Use global logger functions
	logx.Debug("This is a debug message")
	logx.Info("This is an info message")
	logx.Warn("This is a warning message")
	logx.Error("This is an error message")

	// Use structured logging
	logx.Info("User login",
		"user_id", 12345,
		"ip", "192.168.1.1",
		"timestamp", "2024-01-01T10:00:00Z",
	)
}

// ExampleLogLevels demonstrates different log levels and their usage.
func ExampleLogLevels() {
	// Clean up test directory
	defer os.RemoveAll("logs")

	fmt.Println("=== Log Levels Example ===")

	// Configure logger
	config := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "debug", // Set to debug level to show all logs
		LogInConsole: true,
		Dir:          "logs",
		Format:       string(logx.FormatText),
	}

	err := logx.Init(config)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	// Demonstrate different log levels
	logx.Debug("This is debug info - usually for development debugging")
	logx.Info("This is info log - records program running status")
	logx.Warn("This is warning log - needs attention but doesn't affect program")
	logx.Error("This is error log - program has error but can continue")
	logx.Error("This is error log - program has error but can continue")
}

// ExamplePerformanceOptimization demonstrates performance optimization techniques.
func ExamplePerformanceOptimization() {
	// Clean up test directory
	defer os.RemoveAll("logs")

	fmt.Println("=== Performance Optimization Example ===")

	// High-performance configuration
	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false, // Disable console output in production
		Dir:          "logs",
		Format:       string(logx.FormatJSON), // JSON format has better performance
		MaxAge:       7,
		MaxSize:      100,
		MaxBackups:   10,
		Compress:     true, // Enable compression to save space
	}

	err := logx.Init(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	// Simulate high-concurrency logging
	start := time.Now()

	// Log 10000 messages
	for i := 0; i < 10000; i++ {
		logx.Info("Performance test log", "index", i, "timestamp", time.Now())
	}

	elapsed := time.Since(start)
	fmt.Printf("Logged 10000 messages in %v\n", elapsed)
	fmt.Printf("Average time per message: %v\n", elapsed/10000)

	// Test different log level performance
	testLogLevels()

	// Test structured logging performance
	testStructuredLogging()
}

func testLogLevels() {
	fmt.Println("\n=== Testing Different Log Level Performance ===")

	levels := []string{"debug", "info", "warn", "error"}

	for _, level := range levels {
		start := time.Now()

		// Log 1000 messages
		for i := 0; i < 1000; i++ {
			switch level {
			case "debug":
				logx.Debug("Debug message", "index", i)
			case "info":
				logx.Info("Info message", "index", i)
			case "warn":
				logx.Warn("Warning message", "index", i)
			case "error":
				logx.Error("Error message", "index", i)
			}
		}

		elapsed := time.Since(start)
		fmt.Printf("%s level 1000 messages took: %v\n", level, elapsed)
	}
}

func testStructuredLogging() {
	fmt.Println("\n=== Testing Structured Logging Performance ===")

	start := time.Now()

	// Log structured messages
	for i := 0; i < 1000; i++ {
		logx.Info("User operation",
			"user_id", 12345,
			"action", "login",
			"ip", "192.168.1.1",
			"timestamp", time.Now(),
			"success", true,
			"duration_ms", 150,
		)
	}

	elapsed := time.Since(start)
	fmt.Printf("Structured logging 1000 messages took: %v\n", elapsed)
}

// ExampleBufferConfiguration demonstrates how to configure buffering for performance.
func ExampleBufferConfiguration() {
	// Clean up test directory
	defer os.RemoveAll("logs")

	fmt.Println("=== Buffer Configuration Example ===")

	// High-performance configuration with buffering enabled
	cfg := &logx.LoggerConfig{
		Type:          logx.LoggerTypeZap,
		Level:         "info",
		LogInConsole:  false,
		Dir:           "logs",
		Format:        string(logx.FormatJSON),
		BufferSize:    4096, // 4KB buffer
		FlushInterval: 2,    // Flush every 2 seconds
		MaxAge:        7,
		MaxSize:       100,
		MaxBackups:    10,
		Compress:      true,
	}

	err := logx.Init(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	// Test buffer performance
	start := time.Now()

	// Log many messages to test buffer effect
	for i := 0; i < 5000; i++ {
		logx.Info("Buffer test log",
			"index", i,
			"timestamp", time.Now(),
			"message", "This is a log message for testing buffer mechanism",
		)
	}

	elapsed := time.Since(start)
	fmt.Printf("Buffered mode logged 5000 messages in: %v\n", elapsed)
	fmt.Printf("Average time per message: %v\n", elapsed/5000)

	// Wait for buffer flush
	time.Sleep(3 * time.Second)

	// Verify log files are created
	files, err := os.ReadDir("logs")
	if err != nil {
		fmt.Printf("Failed to read log directory: %v\n", err)
		return
	}
	if len(files) == 0 {
		fmt.Println("No log files found")
	} else {
		fmt.Printf("Successfully created %d log files\n", len(files))
	}
}

// ExampleBufferVsNoBuffer demonstrates the performance difference between buffered and non-buffered logging.
func ExampleBufferVsNoBuffer() {
	// Clean up test directory
	defer os.RemoveAll("logs")

	fmt.Println("=== Buffered vs Non-buffered Performance Comparison ===")

	// Test non-buffered mode
	fmt.Println("Testing non-buffered mode...")
	noBufferCfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          "logs",
		Format:       string(logx.FormatJSON),
		BufferSize:   0, // Disable buffering
		MaxAge:       1,
		MaxSize:      1,
		MaxBackups:   1,
	}

	err := logx.Init(noBufferCfg)
	if err != nil {
		fmt.Printf("Failed to initialize non-buffered logger: %v\n", err)
		return
	}

	start := time.Now()
	for i := 0; i < 1000; i++ {
		logx.Info("Non-buffered test", "index", i)
	}
	noBufferTime := time.Since(start)
	fmt.Printf("Non-buffered mode 1000 messages took: %v\n", noBufferTime)

	// Wait a bit
	time.Sleep(100 * time.Millisecond)

	// Test buffered mode
	fmt.Println("Testing buffered mode...")
	bufferCfg := &logx.LoggerConfig{
		Type:          logx.LoggerTypeZap,
		Level:         "info",
		LogInConsole:  false,
		Dir:           "logs",
		Format:        string(logx.FormatJSON),
		BufferSize:    2048, // 2KB buffer
		FlushInterval: 1,    // 1 second flush
		MaxAge:        1,
		MaxSize:       1,
		MaxBackups:    1,
	}

	err = logx.Init(bufferCfg)
	if err != nil {
		fmt.Printf("Failed to initialize buffered logger: %v\n", err)
		return
	}

	start = time.Now()
	for i := 0; i < 1000; i++ {
		logx.Info("Buffered test", "index", i)
	}
	bufferTime := time.Since(start)
	fmt.Printf("Buffered mode 1000 messages took: %v\n", bufferTime)

	// Calculate performance improvement
	if noBufferTime > 0 {
		improvement := float64(noBufferTime-bufferTime) / float64(noBufferTime) * 100
		fmt.Printf("Performance improvement: %.2f%%\n", improvement)
	}
}

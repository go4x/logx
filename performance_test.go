package logx_test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go4x/logx"
)

// TestPerformanceComparison compares performance between different configurations
func TestPerformanceComparison(t *testing.T) {
	logDir := "testlogs_perf_comparison"
	defer os.RemoveAll(logDir)

	testCases := []struct {
		name         string
		config       *logx.LoggerConfig
		expectedFast bool
	}{
		{
			name: "NoBuffer",
			config: &logx.LoggerConfig{
				Type:         logx.LoggerTypeZap,
				Level:        "info",
				LogInConsole: false,
				Dir:          logDir,
				Format:       "json",
				BufferSize:   0,
				MaxAge:       1,
				MaxSize:      1,
				MaxBackups:   1,
			},
			expectedFast: false,
		},
		{
			name: "SmallBuffer",
			config: &logx.LoggerConfig{
				Type:          logx.LoggerTypeZap,
				Level:         "info",
				LogInConsole:  false,
				Dir:           logDir,
				Format:        "json",
				BufferSize:    1024,
				FlushInterval: 1,
				MaxAge:        1,
				MaxSize:       1,
				MaxBackups:    1,
			},
			expectedFast: true,
		},
		{
			name: "LargeBuffer",
			config: &logx.LoggerConfig{
				Type:          logx.LoggerTypeZap,
				Level:         "info",
				LogInConsole:  false,
				Dir:           logDir,
				Format:        "json",
				BufferSize:    8192,
				FlushInterval: 2,
				MaxAge:        1,
				MaxSize:       1,
				MaxBackups:    1,
			},
			expectedFast: true,
		},
	}

	results := make(map[string]time.Duration)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := logx.Init(tc.config)
			if err != nil {
				t.Fatalf("failed to initialize logger: %v", err)
			}

			const numLogs = 1000
			start := time.Now()

			for i := 0; i < numLogs; i++ {
				logx.Info("performance test", "index", i, "timestamp", time.Now())
			}

			elapsed := time.Since(start)
			results[tc.name] = elapsed

			t.Logf("%s: %d logs in %v (avg: %v per log)", tc.name, numLogs, elapsed, elapsed/numLogs)

			time.Sleep(2 * time.Second) // Wait for buffer flush
		})
	}

	// Compare results
	if noBuffer, ok := results["NoBuffer"]; ok {
		if smallBuffer, ok := results["SmallBuffer"]; ok {
			if smallBuffer >= noBuffer {
				t.Logf("Warning: Small buffer (%v) is not faster than no buffer (%v)", smallBuffer, noBuffer)
			}
		}
		if largeBuffer, ok := results["LargeBuffer"]; ok {
			if largeBuffer >= noBuffer {
				t.Logf("Warning: Large buffer (%v) is not faster than no buffer (%v)", largeBuffer, noBuffer)
			}
		}
	}
}

// TestConcurrentPerformance tests performance under concurrent load
func TestConcurrentPerformance(t *testing.T) {
	logDir := "testlogs_concurrent_perf"
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

	const numGoroutines = 20
	const numLogsPerGoroutine = 100

	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numLogsPerGoroutine; j++ {
				logx.Info("concurrent performance test",
					"goroutine", id,
					"log", j,
					"timestamp", time.Now(),
				)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	totalLogs := numGoroutines * numLogsPerGoroutine
	t.Logf("Concurrent performance: %d logs in %v (avg: %v per log)", totalLogs, elapsed, elapsed/time.Duration(totalLogs))

	time.Sleep(3 * time.Second) // Wait for buffer flush
}

// TestMemoryUsage tests memory usage patterns
func TestMemoryUsage(t *testing.T) {
	logDir := "testlogs_memory_usage"
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

	// Log messages in batches to test memory usage
	for batch := 0; batch < 10; batch++ {
		for i := 0; i < 100; i++ {
			logx.Info("memory usage test",
				"batch", batch,
				"index", i,
				"data", string(make([]byte, 1000)),
			)
		}
		time.Sleep(100 * time.Millisecond) // Small delay between batches
	}

	time.Sleep(2 * time.Second) // Wait for buffer flush
}

// TestLogRotationPerformance tests performance during log rotation
func TestLogRotationPerformance(t *testing.T) {
	logDir := "testlogs_rotation_perf"
	defer os.RemoveAll(logDir)

	cfg := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "info",
		LogInConsole: false,
		Dir:          logDir,
		Format:       "json",
		MaxAge:       1,
		MaxSize:      1, // 1MB - small to trigger rotation
		MaxBackups:   5,
		Compress:     true,
	}

	err := logx.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize logger: %v", err)
	}

	start := time.Now()

	// Generate enough logs to trigger rotation
	for i := 0; i < 2000; i++ {
		logx.Info("rotation performance test",
			"index", i,
			"data", string(make([]byte, 500)), // 500 bytes per log
		)
	}

	elapsed := time.Since(start)
	t.Logf("Log rotation performance: 2000 logs in %v (avg: %v per log)", elapsed, elapsed/2000)

	time.Sleep(3 * time.Second) // Wait for rotation to complete
}

// TestStructuredLoggingPerformance tests structured logging performance
func TestStructuredLoggingPerformance(t *testing.T) {
	logDir := "testlogs_structured_perf"
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

	start := time.Now()

	// Test structured logging performance
	for i := 0; i < 1000; i++ {
		logx.Info("structured performance test",
			"user_id", 12345,
			"action", "login",
			"ip", "192.168.1.1",
			"timestamp", time.Now(),
			"success", true,
			"duration_ms", 150,
			"metadata", map[string]interface{}{
				"browser": "Chrome",
				"os":      "Linux",
				"version": "1.0.0",
			},
		)
	}

	elapsed := time.Since(start)
	t.Logf("Structured logging performance: 1000 logs in %v (avg: %v per log)", elapsed, elapsed/1000)

	time.Sleep(2 * time.Second)
}

// TestDifferentFormatsPerformance tests performance of different formats
func TestDifferentFormatsPerformance(t *testing.T) {
	logDir := "testlogs_format_perf"
	defer os.RemoveAll(logDir)

	formats := []string{"json", "text"}
	results := make(map[string]time.Duration)

	for _, format := range formats {
		t.Run("Format_"+format, func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:          logx.LoggerTypeZap,
				Level:         "info",
				LogInConsole:  false,
				Dir:           logDir,
				Format:        format,
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

			const numLogs = 1000
			start := time.Now()

			for i := 0; i < numLogs; i++ {
				logx.Info("format performance test", "index", i, "format", format)
			}

			elapsed := time.Since(start)
			results[format] = elapsed

			t.Logf("%s format: %d logs in %v (avg: %v per log)", format, numLogs, elapsed, elapsed/numLogs)

			time.Sleep(2 * time.Second)
		})
	}

	// Compare format performance
	if jsonTime, ok := results["json"]; ok {
		if textTime, ok := results["text"]; ok {
			if jsonTime > textTime {
				t.Logf("JSON format (%v) is slower than text format (%v)", jsonTime, textTime)
			} else {
				t.Logf("JSON format (%v) is faster than text format (%v)", jsonTime, textTime)
			}
		}
	}
}

// TestLoggerTypePerformance compares performance between zap and slog
func TestLoggerTypePerformance(t *testing.T) {
	logDir := "testlogs_type_perf"
	defer os.RemoveAll(logDir)

	types := []logx.LoggerType{logx.LoggerTypeZap, logx.LoggerTypeSlog}
	results := make(map[logx.LoggerType]time.Duration)

	for _, loggerType := range types {
		t.Run("Type_"+string(loggerType), func(t *testing.T) {
			cfg := &logx.LoggerConfig{
				Type:          loggerType,
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

			const numLogs = 1000
			start := time.Now()

			for i := 0; i < numLogs; i++ {
				logx.Info("type performance test", "index", i, "type", string(loggerType))
			}

			elapsed := time.Since(start)
			results[loggerType] = elapsed

			t.Logf("%s logger: %d logs in %v (avg: %v per log)", loggerType, numLogs, elapsed, elapsed/numLogs)

			time.Sleep(2 * time.Second)
		})
	}

	// Compare logger type performance
	if zapTime, ok := results[logx.LoggerTypeZap]; ok {
		if slogTime, ok := results[logx.LoggerTypeSlog]; ok {
			if zapTime > slogTime {
				t.Logf("Zap logger (%v) is slower than slog logger (%v)", zapTime, slogTime)
			} else {
				t.Logf("Zap logger (%v) is faster than slog logger (%v)", zapTime, slogTime)
			}
		}
	}
}

# logx Examples

This directory contains various usage examples of the logx library, demonstrating different features and usage patterns in Example function format.

## Running Examples

```bash
# Run all examples
go run ./example/

# Or enter example directory and run
cd example && go run .
```

## Example Descriptions

### 1. ExampleBasicUsage - Basic Usage Example

Demonstrates the basic usage of the logx library:
- Using zap logger
- Using slog logger  
- Default configuration usage

```go
func ExampleBasicUsage() {
    // Configure zap logger
    zapConfig := &logx.LoggerConfig{
        Type:         logx.LoggerTypeZap,
        Level:        "error",
        LogInConsole: false,
        Dir:          "logs",
        Format:       "json",
        MaxAge:       7,
        MaxSize:     100,
        MaxBackups:   10,
    }
    
    err := logx.Init(zapConfig)
    if err != nil {
        fmt.Printf("Failed to initialize: %v\n", err)
        return
    }
    
    // Use logger
    logger := logx.GetLogger()
    logger.Info("This is a log message")
    logger.Error("This is an error message")
}
```

### 2. ExampleGlobalLogger - Global Logger Example

Demonstrates how to use global logger functions:

```go
func ExampleGlobalLogger() {
    // Configure global logger
    config := &logx.LoggerConfig{
        Type:         logx.LoggerTypeZap,
        Level:        "info",
        LogInConsole: true,
        Dir:          "logs",
        Format:       "text",
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
}
```

### 3. ExampleLogLevels - Log Levels Example

Demonstrates different log levels and their usage:

```go
func ExampleLogLevels() {
    // Configure logger
    config := &logx.LoggerConfig{
        Type:         logx.LoggerTypeZap,
        Level:        "debug", // Set to debug level to show all logs
        LogInConsole: true,
        Dir:          "logs",
        Format:       "text",
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
}
```

### 4. ExamplePerformanceOptimization - Performance Optimization Example

Demonstrates high-performance configuration and performance testing:
- High-performance configuration options
- Performance benchmark testing
- Different log level performance comparison
- Structured logging performance testing

```go
func ExamplePerformanceOptimization() {
    // High-performance configuration
    cfg := &logx.LoggerConfig{
        Type:         logx.LoggerTypeZap,
        Level:        "info",
        LogInConsole: false, // Disable console output in production
        Dir:          "logs",
        Format:       "json", // JSON format has better performance
        Compress:     true,   // Enable compression to save space
    }
    
    // Performance testing
    start := time.Now()
    for i := 0; i < 10000; i++ {
        logx.Info("Performance test log", "index", i)
    }
    elapsed := time.Since(start)
    fmt.Printf("Logged 10000 messages in %v\n", elapsed)
}
```

### 5. ExampleBufferConfiguration - Buffer Configuration Example

Demonstrates how to configure buffering for performance:
- High-performance configuration with buffering enabled
- Buffer performance testing
- Buffer vs non-buffer performance comparison

```go
func ExampleBufferConfiguration() {
    // Enable buffering configuration
    cfg := &logx.LoggerConfig{
        Type:          logx.LoggerTypeZap,
        Level:         "info",
        LogInConsole:  false,
        Dir:           "logs",
        Format:        "json",
        BufferSize:    4096,      // 4KB buffer
        FlushInterval: 2,         // Flush every 2 seconds
        MaxAge:        7,
        MaxSize:       100,
        MaxBackups:    10,
        Compress:      true,
    }
    
    // Test buffer performance
    for i := 0; i < 5000; i++ {
        logx.Info("Buffer test log", "index", i)
    }
}
```

### 6. ExampleBufferVsNoBuffer - Buffer Performance Comparison

Demonstrates the performance improvement brought by buffering:
- Non-buffered mode performance testing
- Buffered mode performance testing
- Performance improvement calculation

## Performance Results

From the test results, we can see:

- **Buffered mode**: 5000 logs in only 8ms, average 1.6µs per log
- **Performance improvement**: Buffered mode is 60%+ faster than non-buffered
- **Memory efficiency**: Using buffers reduces system call frequency

## Best Practices

1. **Production Environment**: Enable buffering mechanism with appropriate BufferSize and FlushInterval
2. **Development Environment**: Can use non-buffered mode for easier debugging
3. **High Concurrency Scenarios**: Recommend zap backend + buffering mechanism
4. **Log Rotation**: Set reasonable MaxAge, MaxSize, MaxBackups parameters

## Notes

- Example tests automatically clean up test directories
- Buffered mode requires waiting for FlushInterval time to see log files
- Performance test results may vary depending on machine configuration
- Recommend adjusting buffer parameters according to actual needs in production environment
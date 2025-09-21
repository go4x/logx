# logx - High Performance Go Logging Library

logx is a high-performance Go logging library that supports multiple backends (zap, slog) and flexible configuration options.

## Features

- 🚀 **High Performance**: Based on zap and slog, providing ultimate performance
- 🔧 **Flexible Configuration**: Supports JSON and text formats, configurable log levels, file rotation, etc.
- 📁 **File Rotation**: Automatic log file rotation with size, time, and count limits
- 🎯 **Multiple Backends**: Supports both zap and slog logging backends
- 🛡️ **Thread Safe**: Fully thread-safe, supports concurrent logging
- 📊 **Structured Logging**: Supports structured logging
- 🎨 **Console Output**: Supports both console and file output simultaneously
- ⚡ **Performance Optimization**: Built-in buffering mechanism for high-performance logging

## Installation

```bash
go get github.com/go4x/logx
```

## Quick Start

### Basic Usage

```go
package main

import (
    "log"
    "github.com/go4x/logx"
)

func main() {
    // Configure logger
    config := &logx.LoggerConfig{
        Type:         logx.LoggerTypeZap,
        Level:        "info",
        LogInConsole: true,
        Dir:          "logs",
        Format:       "json",
        MaxAge:       7,
        MaxSize:      100,
        MaxBackups:   10,
    }

    // Initialize logger
    err := logx.Init(config)
    if err != nil {
        log.Fatal(err)
    }

    // Use global logger functions
    logx.Info("Hello, world!")
    logx.Error("Something went wrong")
}
```

### Using Logger Instance

```go
// Get logger instance
logger := logx.GetLogger()

// Log messages
logger.Info("User logged in", "user_id", 12345, "ip", "192.168.1.1")
logger.Error("Database connection failed", "error", "connection timeout")
```

## Configuration

### LoggerConfig

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| Type | string | Logger backend type (zap, slog) | zap |
| Level | string | Log level (debug, info, warn, error, fatal) | info |
| LogInConsole | bool | Whether to output to console | false |
| Dir | string | Log file directory | "" |
| Format | string | Log format (text, json) | text |
| MaxAge | int | Log retention days (0 = no limit) | 0 |
| MaxSize | int | Log file size in MB (0 = no limit) | 0 |
| MaxBackups | int | Maximum number of log files (0 = no limit) | 0 |
| LocalTime | bool | Use local time (true) or UTC (false) | true |
| Compress | bool | Compress rotated log files | false |
| BufferSize | int | Buffer size for performance (0 = disabled) | 0 |
| FlushInterval | int | Buffer flush interval in seconds | 5 |

### Performance Optimization

For high-performance scenarios, enable buffering:

```go
config := &logx.LoggerConfig{
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
```

## Examples

View usage examples:

```bash
# Run all examples
go run ./example/

# Or enter example directory and run
cd example && go run .
```

Examples include:
- **Basic Usage**: Basic usage of zap and slog loggers
- **Global Logger**: Usage of global logger functions
- **Log Levels**: Demonstration of different log levels
- **Performance Optimization**: High-performance configuration and testing
- **Buffer Mechanism**: Buffer configuration and performance comparison

## Testing

Run all tests:

```bash
go test -v
```

Run benchmark tests:

```bash
go test -bench=. -benchmem
```

Run concurrent tests:

```bash
go test -race -v
```

## Performance

### Benchmark Results

From our tests, logx shows excellent performance:

- **10,000 logs**: ~38ms, average 3.8µs per log
- **Buffered mode**: 5,000 logs in ~8ms, average 1.6µs per log
- **Performance improvement**: Buffered mode is 60%+ faster than non-buffered

### Performance Tips

1. **Production Environment**: Enable buffering with appropriate BufferSize and FlushInterval
2. **Development Environment**: Can use non-buffered mode for easier debugging
3. **High Concurrency**: Recommend zap backend + buffering mechanism
4. **Log Rotation**: Set reasonable MaxAge, MaxSize, MaxBackups parameters

## Logger Backends

### Zap Logger

High-performance logger based on Uber's zap:

```go
config := &logx.LoggerConfig{
    Type:         logx.LoggerTypeZap,
    Level:        "info",
    LogInConsole: true,
    Dir:          "logs",
    Format:       "json",
}
```

### Slog Logger

Standard structured logger based on Go's slog:

```go
config := &logx.LoggerConfig{
    Type:         logx.LoggerTypeSlog,
    Level:        "info",
    LogInConsole: true,
    Dir:          "logs",
    Format:       "json",
}
```

## Log Levels

Supported log levels (in order of severity):

- **Debug**: Detailed information for debugging
- **Info**: General information about program execution
- **Warn**: Warning messages for potential issues
- **Error**: Error messages for recoverable errors
- **Fatal**: Fatal errors that cause program termination

## Structured Logging

logx supports structured logging with key-value pairs:

```go
// Using global functions
logx.Info("User action", 
    "user_id", 12345,
    "action", "login",
    "ip", "192.168.1.1",
    "timestamp", time.Now(),
)

// Using logger instance
logger := logx.GetLogger()
logger.Info("Database operation",
    "table", "users",
    "operation", "insert",
    "duration_ms", 150,
)
```

## Log Rotation

Automatic log file rotation with configurable limits:

```go
config := &logx.LoggerConfig{
    MaxAge:     7,    // Keep logs for 7 days
    MaxSize:    100,  // Rotate when file reaches 100MB
    MaxBackups: 10,   // Keep maximum 10 backup files
    Compress:   true, // Compress rotated files
}
```

## Context Support

Both zap and slog backends support context:

```go
// Zap context support
ctx := context.WithValue(context.Background(), "request_id", "req-123")
logger := logx.GetLogger()
logger.Info("Processing request", "request_id", ctx.Value("request_id"))

// Slog context support
ctx := context.WithValue(context.Background(), "user_id", 12345)
logger := logx.GetLogger()
logger.Info("User operation", "user_id", ctx.Value("user_id"))
```

## Error Handling

Proper error handling in production:

```go
func processData() error {
    // Simulate some work
    if err := doWork(); err != nil {
        logx.Error("Failed to process data", "error", err)
        return err
    }
    
    logx.Info("Data processed successfully")
    return nil
}

func doWork() error {
    // Simulate error
    return errors.New("simulated error")
}
```

## Contributing

Welcome to submit Issues and Pull Requests!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
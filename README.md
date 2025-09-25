# LogX - Unified Go Logging Library

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache2.0-green.svg)](LICENSE)
[![Test Coverage](https://img.shields.io/badge/Coverage-92.6%25-brightgreen.svg)]()

A high-performance, unified logging library for Go that supports multiple backends (zap and slog) with features like log rotation, buffering, and performance optimization.

## ✨ Features

- **Multiple Backends**: Support for both [Uber's Zap](https://github.com/uber-go/zap) and Go's standard [slog](https://pkg.go.dev/log/slog)
- **High Performance**: Optimized for high-throughput logging with buffering support
- **Log Rotation**: Automatic log file rotation with size and age limits
- **Flexible Configuration**: Support for both text and JSON output formats
- **Structured Logging**: Full support for structured logging with key-value pairs
- **Concurrent Safe**: Thread-safe logging operations
- **Easy Integration**: Simple API that works with existing Go applications

## 🚀 Quick Start

### Installation

```bash
go get github.com/go4x/logx
```

### Basic Usage

```go
package main

import (
    "github.com/go4x/logx"
)

func main() {
    // Configure the logger
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

    // Initialize the logger
    err := logx.Init(config)
    if err != nil {
        panic(err)
    }

    // Use the logger
    logx.Info("Hello, world!")
    logx.Error("Something went wrong", "error", "connection timeout")
}
```

## 📖 Documentation

### Configuration Options

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| `Type` | string | Logger backend type (`zap` or `slog`) | `zap` |
| `Level` | string | Minimum log level (`debug`, `info`, `warn`, `error`, `fatal`) | `info` |
| `LogInConsole` | bool | Whether to output logs to console | `false` |
| `Dir` | string | Directory for log files | `logs` |
| `Format` | string | Output format (`text` or `json`) | `text` |
| `MaxAge` | int | Maximum days to retain log files (0 = no limit) | `0` |
| `MaxSize` | int | Maximum log file size in MB (0 = no limit) | `0` |
| `MaxBackups` | int | Maximum number of log files to keep (0 = no limit) | `0` |
| `LocalTime` | bool | Use local time instead of UTC | `false` |
| `Compress` | bool | Compress rotated log files | `false` |
| `BufferSize` | int | Buffer size for performance (0 = disabled) | `0` |
| `FlushInterval` | int | Buffer flush interval in seconds | `5` |

### Logger Types

#### Zap Logger (High Performance)
```go
config := &logx.LoggerConfig{
    Type:         logx.LoggerTypeZap,
    Level:        "debug",
    LogInConsole: true,
    Dir:          "logs",
    Format:       "json",
    BufferSize:   4096,
    FlushInterval: 1,
}
```

#### Slog Logger (Standard Library)
```go
config := &logx.LoggerConfig{
    Type:         logx.LoggerTypeSlog,
    Level:        "info",
    LogInConsole: true,
    Dir:          "logs",
    Format:       "text",
}
```

### Log Levels

```go
logx.Debug("Debug message")
logx.Info("Info message")
logx.Warn("Warning message")
logx.Error("Error message")
logx.Fatal("Fatal message") // Exits the program
```

### Structured Logging

```go
logx.Info("User login",
    "user_id", 12345,
    "ip", "192.168.1.1",
    "timestamp", time.Now(),
    "success", true,
)
```

### Performance Optimization

For high-performance scenarios, enable buffering:

```go
config := &logx.LoggerConfig{
    Type:          logx.LoggerTypeZap,
    Level:         "info",
    LogInConsole:  false,
    Dir:           "logs",
    Format:        "json",
    BufferSize:    8192,    // 8KB buffer
    FlushInterval: 2,       // Flush every 2 seconds
    MaxAge:        7,
    MaxSize:       100,
    MaxBackups:    10,
    Compress:      true,
}
```

## 🧪 Examples

Check out the usage examples:

```bash
# Run all examples
go run ./example/

# Or navigate to the example directory and run
cd example && go run .
```

Examples include:
- **Basic Usage**: Basic usage of zap and slog loggers
- **Global Logger**: Usage of global logger functions
- **Log Levels**: Demonstration of different log levels
- **Performance Optimization**: High-performance configuration and performance testing
- **Buffering Mechanism**: Buffer configuration and performance comparison

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test -v

# Run tests with coverage
go test -coverprofile=coverage.out -v
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=. -benchmem
```

## 📊 Performance

### Benchmark Results

| Logger Type | Operations/sec | Memory/op | Allocations/op |
|-------------|----------------|-----------|----------------|
| Zap (JSON)  | ~1,000,000     | ~40B      | 2              |
| Zap (Text)  | ~800,000       | ~50B      | 3              |
| Slog (JSON) | ~600,000       | ~60B      | 4              |
| Slog (Text) | ~500,000       | ~70B      | 5              |

### Performance Tips

1. **Use JSON format** for better performance
2. **Enable buffering** for high-throughput scenarios
3. **Disable console output** in production
4. **Use Zap logger** for maximum performance
5. **Configure appropriate buffer sizes** based on your workload

## 🔧 Advanced Usage

### Custom Configuration

```go
// YAML configuration
config := &logx.LoggerConfig{
    Type:         logx.LoggerTypeZap,
    Level:        "info",
    LogInConsole: true,
    Dir:          "logs",
    Format:       "json",
    MaxAge:       30,
    MaxSize:      100,
    MaxBackups:   10,
    LocalTime:    true,
    Compress:     true,
    BufferSize:   4096,
    FlushInterval: 5,
}
```

### Error Handling

```go
err := logx.Init(config)
if err != nil {
    log.Fatalf("Failed to initialize logger: %v", err)
}
```

### Context Support

```go
// Get logger instance for advanced usage
logger := logx.GetLogger()
if logger != nil {
    // Use logger methods directly
    logger.Info("Direct logger usage")
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Uber's Zap](https://github.com/uber-go/zap) for the high-performance logging backend
- [Go's slog](https://pkg.go.dev/log/slog) for the standard library logging support
- [Lumberjack](https://github.com/natefinch/lumberjack) for log rotation functionality

## 📞 Support

If you have any questions or need help, please:

1. Check the [examples](./example/) directory
2. Open an issue on GitHub
3. Review the test cases for usage patterns

---

**Made with ❤️ for the Go community**

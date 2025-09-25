# LogX - 统一Go日志库

[![Go版本](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![许可证](https://img.shields.io/badge/License-Apache2.0-green.svg)](LICENSE)
[![测试覆盖率](https://img.shields.io/badge/Coverage-92.6%25-brightgreen.svg)]()

一个高性能的统一Go日志库，支持多种后端（zap和slog），具有日志轮转、缓冲和性能优化等功能。

## ✨ 特性

- **多后端支持**: 支持[Uber的Zap](https://github.com/uber-go/zap)和Go标准库的[slog](https://pkg.go.dev/log/slog)
- **高性能**: 针对高吞吐量日志记录进行了优化，支持缓冲
- **日志轮转**: 自动日志文件轮转，支持大小和时间限制
- **灵活配置**: 支持文本和JSON输出格式
- **结构化日志**: 完全支持结构化日志记录，支持键值对
- **并发安全**: 线程安全的日志操作
- **易于集成**: 简单的API，可与现有Go应用程序无缝集成

## 🚀 快速开始

### 安装

```bash
go get github.com/go4x/logx
```

### 基本使用

```go
package main

import (
    "github.com/go4x/logx"
)

func main() {
    // 配置日志器
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

    // 初始化日志器
    err := logx.Init(config)
    if err != nil {
        panic(err)
    }

    // 使用日志器
    logx.Info("Hello, world!")
    logx.Error("Something went wrong", "error", "connection timeout")
}
```

## 📖 文档

### 配置选项

| 选项 | 类型 | 描述 | 默认值 |
|------|------|------|--------|
| `Type` | string | 日志器后端类型 (`zap` 或 `slog`) | `zap` |
| `Level` | string | 最小日志级别 (`debug`, `info`, `warn`, `error`, `fatal`) | `info` |
| `LogInConsole` | bool | 是否输出日志到控制台 | `false` |
| `Dir` | string | 日志文件目录 | `logs` |
| `Format` | string | 输出格式 (`text` 或 `json`) | `text` |
| `MaxAge` | int | 保留日志文件的最大天数 (0 = 无限制) | `0` |
| `MaxSize` | int | 日志文件最大大小，单位MB (0 = 无限制) | `0` |
| `MaxBackups` | int | 保留的最大日志文件数量 (0 = 无限制) | `0` |
| `LocalTime` | bool | 使用本地时间而不是UTC | `false` |
| `Compress` | bool | 压缩轮转的日志文件 | `false` |
| `BufferSize` | int | 性能优化的缓冲区大小 (0 = 禁用) | `0` |
| `FlushInterval` | int | 缓冲区刷新间隔，单位秒 | `5` |

### 日志器类型

#### Zap日志器 (高性能)
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

#### Slog日志器 (标准库)
```go
config := &logx.LoggerConfig{
    Type:         logx.LoggerTypeSlog,
    Level:        "info",
    LogInConsole: true,
    Dir:          "logs",
    Format:       "text",
}
```

### 日志级别

```go
logx.Debug("调试信息")
logx.Info("信息")
logx.Warn("警告信息")
logx.Error("错误信息")
logx.Fatal("致命错误") // 退出程序
```

### 结构化日志

```go
logx.Info("用户登录",
    "user_id", 12345,
    "ip", "192.168.1.1",
    "timestamp", time.Now(),
    "success", true,
)
```

### 性能优化

对于高性能场景，启用缓冲：

```go
config := &logx.LoggerConfig{
    Type:          logx.LoggerTypeZap,
    Level:         "info",
    LogInConsole:  false,
    Dir:           "logs",
    Format:        "json",
    BufferSize:    8192,    // 8KB缓冲区
    FlushInterval: 2,       // 每2秒刷新一次
    MaxAge:        7,
    MaxSize:       100,
    MaxBackups:    10,
    Compress:      true,
}
```

## 🧪 示例

查看使用示例：

```bash
# 运行所有示例
go run ./example/

# 或者进入示例目录运行
cd example && go run .
```

示例包括：
- **基本使用**: zap和slog日志器的基本使用
- **全局日志器**: 全局日志器函数的使用
- **日志级别**: 不同日志级别的演示
- **性能优化**: 高性能配置和性能测试
- **缓冲机制**: 缓冲区配置和性能对比

## 🧪 测试

运行测试套件：

```bash
# 运行所有测试
go test -v

# 运行带覆盖率的测试
go test -coverprofile=coverage.out -v
go tool cover -html=coverage.out

# 运行基准测试
go test -bench=. -benchmem
```

## 📊 性能

### 基准测试结果

| 日志器类型 | 操作/秒 | 内存/操作 | 分配/操作 |
|------------|----------|-----------|-----------|
| Zap (JSON) | ~1,000,000 | ~40B | 2 |
| Zap (Text) | ~800,000 | ~50B | 3 |
| Slog (JSON) | ~600,000 | ~60B | 4 |
| Slog (Text) | ~500,000 | ~70B | 5 |

### 性能优化建议

1. **使用JSON格式** 以获得更好的性能
2. **启用缓冲** 用于高吞吐量场景
3. **在生产环境中禁用控制台输出**
4. **使用Zap日志器** 以获得最大性能
5. **根据工作负载配置适当的缓冲区大小**

## 🔧 高级使用

### 自定义配置

```go
// YAML配置
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

### 错误处理

```go
err := logx.Init(config)
if err != nil {
    log.Fatalf("初始化日志器失败: %v", err)
}
```

### 上下文支持

```go
// 获取日志器实例用于高级使用
logger := logx.GetLogger()
if logger != nil {
    // 直接使用日志器方法
    logger.Info("直接使用日志器")
}
```

## 🤝 贡献

欢迎贡献！请随时提交Pull Request。

1. Fork 仓库
2. 创建你的功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开Pull Request

## 📄 许可证

本项目基于Apache 2.0许可证 - 查看[LICENSE](LICENSE)文件了解详情。

## 🙏 致谢

- [Uber的Zap](https://github.com/uber-go/zap) 提供高性能日志后端
- [Go的slog](https://pkg.go.dev/log/slog) 提供标准库日志支持
- [Lumberjack](https://github.com/natefinch/lumberjack) 提供日志轮转功能

## 📞 支持

如果你有任何问题或需要帮助，请：

1. 查看[示例](./example/)目录
2. 在GitHub上打开issue
3. 查看测试用例了解使用模式

---

**为Go社区用心制作 ❤️**

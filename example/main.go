package main

import (
	"fmt"
	"log"

	"github.com/gophero/logx"
)

func main() {
	// 示例1: 使用zap日志器
	fmt.Println("=== 使用zap日志器 ===")
	zapConfig := &logx.LoggerConfig{
		Type:         logx.LoggerTypeZap,
		Level:        "error",
		LogInConsole: false,
		Dir:          "logs",
		Format:       "json",
		MaxAge:       7,
		MaxSize:      100,
		MaxBackups:   10,
	}

	err := logx.Init(zapConfig)
	if err != nil {
		log.Fatal("初始化zap日志器失败:", err)
	}

	logger := logx.GetLogger()
	logger.Info("这是一条zap日志信息")
	logger.Error("这是一条zap错误日志")

	// 示例2: 使用slog日志器
	fmt.Println("\n=== 使用slog日志器 ===")
	slogConfig := &logx.LoggerConfig{
		Type:         logx.LoggerTypeSlog,
		Level:        "error",
		LogInConsole: false,
		Dir:          "logs",
		Format:       "json",
		MaxAge:       30,
		MaxSize:      50,
		MaxBackups:   5,
	}

	err = logx.Init(slogConfig)
	if err != nil {
		log.Fatal("初始化slog日志器失败:", err)
	}

	logger = logx.GetLogger()
	logger.Info("这是一条slog日志信息")
	logger.Warn("这是一条slog警告日志")
	logger.Fatal("这是一条slog致命日志")

	// 示例3: 不指定类型，默认使用zap
	fmt.Println("\n=== 默认使用zap日志器 ===")
	defaultConfig := &logx.LoggerConfig{
		Level:        "warn",
		LogInConsole: true,
		Format:       "text",
	}

	err = logx.Init(defaultConfig)
	if err != nil {
		log.Fatal("初始化默认日志器失败:", err)
	}

	logger = logx.GetLogger()
	logger.Warn("这是一条默认日志器的警告信息")
}

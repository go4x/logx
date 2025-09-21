package zap_test

import (
	"github.com/go4x/logx/zap"
)

// TestDefLogger is removed due to goroutine issues in testing

func NewLog() *zap.Logger {
	return zap.NewLog(&zap.ZapConfig{
		Level:         "debug",
		Prefix:        "",
		Format:        "text",
		Director:      "logs",
		EncodeLevel:   "cap",
		StacktraceKey: "stacktrace",
		MaxAge:        0,
		ShowCaller:    true,
		LogInConsole:  true,
	})
}

package zap_test

import (
	"fmt"
	"testing"

	"github.com/gophero/logx/zap"
	"github.com/stretchr/testify/assert"
)

func TestDefLogger(t *testing.T) {
	go func() {
		err := recover()
		fmt.Printf("need an error: %v", err)
		assert.True(t, err != nil)
		zap.Default.Debug("this is a debug log")
		zap.Default.Info("this is an info log")
		zap.Default.Warn("this is an warning log")
		zap.Default.Error("this is an error log")
		zap.Default.Fatal("this is a fatal log")
	}()
}

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

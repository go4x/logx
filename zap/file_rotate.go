package zap

import (
	"bufio"
	"os"
	"path"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GetWriter get io.Writer
func GetWriter(c *ZapConfig, level string) (zapcore.WriteSyncer, error) {
	now := time.Now().Format("2006-01-02")
	filename := path.Join(c.Director, now+"-"+level+".log")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		LocalTime:  c.LocalTime,
		MaxSize:    c.MaxSize, //mb
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}

	// Set default values
	bufferSize := c.BufferSize
	if bufferSize <= 0 {
		bufferSize = 0 // Default to disable buffering to avoid test issues
	}

	flushInterval := c.FlushInterval
	if flushInterval <= 0 {
		flushInterval = 5 // Default to flush every 5 seconds
	}

	// Create buffered WriteSyncer
	var syncer zapcore.WriteSyncer
	if c.LogInConsole {
		// Output to both console and file
		consoleSyncer := zapcore.AddSync(os.Stdout)
		fileSyncer := zapcore.AddSync(lumberjackLogger)
		syncer = zapcore.NewMultiWriteSyncer(consoleSyncer, fileSyncer)
	} else {
		// Output only to file
		fileSyncer := zapcore.AddSync(lumberjackLogger)
		syncer = fileSyncer
	}

	// If buffering is enabled, create custom buffered WriteSyncer
	if bufferSize > 0 {
		syncer = &BufferedWriteSyncer{
			WriteSyncer:   syncer,
			Size:          bufferSize,
			FlushInterval: flushInterval,
		}
	}

	return syncer, nil
}

// BufferedWriteSyncer implements a buffered WriteSyncer
type BufferedWriteSyncer struct {
	WriteSyncer   zapcore.WriteSyncer
	Size          int
	FlushInterval int
	buf           *bufio.Writer
	mu            sync.Mutex
	stopChan      chan struct{}
}

// Write implements the zapcore.WriteSyncer interface
func (bws *BufferedWriteSyncer) Write(p []byte) (n int, err error) {
	bws.mu.Lock()
	defer bws.mu.Unlock()

	if bws.buf == nil {
		bws.buf = bufio.NewWriterSize(bws.WriteSyncer, bws.Size)
		// Start periodic flush goroutine
		bws.stopChan = make(chan struct{})
		go bws.startFlushTicker()
	}

	return bws.buf.Write(p)
}

// Sync implements the zapcore.WriteSyncer interface
func (bws *BufferedWriteSyncer) Sync() error {
	bws.mu.Lock()
	defer bws.mu.Unlock()

	if bws.buf != nil {
		if err := bws.buf.Flush(); err != nil {
			return err
		}
	}

	return bws.WriteSyncer.Sync()
}

// startFlushTicker starts a periodic flush goroutine
func (bws *BufferedWriteSyncer) startFlushTicker() {
	ticker := time.NewTicker(time.Duration(bws.FlushInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			bws.mu.Lock()
			if bws.buf != nil {
				bws.buf.Flush()
			}
			bws.mu.Unlock()
		case <-bws.stopChan:
			return
		}
	}
}

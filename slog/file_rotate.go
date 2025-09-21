package slog

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"path"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// GetHandler get slog.Handler
func GetHandler(c *SlogConfig) (slog.Handler, error) {
	now := time.Now().Format("2006-01-02")
	filename := path.Join(c.Dir, now+"-"+c.Level+".log")
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

	var writer io.Writer
	if c.LogInConsole && c.LogInFile {
		writer = io.MultiWriter(os.Stdout, lumberjackLogger)
	} else if c.LogInConsole {
		writer = os.Stdout
	} else if c.LogInFile {
		writer = lumberjackLogger
	} else {
		// If neither is enabled, use standard output as default
		writer = os.Stdout
	}

	// If buffering is enabled, create custom buffered Writer
	if bufferSize > 0 {
		bufferedWriter := &BufferedWriter{
			Writer: writer,
			Size:   bufferSize,
		}
		writer = bufferedWriter
		go bufferedWriter.startFlushTicker(flushInterval)
	}
	switch c.Format {
	case "json":
		return slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: getSlogLevel(c.Level),
		}), nil
	default:
		return slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: getSlogLevel(c.Level),
		}), nil
	}
}

// BufferedWriter implements a buffered Writer
type BufferedWriter struct {
	Writer io.Writer
	Size   int
	buf    *bufio.Writer
	mu     sync.Mutex
}

// Write implements the io.Writer interface
func (bw *BufferedWriter) Write(p []byte) (n int, err error) {
	bw.mu.Lock()
	defer bw.mu.Unlock()

	if bw.buf == nil {
		bw.buf = bufio.NewWriterSize(bw.Writer, bw.Size)
	}

	return bw.buf.Write(p)
}

// startFlushTicker starts a periodic flush goroutine
func (bw *BufferedWriter) startFlushTicker(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		bw.mu.Lock()
		if bw.buf != nil {
			bw.buf.Flush()
		}
		bw.mu.Unlock()
	}
}

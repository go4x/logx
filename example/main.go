package main

import (
	"fmt"
	// Import example functions from the current package
)

// main demonstrates all the features of the logx library.
// It runs various examples showing different aspects of logging:
// - Basic usage with zap and slog loggers
// - Global logger functions
// - Different log levels
// - Performance optimization
// - Buffer configuration
// - Performance comparison between buffered and non-buffered logging
func main() {
	fmt.Println("=== logx Examples Demo ===")
	fmt.Println()

	// Basic usage example
	fmt.Println("1. Basic Usage Example")
	ExampleBasicUsage()
	fmt.Println()

	// Global logger example
	fmt.Println("2. Global Logger Example")
	ExampleGlobalLogger()
	fmt.Println()

	// Log levels example
	fmt.Println("3. Log Levels Example")
	ExampleLogLevels()
	fmt.Println()

	// Performance optimization example
	fmt.Println("4. Performance Optimization Example")
	ExamplePerformanceOptimization()
	fmt.Println()

	// Buffer configuration example
	fmt.Println("5. Buffer Configuration Example")
	ExampleBufferConfiguration()
	fmt.Println()

	// Buffer performance comparison example
	fmt.Println("6. Buffer Performance Comparison Example")
	ExampleBufferVsNoBuffer()
	fmt.Println()

	fmt.Println("=== All Examples Demo Completed ===")
}

package slog

// SlogConfig holds the configuration for the slog logger.
type SlogConfig struct {
	// Level specifies the minimum log level (debug, info, warn, error, fatal).
	Level string `mapstructure:"level" yaml:"level"`

	// LogInFile determines whether to write logs to file.
	LogInFile bool `mapstructure:"log-in-file" yaml:"log-in-file"`

	// LogInConsole determines whether to output logs to console.
	LogInConsole bool `mapstructure:"log-in-console" yaml:"log-in-console"`

	// Dir specifies the directory where log files are stored.
	Dir string `mapstructure:"dir" yaml:"dir"`

	// Format specifies the log output format (text or json).
	Format string `mapstructure:"format" yaml:"format"`

	// MaxAge specifies the maximum number of days to retain log files (0 means no limit).
	MaxAge int `mapstructure:"max-age" yaml:"max-age"`

	// MaxSize specifies the maximum size of a log file in MB (0 means no limit).
	MaxSize int `mapstructure:"max-size" yaml:"max-size"`

	// MaxBackups specifies the maximum number of log files to keep (0 means no limit).
	MaxBackups int `mapstructure:"max-backups" yaml:"max-backups"`

	// LocalTime determines whether to use local time (true) or UTC time (false).
	LocalTime bool `mapstructure:"local-time" yaml:"local-time"`

	// Compress determines whether to compress rotated log files.
	Compress bool `mapstructure:"compress" yaml:"compress"`

	// BufferSize specifies the buffer size for performance optimization (0 disables buffering).
	BufferSize int `mapstructure:"buffer-size" yaml:"buffer-size"`

	// FlushInterval specifies the interval in seconds to flush the buffer.
	FlushInterval int `mapstructure:"flush-interval" yaml:"flush-interval"`
}

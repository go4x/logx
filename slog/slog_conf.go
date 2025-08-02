package slog

// SlogConfig slog logger configuration
type SlogConfig struct {
	// log level
	Level string `mapstructure:"level" yaml:"level"`

	// whether to write to file
	LogInFile bool `mapstructure:"log-in-file" yaml:"log-in-file"`

	// whether to output to console
	LogInConsole bool `mapstructure:"log-in-console" yaml:"log-in-console"`

	// log directory
	Dir string `mapstructure:"dir" yaml:"dir"`

	// log format, text or json
	Format string `mapstructure:"format" yaml:"format"`

	// log retention time in days, 0 means no limit
	MaxAge int `mapstructure:"max-age" yaml:"max-age"`

	// log file size in MB, 0 means no limit
	MaxSize int `mapstructure:"max-size" yaml:"max-size"`

	// maximum number of log files to keep, 0 means no limit
	MaxBackups int `mapstructure:"max-backups" yaml:"max-backups"`

	// local time, default is true, if false, use UTC time
	LocalTime bool `mapstructure:"local-time" yaml:"local-time"`

	// compress, default is false
	Compress bool `mapstructure:"compress" yaml:"compress"`
}

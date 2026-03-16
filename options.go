package logger

import "errors"

// Option configures a logger at construction time when passed to New.
type Option func(*Options)

// Options holds configuration for New. New defaults Level to InfoLevel and Output to ConsoleOutput when not set by options.
type Options struct {
	Level       Level       // Minimum log level.
	Output      OutputType  // Where to write (console, file, or both).
	FileOptions FileOptions // Required when Output is FileOutput or BothOutput.
	ServiceName string      // Added as "service" field to every event when non-empty.
	ExitFunc    func(int)   // Called by Fatal after logging and Close; nil uses os.Exit.
}

// FileOptions configures file output (lumberjack). Filename is required for file or both output.
// Zero MaxSize, MaxBackups, MaxAge use internal defaults; set them for explicit rotation behavior.
type FileOptions struct {
	Filename   string // Log file path (required for FileOutput or BothOutput).
	MaxSize    int    // Max megabytes per file before rotation; zero uses default (100).
	MaxBackups int    // Max number of old files to keep; zero uses default (5).
	MaxAge     int    // Max days to keep old files; zero uses default (30).
	Compress   bool   // Whether to gzip rotated files.
}

// ErrEmptyFilename is returned by New when Output is FileOutput or BothOutput but FileOptions.Filename is empty.
var ErrEmptyFilename = errors.New("logger: filename required for file output")

// WithLevel sets the minimum log level (events below this level are not emitted).
func WithLevel(level Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// WithOutput sets the output destination (ConsoleOutput, FileOutput, or BothOutput).
func WithOutput(output OutputType) Option {
	return func(o *Options) {
		o.Output = output
	}
}

// WithFileOptions sets file output options. Required when Output is FileOutput or BothOutput; Filename must be set.
func WithFileOptions(fo FileOptions) Option {
	return func(o *Options) {
		o.FileOptions = fo
	}
}

// WithServiceName sets the service name; adds a "service" field to every log event. Use for multi-service deployments.
func WithServiceName(name string) Option {
	return func(o *Options) {
		o.ServiceName = name
	}
}

// WithExitFunc sets the function called by Fatal after logging and Close. Use for tests or to avoid os.Exit (e.g. graceful shutdown).
// If nil, Fatal calls os.Exit(1).
func WithExitFunc(fn func(int)) Option {
	return func(o *Options) {
		o.ExitFunc = fn
	}
}

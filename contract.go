package logger

// Level is the minimum level at which log events are emitted. Unknown values are treated as InfoLevel.
type Level int

// Log levels. DebugLevel is the lowest, FatalLevel the highest.
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// OutputType selects where log output is written (console, file, or both).
type OutputType int

// Output destinations: console only, file only, or both.
const (
	ConsoleOutput OutputType = iota
	FileOutput
	BothOutput
)

// Fields is a key-value map attached to a log event. Multiple Fields passed to one call are merged; later keys override earlier ones.
// Fields are shallow-copied when passed to WithFields or to log methods; nested maps and slices are shared and must not be mutated concurrently. Do not mutate a Fields map after passing it to a log method or WithFields.
type Fields = map[string]any

// Logger is the interface for structured logging. All level methods accept an optional variadic Fields.
// Close releases resources (e.g. file handles); typically call only on the root logger.
type Logger interface {
	// Debug logs at debug level.
	Debug(msg string, fields ...Fields)
	// Info logs at info level.
	Info(msg string, fields ...Fields)
	// Warn logs at warn level.
	Warn(msg string, fields ...Fields)
	// Error logs at error level.
	Error(msg string, fields ...Fields)
	// Fatal logs at fatal level, then closes the logger and calls the configured exit function (default os.Exit(1)). Deferred functions in the caller are not run.
	Fatal(msg string, fields ...Fields)
	// WithFields returns a logger that includes the given fields in every subsequent log event.
	WithFields(fields Fields) Logger
	// WithError returns a logger that includes err in the next log event (e.g. for structured error logging).
	WithError(err error) Logger
	// Close releases resources. Safe to call multiple times; only the first call has effect.
	Close() error
}

// Package logger provides a minimal structured logging interface backed by zerolog.
//
// # Building a logger
//
// New builds a Logger from options. Defaults are InfoLevel and ConsoleOutput. Use WithLevel, WithOutput,
// WithFileOptions, WithServiceName, and WithExitFunc to configure. Returns ErrEmptyFilename when Output is FileOutput or
// BothOutput and FileOptions.Filename is empty. Options.Level zero value is DebugLevel (0); New applies
// InfoLevel when no option sets Level. Unknown Level is mapped to InfoLevel.
//
// # Output
//
// ConsoleOutput writes to stdout (human-readable). FileOutput uses lumberjack for rotation; set FileOptions
// (Filename required; MaxSize, MaxBackups, MaxAge have defaults). BothOutput writes to both. WithFields and
// WithError return loggers that share the root's output and closer; calling Close on any of them closes the
// underlying file for all—typically only close the root logger.
//
// # Levels and methods
//
// Logger provides Debug, Info, Warn, Error, and Fatal. Each accepts a message and optional variadic Fields;
// multiple Fields are merged (later keys override). Fatal writes the event, closes the logger, and calls
// os.Exit(1); deferred functions in the caller are not run. Prefer Error plus explicit shutdown for graceful exit.
// Noop returns a logger that discards all output; Fatal on the noop logger does not exit (for tests).
//
// # Context
//
// IntoContext stores a logger in the request context; FromContext retrieves it (or Noop() if absent or nil).
// Use in handlers to get the request-scoped logger set by middleware.
//
// # Conventions
//
// Do not log secrets (passwords, tokens, API keys) or unredacted PII in Fields. Use consistent keys from
// the fields helpers: TraceID, RequestID, UserID, Error, Duration, Component. Control characters (including \r and \n)
// in message strings and in field keys and values are replaced with space to reduce log injection.
//
// # Thread safety
//
// Logger methods (Debug, Info, Warn, Error, Fatal, WithFields, WithError) are safe for concurrent use. Close is safe to call from multiple goroutines; only the first call has effect and closes the underlying output for all loggers sharing it.
package logger

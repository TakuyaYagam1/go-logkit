// Package logger provides a minimal structured logging interface backed by zerolog.
//
// Logger has level-based methods (Debug, Info, Warn, Error, Fatal), WithFields, WithError.
// Build with New and options (WithLevel, WithOutput, WithFileOptions, WithServiceName).
// Output can be console, file, or both; file output uses lumberjack for rotation.
//
// Context: IntoContext stores a logger in context; FromContext retrieves it (or Noop() if absent).
// Use TraceID, RequestID, UserID, Error, Duration, Component from the fields package for consistent keys.
// Noop returns a silent logger for tests or when logging is disabled.
package logger

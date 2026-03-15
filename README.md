# go-logkit

Structured logging interface backed by zerolog.

## Install

```bash
go get github.com/TakuyaYagam1/go-logkit
```

```go
import logger "github.com/TakuyaYagam1/go-logkit"
```

## Features

- **Logger** interface: Debug, Info, Warn, Error, Fatal; WithFields, WithError
- **Output**: console, file, or both (lumberjack rotation for file)
- **Context**: IntoContext / FromContext to pass logger in request context
- **Fields**: TraceID, RequestID, UserID, Error, Duration, Component for consistent keys
- **Options**: WithLevel, WithOutput, WithFileOptions, WithServiceName
- **Noop()**: silent logger for tests

## Example

```go
l, err := logger.New(
    logger.WithLevel(logger.InfoLevel),
    logger.WithOutput(logger.ConsoleOutput),
    logger.WithServiceName("api"),
)
if err != nil {
    log.Fatal(err)
}
l.Info("started", logger.RequestID("req-1"), logger.Duration(time.Second))

ctx := logger.IntoContext(r.Context(), l)
lFromCtx := logger.FromContext(ctx)
```

## API

| Symbol | Description |
|--------|-------------|
| Logger | Interface: Debug, Info, Warn, Error, Fatal(msg, fields...); WithFields, WithError |
| Level | DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel |
| OutputType | ConsoleOutput, FileOutput, BothOutput |
| Fields | map[string]any for log event key-value data |
| New(opts...) | Build Logger; returns ErrEmptyFilename if file output without Filename |
| Noop() | Logger that discards all output |
| IntoContext, FromContext | Store/retrieve Logger in context |
| TraceID, RequestID, UserID, Error, Duration, Component | Field helpers |
| WithLevel, WithOutput, WithFileOptions, WithServiceName | Options |

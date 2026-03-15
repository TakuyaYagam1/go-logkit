package logger

import "context"

type contextKey struct{}

// IntoContext stores the logger in ctx. Use FromContext in handlers to retrieve it.
func IntoContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

// FromContext returns the logger from ctx, or Noop() if none is set or the value is nil.
func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(contextKey{}).(Logger); ok && l != nil {
		return l
	}
	return Noop()
}

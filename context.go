package logger

import "context"

type contextKey struct{}

// IntoContext stores the logger in ctx. Use FromContext in handlers to retrieve the request-scoped logger. If ctx is nil, uses context.Background().
func IntoContext(ctx context.Context, l Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contextKey{}, l)
}

// FromContext returns the logger from ctx (set by IntoContext), or Noop() if absent or nil.
func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return Noop()
	}
	if l, ok := ctx.Value(contextKey{}).(Logger); ok && l != nil {
		return l
	}
	return Noop()
}

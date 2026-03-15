package logger

import "time"

const (
	keyTraceID   = "trace_id"
	keyRequestID = "request_id"
	keyUserID    = "user_id"
	keyError     = "error"
	keyDuration  = "duration"
	keyComponent = "component"
)

// TraceID returns Fields with the trace_id key for distributed tracing.
func TraceID(id string) Fields {
	return Fields{keyTraceID: id}
}

// RequestID returns Fields with the request_id key.
func RequestID(id string) Fields {
	return Fields{keyRequestID: id}
}

// UserID returns Fields with the user_id key.
func UserID(id string) Fields {
	return Fields{keyUserID: id}
}

// Error returns Fields with the error key, or nil if err is nil.
func Error(err error) Fields {
	if err == nil {
		return nil
	}
	return Fields{keyError: err}
}

// Duration returns Fields with the duration key.
func Duration(d time.Duration) Fields {
	return Fields{keyDuration: d}
}

// Component returns Fields with the component key (e.g. handler or service name).
func Component(name string) Fields {
	return Fields{keyComponent: name}
}

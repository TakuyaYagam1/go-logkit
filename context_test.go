package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntoContext_FromContext(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	assert.NotNil(t, FromContext(ctx))
	assert.Equal(t, Noop(), FromContext(ctx))

	l, err := New(WithLevel(InfoLevel), WithOutput(ConsoleOutput))
	require.NoError(t, err)
	ctx = IntoContext(ctx, l)
	assert.Same(t, l, FromContext(ctx))
}

func TestFromContext_NilLogger(t *testing.T) {
	t.Parallel()
	var nilLogger Logger
	ctx := context.WithValue(context.Background(), contextKey{}, nilLogger)
	assert.Equal(t, Noop(), FromContext(ctx))
}

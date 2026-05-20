package tx

import (
	"context"
)

type contextKey struct{}

func FactoryFromContext(ctx context.Context) (Factory, bool) {
	factory, ok := ctx.Value(contextKey{}).(Factory)
	return factory, ok
}

func WithFactory(ctx context.Context, factory Factory) context.Context {
	return context.WithValue(ctx, contextKey{}, factory)
}

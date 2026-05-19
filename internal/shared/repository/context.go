package repository

import "context"

type contextKey struct{}

func WithProvider(ctx context.Context, provider Provider) context.Context {
	return context.WithValue(ctx, contextKey{}, provider)
}

func GetProvider(ctx context.Context) (Provider, bool) {
	p, ok := ctx.Value(contextKey{}).(Provider)
	return p, ok
}

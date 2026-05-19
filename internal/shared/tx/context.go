package tx

import (
	"context"
)

type contextKey struct{}

func WithTransaction(ctx context.Context, tx Transaction) context.Context {
	return context.WithValue(
		ctx,
		contextKey{},
		tx,
	)
}

func GetTransaction(ctx context.Context) (Transaction, bool) {
	t, ok := ctx.Value(contextKey{}).(Transaction)

	return t, ok
}

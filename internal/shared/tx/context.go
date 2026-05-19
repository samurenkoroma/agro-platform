package tx

import (
	"context"
)

type ctxKey struct{}
type contextKey struct{}
type factoryKey struct{}

func WithUnitOfWork(ctx context.Context, uow UnitOfWork) context.Context {
	return context.WithValue(ctx, ctxKey{}, uow)
}
func FromContext(ctx context.Context) (UnitOfWork, bool) {
	uow, ok := ctx.Value(ctxKey{}).(UnitOfWork)
	return uow, ok
}

func FactoryFromContext(ctx context.Context) (Factory, bool) {
	factory, ok := ctx.Value(factoryKey{}).(Factory)
	return factory, ok
}

func WithFactory(ctx context.Context, factory Factory) context.Context {
	return context.WithValue(ctx, factoryKey{}, factory)
}

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

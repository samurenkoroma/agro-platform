package query

import "context"

type Handler[TQuery Query, TResult any] interface {
	Handle(ctx context.Context, q TQuery) (TResult, error)
}

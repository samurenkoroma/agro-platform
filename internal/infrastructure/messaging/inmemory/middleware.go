package inmemory

import "context"

type Next func(ctx context.Context, msg any) error

type Middleware func(ctx context.Context, msg any, next Next) error

func Chain(mw []Middleware, final Next) Next {

	if len(mw) == 0 {
		return final
	}

	next := final

	for i := len(mw) - 1; i >= 0; i-- {
		current := mw[i]
		prev := next

		next = func(ctx context.Context, msg any) error {
			return current(ctx, msg, prev)
		}
	}

	return next
}

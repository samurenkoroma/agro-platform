package query

import "context"

type Query interface {
	QueryName() string
}

type Bus interface {
	Ask(ctx context.Context, q Query) (any, error)
}

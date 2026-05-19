package inmemory

import (
	"context"
	"fmt"

	qb "github.com/samurenkoroma/agro-platform/internal/shared/bus/query"
)

type QueryBus struct {
	registry *Registry
}

func NewQueryBus(reg *Registry) *QueryBus {
	return &QueryBus{registry: reg}
}

func (b *QueryBus) Ask(ctx context.Context, q qb.Query) (any, error) {
	name := q.QueryName()

	raw, ok := b.registry.Get(name)

	if !ok {
		return nil, qb.ErrQueryHandlerNotFound
	}

	handler, ok := raw.(func(context.Context, qb.Query) (any, error))

	if !ok {
		return nil, fmt.Errorf("invalid query handler %s", name)
	}

	return handler(ctx, q)
}

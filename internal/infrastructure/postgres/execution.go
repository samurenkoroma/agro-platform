package postgres

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
)

type execution struct {
	ctx context.Context

	aggregates []aggregate.Aggregate
}

func (e *execution) RegisterAggregate(agg aggregate.Aggregate) {
	e.aggregates = append(e.aggregates, agg)
}

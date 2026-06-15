package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

var (
	ErrAlreadyCommitted  = errors.New("unit of work already committed")
	ErrAlreadyRolledBack = errors.New("unit of work already rolled back")
)

type unitOfWork struct {
	pool *pgxpool.Pool
	bus  bus.EventBus
}

func NewUnitOfWork(pool *pgxpool.Pool, bus bus.EventBus) uow.UnitOfWork {
	return &unitOfWork{
		pool: pool,
		bus:  bus,
	}
}

func (u *unitOfWork) Execute(ctx context.Context, build func(db uow.DB) repository.RepositoryProvider, fn func(provider repository.RepositoryProvider, exec uow.Execution) (any, error)) (any, error) {

	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	exec := &execution{
		ctx:        ctx,
		aggregates: make([]aggregate.Aggregate, 0),
	}

	provider := build(tx)

	data, err := fn(provider, exec)
	if err != nil {

		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return nil, fmt.Errorf("rollback error: %v, original error: %w", rbErr, err)
		}

		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	if err := u.dispatchEvents(ctx, exec); err != nil {
		return nil, err
	}

	return data, nil
}

func (u *unitOfWork) dispatchEvents(ctx context.Context, exec *execution) error {

	for {

		var events []event.DomainEvent

		for _, agg := range exec.aggregates {
			events = append(events, agg.PullEvents()...)
		}

		exec.aggregates = nil

		if len(events) == 0 {
			return nil
		}

		if err := u.bus.Publish(ctx, events); err != nil {
			return err
		}
	}
}

package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
	"github.com/samurenkoroma/agro-platform/pkg/logger"
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
	log := logger.FromContext(ctx)
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	exec := &execution{
		ctx:        ctx,
		aggregates: make([]aggregate.Aggregate, 0),
	}

	provider := build(tx)

	log.Debug("uow: transaction started", "provider", provider.ProviderName())

	data, err := fn(provider, exec)
	if err != nil {

		if rbErr := tx.Rollback(ctx); rbErr != nil {
			log.Error("uow: rollback failed",
				"rollback_error", rbErr,
				"original_error", err,
			)

			return nil, fmt.Errorf("rollback error: %v, original error: %w", rbErr, err)
		}
		log.Warn("uow: transaction rolled back", "error", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error("uow: commit failed", "error", err)
		return nil, err
	}
	log.Debug("uow: transaction committed")
	if err := u.dispatchEvents(ctx, exec); err != nil {
		return nil, err
	}
	log.Debug("uow: events dispatched")
	return data, nil
}

func (u *unitOfWork) dispatchEvents(ctx context.Context, exec *execution) error {
	log := logger.FromContext(ctx)
	for {

		var events []event.DomainEvent

		for _, agg := range exec.aggregates {
			events = append(events, agg.PullEvents()...)
		}

		exec.aggregates = nil

		if len(events) == 0 {
			return nil
		}
		log.Debug("uow: dispatching events", "count", len(events))

		for _, e := range events {
			log.Debug("uow: event dispatched",
				"event_type", e.EventType(),
				"aggregate_id", e.AggregateID(),
			)
		}

		if err := u.bus.Publish(ctx, events); err != nil {
			return err
		}
	}
}

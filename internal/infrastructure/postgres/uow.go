package postgres

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
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
	committed  bool
	rolledBack bool

	mu         sync.Mutex
	aggregates []aggregate.Aggregate
	pool       *pgxpool.Pool
	ctx        context.Context
	bus        bus.EventBus
}

func NewUnitOfWork(ctx context.Context, pool *pgxpool.Pool, bus bus.EventBus) uow.UnitOfWork {
	return &unitOfWork{
		pool: pool,
		ctx:  ctx,
		bus:  bus,
	}
}

func (u *unitOfWork) Execute(ctx context.Context, build func(db uow.DB) repository.RepositoryProvider, fn func(provider repository.RepositoryProvider) (any, error)) (any, error) {
	// Создаем провайдер для этой транзакции
	provider := build(u.pool)

	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	// Выполняем бизнес-логику
	data, err := fn(provider)
	if err != nil {
		// В случае ошибки — откат
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return nil, fmt.Errorf("rollback error: %v, original error: %w", rbErr, err)
		}
		return nil, err
	}

	// Пробуем закоммитить
	if err := u.Commit(tx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return data, nil
}

func (u *unitOfWork) RegisterAggregate(agg aggregate.Aggregate) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.aggregates = append(u.aggregates, agg)
}

func (u *unitOfWork) Commit(tx pgx.Tx) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.committed {
		return ErrAlreadyCommitted
	}
	if u.rolledBack {
		return ErrAlreadyRolledBack
	}

	if err := tx.Commit(u.ctx); err != nil {
		return err
	}

	return u.dispatchEvents()
}

func (u *unitOfWork) Rollback(tx pgx.Tx) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.committed {
		return ErrAlreadyCommitted
	}
	if u.rolledBack {
		return ErrAlreadyRolledBack
	}
	u.rolledBack = true
	return tx.Rollback(u.ctx)
}

func (u *unitOfWork) dispatchEvents() error {
	var allEvents []event.DomainEvent

	for _, agg := range u.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return u.bus.Publish(u.ctx, allEvents)
}

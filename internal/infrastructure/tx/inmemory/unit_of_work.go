package inmemory

import (
	"context"

	de "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"

	tx "github.com/samurenkoroma/agro-platform/internal/shared/tx"
)

type UnitOfWork struct {
	transaction tx.Transaction
	events      []de.DomainEvent
}

func NewUnitOfWork(transaction tx.Transaction) *UnitOfWork {
	return &UnitOfWork{
		transaction: transaction,

		events: make([]de.DomainEvent, 0),
	}
}

func (u *UnitOfWork) Transaction() tx.Transaction {
	return u.transaction
}

func (u *UnitOfWork) AddEvents(events ...de.DomainEvent) {
	u.events = append(u.events, events...)
}

func (u *UnitOfWork) Events() []de.DomainEvent {
	return u.events
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	return u.transaction.Commit(ctx)
}

func (u *UnitOfWork) Rollback(ctx context.Context) error {
	return u.transaction.Rollback(ctx)
}

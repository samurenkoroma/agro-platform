package tx

import (
	"context"

	de "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
)

type UnitOfWork interface {
	Transaction() Transaction
	AddEvents(events ...de.DomainEvent)
	Events() []de.DomainEvent
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

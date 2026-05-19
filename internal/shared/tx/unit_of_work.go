package tx

import (
	"context"
	"database/sql"

	de "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type UnitOfWork interface {
	Execute(ctx context.Context, build func(tx *sql.Tx) repository.RepositoryProvider, fn func(provider repository.RepositoryProvider) (any, error)) (any, error)
	AddEvents(events ...de.DomainEvent)
	Events() []de.DomainEvent
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

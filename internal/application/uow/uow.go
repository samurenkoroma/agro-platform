package uow

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type UnitOfWork interface {
	Execute(
		ctx context.Context,
		build func(DB) repository.RepositoryProvider,
		fn func(provider repository.RepositoryProvider, exec Execution) (any, error),
	) (any, error)
}

type Execution interface {
	RegisterAggregate(aggregate.Aggregate)
}

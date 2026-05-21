package uow

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/provider"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type UnitOfWork interface {
	Execute(context.Context, provider.ProviderDeps, func(provider repository.RepositoryProvider) (any, error)) (any, error)
	RegisterAggregate(agg aggregate.Aggregate)
}

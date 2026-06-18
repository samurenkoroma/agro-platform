package testutil

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

// FakeUoW исполняет fn сразу с переданным провайдером — без БД и транзакций.
type FakeUoW struct {
	Provider repository.RepositoryProvider
}

func (f *FakeUoW) Execute(
	ctx context.Context,
	_ func(db uow.DB) repository.RepositoryProvider,
	fn func(repository.RepositoryProvider) (any, error),
) (any, error) {
	return fn(f.Provider)
}

func (f *FakeUoW) RegisterAggregate(_ aggregate.Aggregate) {}

var _ uow.UnitOfWork = (*FakeUoW)(nil)

package testutil

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

// fakeExecution — минимальная реализация uow.Execution для тестов.
// Собирает зарегистрированные агрегаты и пуллит их события после выполнения,
// как это делает настоящий UoW после commit.
type fakeExecution struct {
	aggregates []aggregate.Aggregate
}

func (e *fakeExecution) RegisterAggregate(agg aggregate.Aggregate) {
	e.aggregates = append(e.aggregates, agg)
}

// FakeUoW исполняет fn сразу с переданным провайдером — без БД и транзакций.
// События зарегистрированных агрегатов дренируются (PullEvents), но никуда
// не публикуются — если тесту нужно проверить факт публикации события,
// используйте настоящий inmemory.EventBus напрямую в application-тесте.
type FakeUoW struct {
	Provider repository.RepositoryProvider
}

func (f *FakeUoW) Execute(
	ctx context.Context,
	_ func(db uow.DB) repository.RepositoryProvider,
	fn func(provider repository.RepositoryProvider, exec uow.Execution) (any, error),
) (any, error) {
	exec := &fakeExecution{}
	data, err := fn(f.Provider, exec)
	if err != nil {
		return nil, err
	}
	for _, agg := range exec.aggregates {
		agg.PullEvents() // дренируем, чтобы повторное использование агрегата не копило старые события
	}
	return data, nil
}

var _ uow.UnitOfWork = (*FakeUoW)(nil)

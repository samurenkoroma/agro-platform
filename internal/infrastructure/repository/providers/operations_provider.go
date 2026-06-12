package providers

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	inmemory "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/operations"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/operations"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type operationsProvider struct {
	db         uow.DB
	inMemory   bool
	tasks      domain.TaskRepository
	timelines  domain.TimeLineRepository
	operations domain.OperationRepository
}

func NewOperationsProvider(db uow.DB) repository.RepositoryProvider {
	return &operationsProvider{db: db, inMemory: false}
}

func (p *operationsProvider) ProviderName() string { return "operations" }

func (p *operationsProvider) Tasks() domain.TaskRepository {
	if p.tasks != nil {
		return p.tasks
	}
	if p.inMemory {
		p.tasks = inmemory.NewTaskRepository()
	} else {
		p.tasks = postgres.NewTaskRepository(p.db)
	}
	return p.tasks
}

func (p *operationsProvider) Timelines() domain.TimeLineRepository {
	if p.timelines != nil {
		return p.timelines
	}
	if p.inMemory {
		p.timelines = inmemory.NewTimelineRepository()
	} else {
		p.timelines = postgres.NewTimelineRepository(p.db)
	}
	return p.timelines
}

func (p *operationsProvider) Operations() domain.OperationRepository {
	if p.operations != nil {
		return p.operations
	}
	if p.inMemory {
		p.operations = inmemory.NewOperationRepository()
	} else {
		p.operations = postgres.NewOperationRepository(p.db)
	}
	return p.operations
}

var _ domain.OperationsProvider = (*operationsProvider)(nil)

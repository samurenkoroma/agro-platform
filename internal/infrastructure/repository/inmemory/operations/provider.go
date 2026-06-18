package operations

import (
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
)

type provider struct {
	tasks      repository.TaskRepository
	timelines  repository.TimeLineRepository
	operations repository.OperationRepository
}

func NewProvider() repository.OperationsProvider {
	return &provider{
		tasks:      NewTaskRepository(),
		timelines:  NewTimelineRepository(),
		operations: NewOperationRepository(),
	}
}

func (p *provider) Tasks() repository.TaskRepository           { return p.tasks }
func (p *provider) Timelines() repository.TimeLineRepository   { return p.timelines }
func (p *provider) Operations() repository.OperationRepository { return p.operations }
func (p *provider) ProviderName() string                       { return "operations_inmemory" }

var _ repository.OperationsProvider = (*provider)(nil)

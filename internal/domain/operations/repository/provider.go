package repository

import "github.com/samurenkoroma/agro-platform/internal/shared/repository"

type OperationsProvider interface {
	repository.RepositoryProvider
	Tasks() TaskRepository
	Timelines() TimeLineRepository
	Operations() OperationRepository
}

package repository

type OperationsProvider interface {
	Tasks() TaskRepository
	Timelines() TimeLineRepository
	Operations() OperationRepository
}

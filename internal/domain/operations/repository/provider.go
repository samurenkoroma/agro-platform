package repository

type OperationsProvider interface {
	Tasks() TaskRepository
}

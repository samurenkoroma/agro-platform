package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"

type taskRepository struct {
}

func NewTaskRepository() repository.TaskRepository {
	return &taskRepository{}
}

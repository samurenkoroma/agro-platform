package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type TaskFilter struct {
	FarmID           vo.ID
	GrowingCycleID   *vo.ID
	ProductionUnitID *vo.ID
	Status           *task.Status
	AssignedTo       *vo.ID
}

type TaskRepository interface {
	Save(ctx context.Context, t *task.Task) error
	GetByID(ctx context.Context, id vo.ID) (*task.Task, error)
	List(ctx context.Context, filter TaskFilter) ([]*task.Task, error)
	Delete(ctx context.Context, id vo.ID) error
}

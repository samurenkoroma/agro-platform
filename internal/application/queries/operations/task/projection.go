package task

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type TaskDTO struct {
	ID               vo.ID      `json:"id"`
	Title            string     `json:"title"`
	Description      *string    `json:"description,omitempty"`
	OperationType    *string    `json:"operationType,omitempty"`
	ProductionUnitID *vo.ID     `json:"productionUnitId,omitempty"`
	GrowingCycleID   *vo.ID     `json:"growingCycleId,omitempty"`
	AssignedTo       *string    `json:"assignedTo,omitempty"`
	Status           string     `json:"status"`
	Priority         string     `json:"priority"`
	DueDate          *time.Time `json:"dueDate,omitempty"`
	CompletedAt      *time.Time `json:"completedAt,omitempty"`
	CreatedAt        time.Time  `json:"createdAt"`
}

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*TaskDTO, error)
	List(ctx context.Context, farmID vo.ID, cycleID *vo.ID) ([]*TaskDTO, error)
}

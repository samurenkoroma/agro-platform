package repository

import (
	"context"
	"time"

	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type OperationFilter struct {
	FarmID           vo.ID
	GrowingCycleID   *vo.ID
	ProductionUnitID *vo.ID
	Type             *operationevent.OperationType
	From             *time.Time
	To               *time.Time
}

type OperationRepository interface {
	Save(ctx context.Context, e *operationevent.OperationEvent) error
	GetByID(ctx context.Context, id vo.ID) (*operationevent.OperationEvent, error)
	List(ctx context.Context, filter OperationFilter) ([]*operationevent.OperationEvent, error)
}

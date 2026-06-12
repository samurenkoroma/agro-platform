package operationevent

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type OperationDTO struct {
	ID               vo.ID          `json:"id"`
	Type             string         `json:"type"`
	ProductionUnitID *vo.ID         `json:"productionUnitId,omitempty"`
	GrowingCycleID   *vo.ID         `json:"growingCycleId,omitempty"`
	PerformedBy      *vo.ID         `json:"performedBy,omitempty"`
	Payload          map[string]any `json:"payload"`
	Timestamp        time.Time      `json:"timestamp"`
}

type Projection interface {
	List(ctx context.Context, farmID vo.ID, cycleID *vo.ID) ([]*OperationDTO, error)
}

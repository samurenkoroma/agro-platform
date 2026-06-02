package allocation

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type DTO struct {
	ID                 vo.ID      `json:"id"`
	CycleID            vo.ID      `json:"cycleID"`
	ProductionUnitID   vo.ID      `json:"productionUnitID"`
	ProductionUnitName string     `json:"productionUnitName"`
	Area               float64    `json:"area"`
	StartedAt          time.Time  `json:"startedAt"`
	EndedAt            *time.Time `json:"endedAt"`
}

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*DTO, error)
	List(ctx context.Context, ownerId vo.ID) ([]*DTO, error)
}

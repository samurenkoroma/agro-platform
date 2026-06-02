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

type OccupancyDTO struct {
	ProductionUnitID   vo.ID  `json:"production_unit_id"`
	ProductionUnitName string `json:"production_unit_name"`

	TotalArea     float64 `json:"total_area"`
	AllocatedArea float64 `json:"allocated_area"`
	FreeArea      float64 `json:"free_area"`
}

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*DTO, error)
	List(ctx context.Context, ownerId vo.ID) ([]*DTO, error)
	ListByCycleID(ctx context.Context, cycleID vo.ID) ([]*DTO, error)

	ListByProductionUnitID(ctx context.Context, productionUnitID vo.ID) ([]*DTO, error)

	GetOccupancy(ctx context.Context, productionUnitID vo.ID) (*OccupancyDTO, error)
}

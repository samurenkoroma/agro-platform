package growingcycle

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type DTO struct {
	CropName      string          `json:"cropName"`
	AllocatedArea float64         `json:"allocatedArea"`
	TasksCount    *int            `json:"tasksCount"`
	Progress      int             `json:"progress"`
	Count         int             `json:"count"`
	Allocations   []AllocationDTO `json:"allocations"`
}

type AllocationDTO struct {
	CycleName          string     `json:"cycleName"`
	CycleId            string     `json:"cycleId"`
	VarietyName        *string    `json:"varietyName"`
	Status             string     `json:"status"`
	Stage              string     `json:"stage"`
	ProductionUnitId   *string    `json:"productionUnitId"`
	ProductionUnitName *string    `json:"productionUnitName"`
	Area               *float64   `json:"area"`
	Progress           int        `json:"progress"`
	DaysToMaturity     int        `json:"daysToMaturity"`
	StartDate          *time.Time `json:"startDate"`
	EndDate            *time.Time `json:"endDate"`
}

type SummaryDTO struct {
	ID                vo.ID
	Name              string
	Status            string
	AllocatedArea     float64
	PlantedQuantity   float64
	HarvestedQuantity float64
	//ExpectedHarvestAt *time.Time
}

type Projection interface {
	//Get(ctx context.Context, id vo.ID) (*DTO, error)
	List(ctx context.Context, ownerId vo.ID) ([]*DTO, error)
	Summary(ctx context.Context, ownerId vo.ID, cycleId vo.ID) (*SummaryDTO, error)
}

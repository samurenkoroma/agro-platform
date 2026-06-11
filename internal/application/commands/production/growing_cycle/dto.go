package growingcycle

import (
	"time"

	growingcycle "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CreateCommand struct {
	CropID    vo.ID  `json:"cropID" validate:"required"`
	VarietyID *vo.ID `json:"varietyID,omitempty"`

	Name       string                        `json:"name" validate:"required"`
	Code       string                        `json:"code" validate:"required"`
	Method     growingcycle.ProductionMethod `json:"method" validate:"required"`
	ProtocolID *vo.ID                        `json:"protocolID,omitempty"`
}
type StartGrowingCycleCMD struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	CropID vo.ID  `json:"cropID"`

	VarietyID  *vo.ID                        `json:"varietyID"`
	ProtocolID *vo.ID                        `json:"protocolID"`
	Status     growingcycle.CycleStatus      `json:"status"`
	Stage      growingcycle.CycleStage       `json:"stage"`
	Method     growingcycle.ProductionMethod `json:"method"`

	ExpectedHarvestAt *time.Time `json:"expectedHarvestAt"`

	Allocations []AllocationDTO `json:"allocations"`
	Plantings   []PlantingDTO   `json:"plantings"`
}

type AllocationDTO struct {
	ProductionUnitID vo.ID     `json:"productionUnitID"`
	Area             float64   `json:"area"`
	StartedAt        time.Time `json:"startedAt"`
}

type PlantingDTO struct {
	PlantedAt time.Time `json:"plantedAt"`
	Quantity  float64   `json:"quantity"`
}

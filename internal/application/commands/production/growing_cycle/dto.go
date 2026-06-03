package growingcycle

import (
	"time"

	growingcycle "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CreateCommand struct {
	ProductionUnitID  vo.ID      `json:"productionUnitID" validate:"required"`
	CropID            vo.ID      `json:"cropID" validate:"required"`
	Area              float64    `json:"area" validate:"required"`
	Name              string     `json:"name" validate:"required"`
	Code              string     `json:"code" validate:"required"`
	Method            string     `json:"method" validate:"required"`
	VarietyID         *vo.ID     `json:"varietyID,omitempty"`
	ProtocolID        *vo.ID     `json:"protocolID,omitempty"`
	StartedAt         *time.Time `json:"startedAt,omitempty"`
	ExpectedHarvestAt *time.Time `json:"expectedHarvestAt,omitempty"`
}

type UpdateCommand struct {
	ID                vo.ID                    `json:"id" validate:"required"`
	CropID            vo.ID                    `json:"cropID" validate:"required"`
	VarietyID         *vo.ID                   `json:"varietyID,omitempty"`
	ProtocolID        *vo.ID                   `json:"protocolID,omitempty"`
	Name              string                   `json:"name" validate:"required"`
	Code              string                   `json:"code" validate:"required"`
	Method            string                   `json:"method" validate:"required"`
	Status            growingcycle.CycleStatus `json:"status"`
	Stage             growingcycle.CycleStage  `json:"stage"`
	ExpectedHarvestAt *time.Time               `json:"expectedHarvestAt"`
}

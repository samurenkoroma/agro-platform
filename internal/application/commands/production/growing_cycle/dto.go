package growingcycle

import (
	"time"

	growingcycle "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CreateCommand struct {
	CropID            vo.ID      `json:"cropId" validate:"required"`
	VarietyID         *vo.ID     `json:"varietyId,omitempty"`
	ProtocolID        *vo.ID     `json:"protocolId,omitempty"`
	Name              string     `json:"name" validate:"required"`
	Code              string     `json:"code" validate:"required"`
	Method            string     `json:"method" validate:"required"`
	ExpectedHarvestAt *time.Time `json:"expectedHarvestAt"`
}

type UpdateCommand struct {
	ID                vo.ID                    `json:"id" validate:"required"`
	CropID            vo.ID                    `json:"cropId" validate:"required"`
	VarietyID         *vo.ID                   `json:"varietyId,omitempty"`
	ProtocolID        *vo.ID                   `json:"protocolId,omitempty"`
	Name              string                   `json:"name" validate:"required"`
	Code              string                   `json:"code" validate:"required"`
	Method            string                   `json:"method" validate:"required"`
	Status            growingcycle.CycleStatus `json:"status"`
	Stage             growingcycle.CycleStage  `json:"stage"`
	ExpectedHarvestAt *time.Time               `json:"expectedHarvestAt"`
}

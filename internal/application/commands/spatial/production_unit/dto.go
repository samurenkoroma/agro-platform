package productionunit

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type CreateCommand struct {
	Code         string                  `json:"code" validate:"required"`
	Type         pu.ProductionUnitType   `json:"type" validate:"required"`
	Status       pu.ProductionUnitStatus `json:"status" validate:"required"`
	ParentID     *vo.ID                  `json:"parentId,omitempty"`
	Capabilities []string                `json:"capabilities,omitempty"`
	Dimensions   *pu.Dimensions          `json:"dimensions,omitempty"`
}

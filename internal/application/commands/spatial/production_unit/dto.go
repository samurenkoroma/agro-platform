package productionunit

import (
	"encoding/json"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type CreateCommand struct {
	Type         pu.ProductionUnitType   `json:"type" validate:"required"`
	Status       pu.ProductionUnitStatus `json:"status" validate:"required"`
	ParentID     *vo.ID                  `json:"parentId,omitempty"`
	Capabilities []string                `json:"capabilities,omitempty"`
	Name         *string                 `json:"name"`
	Dimensions   *pu.Dimensions          `json:"dimensions,omitempty"`
	CreateChild  bool                    `json:"createChild,omitempty"`
}

type UpdateCommand struct {
	Id     vo.ID           `json:"id"`
	Schema json.RawMessage `json:"schema,omitempty"`
}

type ConfigureCommand struct {
	Id     vo.ID           `json:"id"`
	Schema pu.LayoutSchema `json:"schema"`
}

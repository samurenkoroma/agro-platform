package createproductionunit

import (
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const CommandName = "spatial.create_production_unit"

type Command struct {
	FarmID   vo.ID
	Name     string
	Type     pu.ProductionUnitType
	ParentID *vo.ID
}

func (Command) CommandName() string {
	return CommandName
}

package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type ProductionUnitRepository interface {
	Save(aggregate *pu.ProductionUnit) error
	GetByID(id vo.ID) (*pu.ProductionUnit, error)
	GetChildren(parentID vo.ID) ([]*pu.ProductionUnit, error)
	Delete(id vo.ID) error
}

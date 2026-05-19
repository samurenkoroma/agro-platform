package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type ProductionUnitRepository interface {
	Save(*pu.ProductionUnit) error
	GetByID(vo.ID) (*pu.ProductionUnit, error)
	GetChildren(vo.ID) ([]*pu.ProductionUnit, error)
	Delete(vo.ID) error
}

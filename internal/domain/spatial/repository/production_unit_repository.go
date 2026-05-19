package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type ProductionUnitRepository interface {
	Save(aggregate *pu.Aggregate) error
	GetByID(id vo.ID) (*pu.Aggregate, error)
	GetChildren(parentID vo.ID) ([]*pu.Aggregate, error)
	Delete(id vo.ID) error
}

package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
)

type GrowingCycleRepository interface {
	Save(aggregate *gc.Aggregate) error
	GetByID(id vo.ID) (*gc.Aggregate, error)
	GetActiveByUnit(unitID vo.ID) ([]*gc.Aggregate, error)
	Delete(id vo.ID) error
}

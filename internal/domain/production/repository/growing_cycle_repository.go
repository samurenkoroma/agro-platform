package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
)

type GrowingCycleRepository interface {
	Save(aggregate *gc.GrowingCycle) error
	GetByID(id vo.ID) (*gc.GrowingCycle, error)
	GetActiveByUnit(unitID vo.ID) ([]*gc.GrowingCycle, error)
	Delete(id vo.ID) error
}

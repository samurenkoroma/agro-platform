package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/slot"
)

type SlotRepository interface {
	Save(aggregate *slot.Aggregate) error
	GetByID(id vo.ID) (*slot.Aggregate, error)
	GetByUnit(unitID vo.ID) ([]*slot.Aggregate, error)
}

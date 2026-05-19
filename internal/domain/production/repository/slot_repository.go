package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/slot"
)

type SlotRepository interface {
	Save(aggregate *slot.Slot) error
	GetByID(id vo.ID) (*slot.Slot, error)
	GetByUnit(unitID vo.ID) ([]*slot.Slot, error)
}

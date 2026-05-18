package slot

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Slot struct {
	ID               vo.ID
	ProductionUnitID vo.ID
	Code             string
	Position         *vo.Coordinates
	Status           SlotStatus
	Capacity         Capacity
	Metadata         vo.Metadata
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Slot
}

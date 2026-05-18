package timeline

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Timeline struct {
	ID               vo.ID
	FarmID           vo.ID
	ProductionUnitID *vo.ID
	GrowingCycleID   *vo.ID
	PlantID          *vo.ID
	Items            []Item
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Timeline
}

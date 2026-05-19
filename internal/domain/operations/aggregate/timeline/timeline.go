package timeline

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Timeline struct {
	ev.AggregateRoot
	ID               vo.ID
	FarmID           vo.ID
	ProductionUnitID *vo.ID
	GrowingCycleID   *vo.ID
	PlantID          *vo.ID
	Items            []Item
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func New(farmID vo.ID) *Timeline {

	now := time.Now()
	root := &Timeline{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Items:     make([]Item, 0),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewTimelineCreated(root.ID))

	return root
}

func (a *Timeline) AddItem(item Item) {
	a.Items = append(a.Items, item)

	a.UpdatedAt = time.Now()

	a.AddEvent(NewItemAdded(a.ID, item.ID))
}

func (a *Timeline) Archive() {
	a.AddEvent(NewTimelineArchived(a.ID))
}

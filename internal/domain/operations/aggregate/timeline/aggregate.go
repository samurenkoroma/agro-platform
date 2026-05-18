package timeline

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID) *Aggregate {

	now := time.Now()
	root := Timeline{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Items:     make([]Item, 0),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewTimelineCreated(root.ID))

	return a
}

func (a *Aggregate) AddItem(item Item) {
	a.Root.Items = append(a.Root.Items, item)

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewItemAdded(a.Root.ID, item.ID))
}

func (a *Aggregate) Archive() {
	a.AddEvent(NewTimelineArchived(a.Root.ID))
}

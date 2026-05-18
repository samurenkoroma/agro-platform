package warehouse

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, name string) *Aggregate {
	now := time.Now()

	root := Warehouse{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Name:      name,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewWarehouseCreated(root.ID))

	return a
}

func (a *Aggregate) Archive() {
	now := time.Now()

	a.Root.ArchivedAt = &now
	a.Root.UpdatedAt = now

	a.AddEvent(NewWarehouseArchived(a.Root.ID))
}

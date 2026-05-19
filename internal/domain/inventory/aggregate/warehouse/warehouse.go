package warehouse

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Warehouse struct {
	ev.AggregateRoot
	ID         vo.ID
	FarmID     vo.ID
	Name       string
	Code       *string
	Metadata   vo.Metadata
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

func New(farmID vo.ID, name string) *Warehouse {
	now := time.Now()

	root := &Warehouse{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Name:      name,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewWarehouseCreated(root.ID))

	return root
}

func (a *Warehouse) Archive() {
	now := time.Now()

	a.ArchivedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewWarehouseArchived(a.ID))
}

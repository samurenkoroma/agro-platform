package inventoryitem

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Item struct {
	ID vo.ID

	Name string

	Type Type

	Unit Unit

	WarehouseID *vo.ID

	Stock Stock

	Metadata vo.Metadata

	CreatedAt time.Time

	UpdatedAt time.Time

	ArchivedAt *time.Time
}

type Aggregate struct {
	ev.AggregateRoot

	Root Item
}

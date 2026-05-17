package layoutsnapshot

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pus "github.com/samurenkoroma/agro-platform/internal/domain/spatial/entity/production_unit_snapshot"
)

type LayoutSnapshot struct {
	ID vo.ID

	FarmID vo.ID

	Version int

	Description *string

	CreatedBy vo.ID

	CreatedAt time.Time

	Units []pus.ProductionUnitSnapshot
}

type Aggregate struct {
	ev.AggregateRoot

	Root LayoutSnapshot
}

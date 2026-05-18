package harvestbatch

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type HarvestBatch struct {
	ID               vo.ID
	GrowingCycleID   vo.ID
	ProductionUnitID vo.ID
	Quantity         vo.Quantity
	HarvestedArea    *vo.Area
	Grade            QualityGrade
	Marketable       bool
	Notes            *string
	HarvestedAt      time.Time
	Metadata         vo.Metadata
	CreatedAt        time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root HarvestBatch
}

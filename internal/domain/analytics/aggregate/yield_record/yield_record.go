package yieldrecord

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type YieldRecord struct {
	ID vo.ID

	FarmID vo.ID

	GrowingCycleID vo.ID

	ProductionUnitID vo.ID

	HarvestBatchID *vo.ID

	YieldBatchID *vo.ID

	Source Source

	Quantity float64

	Unit string

	HarvestedAt time.Time

	Metadata vo.Metadata
}

type Aggregate struct {
	ev.AggregateRoot

	Root YieldRecord
}

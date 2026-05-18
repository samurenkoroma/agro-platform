package yieldbatch

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type YieldBatch struct {
	ID             vo.ID
	GrowingCycleID vo.ID
	PlantID        vo.ID
	Quantity       vo.Quantity
	FruitCount     *int
	Grade          QualityGrade
	Marketable     bool
	Notes          *string
	HarvestedAt    time.Time
	Metadata       vo.Metadata
	CreatedAt      time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root YieldBatch
}

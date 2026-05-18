package cropstage

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropStage struct {
	ID           vo.ID
	CropID       vo.ID
	Code         string
	Name         string
	Order        int
	BBCHCode     *string
	ExpectedDays *int
	ExpectedGDD  *float64
	Metadata     vo.Metadata
	CreatedAt    time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root CropStage
}

package cropstage

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropStage struct {
	ID           vo.ID
	CropID       vo.ID
	Code         string
	Name         string
	Order        int
	BBCH         *string
	DurationDays *int
	ExpectedGDD  *float64
	Metadata     vo.Metadata
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ArchivedAt   *time.Time
}

type Aggregate struct {
	ev.BaseAggregate
	Root CropStage
}

package cropprotocol

import (
	"time"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropProtocol struct {
	ID            vo.ID
	CropID        vo.ID
	Name          string
	GrowingMethod gc.GrowingMethod
	Description   *string
	StageProfiles []StageProfile
	Metadata      vo.Metadata
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root CropProtocol
}

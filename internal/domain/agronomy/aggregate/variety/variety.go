package variety

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Variety struct {
	ID         vo.ID
	CropID     vo.ID
	Name       string
	Breeder    *string
	Maturity   MaturityProfile
	Growth     GrowthProfile
	Spacing    PlantSpacing
	Harvest    HarvestProfile
	Yield      YieldPotential
	Tolerance  EnvironmentTolerance
	Metadata   vo.Metadata
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Variety
}

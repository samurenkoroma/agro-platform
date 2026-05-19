package growingcycle

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type GrowingCycle struct {
	ID vo.ID

	FarmID vo.ID

	CropID vo.ID

	VarietyID *vo.ID

	ProductionUnitID vo.ID

	Method GrowingMethod

	Granularity ProductionGranularity

	ProtocolID *vo.ID

	Status GrowingStatus

	CurrentStageID *vo.ID

	LayoutSnapshotID *vo.ID

	ExpectedHarvestAt *time.Time

	StartedAt *time.Time

	CompletedAt *time.Time

	Metadata vo.Metadata

	CreatedAt time.Time

	UpdatedAt time.Time
}

type Aggregate struct {
	ev.AggregateRoot

	Root GrowingCycle
}

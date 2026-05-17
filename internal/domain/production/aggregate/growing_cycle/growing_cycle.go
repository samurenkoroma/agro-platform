package growingcycle

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type GrowingCycle struct {
	ID vo.ID

	FarmID vo.ID

	CropID vo.ID

	VarietyID *vo.ID

	ProductionUnitID vo.ID

	Method GrowingMethod

	ProtocolID *vo.ID

	Status GrowingStatus

	CurrentStageID *vo.ID

	LayoutSnapshotID *vo.ID

	StartedAt time.Time

	ExpectedHarvestAt *time.Time

	CompletedAt *time.Time

	Metadata vo.Metadata
}

type Aggregate struct {
	ev.AggregateRoot

	Root GrowingCycle
}

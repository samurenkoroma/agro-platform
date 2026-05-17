package plant

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Plant struct {
	ID vo.ID

	GrowingCycleID vo.ID

	CropID vo.ID

	VarietyID *vo.ID

	ProductionUnitID vo.ID

	SlotID *vo.ID

	SubstrateID *vo.ID

	Status PlantStatus

	Health PlantHealth

	CurrentStageID *vo.ID

	PlantedAt time.Time

	TransplantedAt *time.Time

	HarvestedAt *time.Time

	DiscardedAt *time.Time

	Metadata vo.Metadata

	CreatedAt time.Time

	UpdatedAt time.Time
}

type Aggregate struct {
	ev.AggregateRoot

	Root Plant
}

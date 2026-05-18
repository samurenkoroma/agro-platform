package sensor

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Sensor struct {
	ID vo.ID

	Type Type

	FarmID vo.ID

	ProductionUnitID *vo.ID

	ClimateZoneID *vo.ID

	Status Status

	Value Value

	Metadata vo.Metadata

	CreatedAt time.Time

	UpdatedAt time.Time

	ArchivedAt *time.Time
}

type Aggregate struct {
	ev.AggregateRoot

	Root Sensor
}

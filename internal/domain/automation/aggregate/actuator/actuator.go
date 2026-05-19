package actuator

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Actuator struct {
	ID               vo.ID
	Type             Type
	FarmID           vo.ID
	ProductionUnitID *vo.ID
	Status           Status
	Metadata         vo.Metadata
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ArchivedAt       *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Actuator
}

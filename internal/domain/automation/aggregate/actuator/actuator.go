package actuator

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Actuator struct {
	ev.BaseAggregate
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

func New(farmID vo.ID, actuatorType Type) *Actuator {
	now := time.Now()

	root := &Actuator{
		ID:       vo.NewID(),
		Type:     actuatorType,
		FarmID:   farmID,
		Status:   Enabled,
		Metadata: vo.NewMetadata(),

		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewActuatorCreated(root.ID))

	return root
}

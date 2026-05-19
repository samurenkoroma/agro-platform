package actuator

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, actuatorType Type) *Aggregate {
	now := time.Now()

	root := Actuator{
		ID:       vo.NewID(),
		Type:     actuatorType,
		FarmID:   farmID,
		Status:   Enabled,
		Metadata: vo.NewMetadata(),

		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewActuatorCreated(root.ID))

	return a
}

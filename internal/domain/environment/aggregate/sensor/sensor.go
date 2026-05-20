package sensor

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Sensor struct {
	ev.BaseAggregate
	ID               vo.ID
	Type             Type
	FarmID           vo.ID
	ProductionUnitID *vo.ID
	ClimateZoneID    *vo.ID
	Status           Status
	Value            Value
	Metadata         vo.Metadata
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ArchivedAt       *time.Time
}

func New(farmID vo.ID, sensorType Type) *Sensor {
	now := time.Now()

	root := &Sensor{
		ID:        vo.NewID(),
		Type:      sensorType,
		FarmID:    farmID,
		Status:    Online,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewSensorCreated(root.ID))

	return root
}

func (a *Sensor) UpdateValue(v float64, ts time.Time) {
	a.Value.Current = &v
	a.Value.LastTimestamp = &ts

	a.UpdatedAt = time.Now()

	a.AddEvent(NewSensorUpdated(a.ID))
}

func (a *Sensor) SetStatus(status Status) error {

	if a.ArchivedAt != nil {
		return ErrArchivedSensor
	}

	a.Status = status

	a.UpdatedAt = time.Now()

	a.AddEvent(NewSensorStatusChanged(a.ID))

	return nil
}

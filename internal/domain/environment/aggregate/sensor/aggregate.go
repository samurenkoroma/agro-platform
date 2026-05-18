package sensor

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, sensorType Type) *Aggregate {
	now := time.Now()

	root := Sensor{
		ID:        vo.NewID(),
		Type:      sensorType,
		FarmID:    farmID,
		Status:    Online,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewSensorCreated(root.ID))

	return a
}

func (a *Aggregate) UpdateValue(v float64, ts time.Time) {
	a.Root.Value.Current = &v
	a.Root.Value.LastTimestamp = &ts

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewSensorUpdated(a.Root.ID))
}

func (a *Aggregate) SetStatus(status Status) error {

	if a.Root.ArchivedAt != nil {
		return ErrArchivedSensor
	}

	a.Root.Status = status

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewSensorStatusChanged(a.Root.ID))

	return nil
}

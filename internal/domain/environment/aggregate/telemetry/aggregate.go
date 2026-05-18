package telemetry

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(sensorID vo.ID, value float64) *Aggregate {
	root := Telemetry{
		ID:       vo.NewID(),
		SensorID: sensorID,
		Value: Value{
			Value: value,
		},
		Timestamp: time.Now(),
		Metadata:  vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewTelemetryCreated(root.ID))

	return a
}

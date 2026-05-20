package telemetry

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Telemetry struct {
	ev.BaseAggregate
	ID        vo.ID
	SensorID  vo.ID
	Value     Value
	Timestamp time.Time
	Metadata  vo.Metadata
}

func New(sensorID vo.ID, value float64) *Telemetry {
	root := &Telemetry{
		ID:       vo.NewID(),
		SensorID: sensorID,
		Value: Value{
			Value: value,
		},
		Timestamp: time.Now(),
		Metadata:  vo.NewMetadata(),
	}

	root.AddEvent(NewTelemetryCreated(root.ID))

	return root
}

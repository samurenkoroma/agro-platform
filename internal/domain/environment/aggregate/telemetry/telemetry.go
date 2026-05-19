package telemetry

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Telemetry struct {
	ID vo.ID

	SensorID vo.ID

	Value Value

	Timestamp time.Time

	Metadata vo.Metadata
}

type Aggregate struct {
	ev.AggregateRoot

	Root Telemetry
}

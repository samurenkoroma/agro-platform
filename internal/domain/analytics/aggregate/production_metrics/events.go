package productionmetrics

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventMetricsCalculated = "metrics.calculated"
	EventMetricsUpdated    = "metrics.updated"
)

type MetricsCalculated struct {
	ev.BaseEvent
}

func NewMetricsCalculated(id vo.ID) MetricsCalculated {
	return MetricsCalculated{ev.NewBaseEvent(id, EventMetricsCalculated)}
}

type MetricsUpdated struct {
	ev.BaseEvent
}

func NewMetricsUpdated(id vo.ID) MetricsUpdated {
	return MetricsUpdated{ev.NewBaseEvent(id, EventMetricsUpdated)}
}

package telemetry

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventTelemetryCreated  = "telemetry.created"
	EventTelemetryImported = "telemetry.imported"
	EventTelemetryArchived = "telemetry.archived"
)

type TelemetryCreated struct {
	ev.BaseEvent
}

func NewTelemetryCreated(id vo.ID) TelemetryCreated {
	return TelemetryCreated{BaseEvent: ev.NewBaseEvent(id, EventTelemetryCreated)}
}

type TelemetryImported struct {
	ev.BaseEvent
}

func NewTelemetryImported(id vo.ID) TelemetryImported {
	return TelemetryImported{BaseEvent: ev.NewBaseEvent(id, EventTelemetryImported)}
}

type TelemetryArchived struct {
	ev.BaseEvent
}

func NewTelemetryArchived(id vo.ID) TelemetryArchived {
	return TelemetryArchived{BaseEvent: ev.NewBaseEvent(id, EventTelemetryArchived)}
}

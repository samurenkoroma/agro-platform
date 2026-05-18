package climatezone

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventClimateZoneCreated       = "climate_zone.created"
	EventClimateZoneTargetUpdated = "climate_zone.target.updated"
	EventClimateZoneArchived      = "climate_zone.archived"
)

type ClimateZoneCreated struct {
	ev.BaseEvent
}

func NewClimateZoneCreated(id vo.ID) ClimateZoneCreated {
	return ClimateZoneCreated{ev.NewBaseEvent(id, EventClimateZoneCreated)}
}

type ClimateZoneTargetUpdated struct {
	ev.BaseEvent
}

func NewClimateZoneTargetUpdated(id vo.ID) ClimateZoneTargetUpdated {
	return ClimateZoneTargetUpdated{ev.NewBaseEvent(id, EventClimateZoneTargetUpdated)}
}

type ClimateZoneArchived struct {
	ev.BaseEvent
}

func NewClimateZoneArchived(id vo.ID) ClimateZoneArchived {
	return ClimateZoneArchived{ev.NewBaseEvent(id, EventClimateZoneArchived)}
}

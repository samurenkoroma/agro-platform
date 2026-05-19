package actuator

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventActuatorCreated       = "actuator.created"
	EventActuatorStatusChanged = "actuator.status.changed"
	EventActuatorArchived      = "actuator.archived"
)

type ActuatorCreated struct {
	ev.BaseEvent
}

func NewActuatorCreated(id vo.ID) ActuatorCreated {
	return ActuatorCreated{ev.NewBaseEvent(id, EventActuatorCreated)}
}

type ActuatorStatusChanged struct {
	ev.BaseEvent
}

func NewActuatorStatusChanged(id vo.ID) ActuatorStatusChanged {
	return ActuatorStatusChanged{ev.NewBaseEvent(id, EventActuatorStatusChanged)}
}

type ActuatorArchived struct {
	ev.BaseEvent
}

func NewActuatorArchived(id vo.ID) ActuatorArchived {
	return ActuatorArchived{ev.NewBaseEvent(id, EventActuatorArchived)}
}

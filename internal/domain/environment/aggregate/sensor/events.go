package sensor

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventSensorCreated        = "sensor.created"
	EventSensorUpdated        = "sensor.updated"
	EventSensorAttachedToUnit = "sensor.production_unit.attached"
	EventSensorAttachedToZone = "sensor.climate_zone.attached"
	EventSensorStatusChanged  = "sensor.status.changed"
	EventSensorArchived       = "sensor.archived"
)

type SensorCreated struct{ ev.BaseEvent }

func NewSensorCreated(id vo.ID) SensorCreated {
	return SensorCreated{ev.NewBaseEvent(id, EventSensorCreated)}
}

type SensorUpdated struct{ ev.BaseEvent }

func NewSensorUpdated(id vo.ID) SensorUpdated {
	return SensorUpdated{ev.NewBaseEvent(id, EventSensorUpdated)}
}

type SensorArchived struct{ ev.BaseEvent }

func NewSensorArchived(id vo.ID) SensorArchived {
	return SensorArchived{ev.NewBaseEvent(id, EventSensorArchived)}
}

type SensorStatusChanged struct{ ev.BaseEvent }

func NewSensorStatusChanged(id vo.ID) SensorStatusChanged {
	return SensorStatusChanged{ev.NewBaseEvent(id, EventSensorStatusChanged)}
}

type SensorAttachedToUnit struct{ ev.BaseEvent }

func NewSensorAttachedToUnit(id vo.ID) SensorAttachedToUnit {
	return SensorAttachedToUnit{ev.NewBaseEvent(id, EventSensorAttachedToUnit)}
}

type SensorAttachedToZone struct{ ev.BaseEvent }

func NewSensorAttachedToZone(id vo.ID) SensorAttachedToZone {
	return SensorAttachedToZone{ev.NewBaseEvent(id, EventSensorAttachedToZone)}
}

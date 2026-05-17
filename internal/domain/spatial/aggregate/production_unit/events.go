package productionunit

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventCreated = "production_unit.created"

	EventAttached = "production_unit.attached"
)

type ProductionUnitCreated struct {
	ev.BaseEvent
}

func NewProductionUnitCreated(
	id vo.ID,
) ProductionUnitCreated {

	return ProductionUnitCreated{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventCreated,
		),
	}
}

type ProductionUnitAttached struct {
	ev.BaseEvent

	ParentID vo.ID
}

func NewProductionUnitAttached(
	id vo.ID,
	parent vo.ID,
) ProductionUnitAttached {

	return ProductionUnitAttached{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventAttached,
		),

		ParentID: parent,
	}
}
